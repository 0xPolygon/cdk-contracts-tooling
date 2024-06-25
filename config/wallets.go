package config

import (
	"bufio"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

type Wallets struct {
	Wallets map[string]WalletConfig
}

type WalletConfig struct {
	Type       walletType
	FilePath   string
	PrivateKey string
}

type walletType string

const (
	wtText walletType = "PlainTextPrivateKey"
	wtFile walletType = "EncryptedFile"
)

func LoadWallets() (*Wallets, error) {
	var w Wallets
	return &w, LoadFromFile("./wallets.toml", &w)
}

func (w *Wallets) GetPrivateKey(address, pass string) (*ecdsa.PrivateKey, error) {
	wallet, ok := w.Wallets[address]
	if !ok {
		return nil, fmt.Errorf("there is no wallet config for address %s. Please add it to wallets.toml", address)
	}
	switch wallet.Type {
	case wtText:
		return crypto.HexToECDSA(strings.TrimPrefix(wallet.PrivateKey, "0x"))
	case wtFile:
		return loadKeyFromFile(wallet.FilePath, pass)
	default:
		return nil, fmt.Errorf("unsupported wallet type: %s. Should be one of [%s, %s]", wallet.Type, wtText, wtFile)
	}
}

func (w *Wallets) GetAuth(address, pass string, chainID uint64) (*bind.TransactOpts, error) {
	pk, err := w.GetPrivateKey(address, pass)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(pk, new(big.Int).SetUint64(chainID))
}

// loadKeyFromFile creates an instance of a keystore key from a keystore file
func loadKeyFromFile(path, pass string) (*ecdsa.PrivateKey, error) {
	keystoreEncrypted, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		return nil, err
	}
	if pass == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Enter password for the encrypted private key at: %s\n\nINPUT PASSWORD:", path)
		pass, err = reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		pass = strings.TrimSuffix(pass, "\n")
	}
	key, err := keystore.DecryptKey(keystoreEncrypted, pass)
	if err != nil {
		return nil, err
	}
	return key.PrivateKey, nil
}
