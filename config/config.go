package config

import (
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/rollup"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
	"github.com/ethereum/go-ethereum/common"
)

// Load RPCs configuration, rollup manager metadata and rollup metadata
func Load(l1Network, rmAlias, rAlias, baseDir string) (*RPCs, *rollupmanager.RollupManager, *rollup.RollupMetadata, error) {
	fmt.Println("loading the rollup manager info from file")
	rollupManagerPath := path.Join(baseDir, "networks", l1Network, rmAlias)
	rm, err := rollupmanager.LoadFromFile(nil, path.Join(rollupManagerPath, "rollupManager.json"))
	if err != nil {
		return nil, nil, nil, err
	}

	fmt.Println("loading RPC info")
	rpcs, err := LoadRPCs()
	if err != nil {
		return nil, nil, nil, err
	}

	rollupPath := path.Join(rollupManagerPath, "rollups")
	rollupFile := path.Join(rollupPath, rAlias+".json")
	fmt.Println("loading the rollup info from file", rollupFile)
	r, err := rollup.LoadMetadataFromFile(rollupFile)
	if err != nil {
		return nil, nil, nil, err
	}

	return rpcs, rm, r, nil
}

// GenesisAllocs represents the full genesis JSON structure.
type GenesisAllocs struct {
	Root    common.Hash      `json:"root"`
	Genesis []GenesisAccount `json:"genesis"`
}

// GenesisAccount represents a single account or contract entry within the genesis array.
type GenesisAccount struct {
	ContractName string                      `json:"contractName,omitempty"`
	AccountName  string                      `json:"accountName,omitempty"`
	Balance      BigInt                      `json:"balance"`
	Nonce        uint64                      `json:"nonce,string"`
	Address      common.Address              `json:"address"`
	Bytecode     string                      `json:"bytecode,omitempty"`
	Storage      map[common.Hash]common.Hash `json:"storage,omitempty"`
}

// BigInt handles both numeric and string JSON encodings for balances.
type BigInt struct {
	*big.Int
}

func (b *BigInt) UnmarshalJSON(data []byte) error {
	// Try as a string first (e.g., "0" or "3402823...")
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		i, ok := new(big.Int).SetString(s, 10)
		if !ok {
			return fmt.Errorf("invalid big.Int string: %q", s)
		}
		b.Int = i
		return nil
	}

	// Try as a raw number (unquoted)
	var i int64
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	b.Int = big.NewInt(i)
	return nil
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	if b.Int == nil {
		return []byte(`"0"`), nil
	}
	return json.Marshal(b.String())
}

// LoadGenesisAllocs loads the genessis allocations from the provided baseDir followed by the relative path of the genesis allocs file
func LoadGenesisAllocs(baseDir, genesisPath string) (*GenesisAllocs, error) {
	fmt.Println("loading genesis allocations file")
	genesisData, err := os.ReadFile(path.Join(baseDir, genesisPath))
	if err != nil {
		return nil, err
	}

	var genesis GenesisAllocs
	err = json.Unmarshal(genesisData, &genesis)
	if err != nil {
		return nil, err
	}

	return &genesis, nil
}
