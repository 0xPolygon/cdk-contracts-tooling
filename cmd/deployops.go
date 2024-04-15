package main

import (
	"errors"
	"math/big"
	"sync"

	"github.com/0xPolygon/cdk-contracts-tooling/contracts/etrog/polygonzkevmdeployer"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/sha3"
)

// hasherPool holds LegacyKeccak256 hashers for rlpHash.
var hasherPool = sync.Pool{
	New: func() interface{} { return sha3.NewLegacyKeccak256() },
}

type scalableSigner struct{}

func (s scalableSigner) ChainID() *big.Int {
	return nil
}

func (s scalableSigner) Equal(s2 types.Signer) bool {
	_, ok := s2.(types.FrontierSigner)
	return ok
}

func (fs scalableSigner) Sender(tx *types.Transaction) (common.Address, error) {
	if tx.Type() != types.LegacyTxType {
		return common.Address{}, types.ErrTxTypeNotSupported
	}
	v, r, s := tx.RawSignatureValues()
	return recoverPlain(fs.Hash(tx), r, s, v, false)
}

// SignatureValues returns signature values. This signature
// needs to be in the [R || S || V] format where V is 0 or 1.
func (fs scalableSigner) SignatureValues(tx *types.Transaction, sig []byte) (r, s, v *big.Int, err error) {
	if tx.Type() != types.LegacyTxType {
		return nil, nil, nil, types.ErrTxTypeNotSupported
	}
	r, s, v = hardcodedSignature()
	return r, s, v, nil
}

// rlpHash encodes x and hashes the encoded bytes.
func rlpHash(x interface{}) (h common.Hash) {
	sha := hasherPool.Get().(crypto.KeccakState)
	defer hasherPool.Put(sha)
	sha.Reset()
	rlp.Encode(sha, x)
	sha.Read(h[:])
	return h
}

// Hash returns the hash to be signed by the sender.
// It does not uniquely identify the transaction.
func (fs scalableSigner) Hash(tx *types.Transaction) common.Hash {
	return rlpHash([]interface{}{
		tx.Nonce(),
		tx.GasPrice(),
		tx.Gas(),
		tx.To(),
		tx.Value(),
		tx.Data(),
	})
}

func hardcodedSignature() (r, s, v *big.Int) {
	r = new(big.Int).SetBytes(common.Hex2Bytes("0x00000000000000000000000000000000000000000000000000000005ca1ab1e0"))
	s = new(big.Int).SetBytes(common.Hex2Bytes("0x00000000000000000000000000000000000000000000000000000005ca1ab1e0"))
	v = new(big.Int).SetBytes([]byte{27})
	return r, s, v
}

func recoverPlain(sighash common.Hash, R, S, Vb *big.Int, homestead bool) (common.Address, error) {
	if Vb.BitLen() > 8 {
		return common.Address{}, types.ErrInvalidSig
	}
	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S, homestead) {
		return common.Address{}, types.ErrInvalidSig
	}
	// encode the signature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, crypto.SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
	// recover the public key from the signature
	pub, err := crypto.Ecrecover(sighash[:], sig)
	if err != nil {
		return common.Address{}, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return common.Address{}, errors.New("invalid public key")
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pub[1:])[12:])
	return addr, nil
}

func create2Deployment(
	cliCtx *cli.Context,
	client *ethclient.Client,
	auth *bind.TransactOpts,
	polgonZKEVMDeployerContract *polygonzkevmdeployer.Polygonzkevmdeployer,
	polgonZKEVMDeployerContractAddr common.Address,
	salt [32]byte,
	deployTransaction []byte,
	dataCall []byte,
	confirmationQuestion string,
) (common.Address, error) {
	// Precalculate addr
	initHash := crypto.Keccak256(deployTransaction)
	precalculatedAddressDeployed := crypto.CreateAddress2(polgonZKEVMDeployerContractAddr, salt, initHash)

	// Check if already deployed
	code, err := client.CodeAt(cliCtx.Context, precalculatedAddressDeployed, nil)
	if err != nil {
		return common.Address{}, err
	}
	if len(code) > 0 {
		return precalculatedAddressDeployed, nil
	}

	// Deploy
	if len(dataCall) > 0 {
		// Using create2 and call
		_, err := sendTxWithConfirmation(
			cliCtx, client, confirmationQuestion,
			func() (*types.Transaction, common.Address, error) {
				tx, err := polgonZKEVMDeployerContract.DeployDeterministicAndCall(auth, nil, salt, deployTransaction, dataCall)
				return tx, common.Address{}, err
			},
		)
		if err != nil {
			return common.Address{}, err
		}
	} else {
		// Using create2
		_, err := sendTxWithConfirmation(
			cliCtx, client, confirmationQuestion,
			func() (*types.Transaction, common.Address, error) {
				tx, err := polgonZKEVMDeployerContract.DeployDeterministic(auth, nil, salt, deployTransaction)
				return tx, common.Address{}, err
			},
		)
		if err != nil {
			return common.Address{}, err
		}
	}

	return precalculatedAddressDeployed, nil
}
