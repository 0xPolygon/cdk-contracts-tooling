package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/rollup"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
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

// LoadGenesisAllocs loads the genessis allocations from the provided baseDir followed by the relative path of the genesis allocs file
func LoadGenesisAllocs(baseDir, genesisPath string) (interface{}, error) {
	fmt.Println("loading genesis allocations file")
	genesisData, err := os.ReadFile(path.Join(baseDir, genesisPath))
	if err != nil {
		return nil, err
	}

	var genesis interface{}
	err = json.Unmarshal(genesisData, &genesis)
	if err != nil {
		return nil, err
	}

	return genesis, nil
}
