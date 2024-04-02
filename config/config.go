package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/0xPolygon/cdk-contracts-tooling/rollup"
	"github.com/0xPolygon/cdk-contracts-tooling/rollupmanager"
)

func Load(l1Network, rmAlias, rAlias, baseDir string) (*RPCs, *rollupmanager.RollupManager, *rollup.Rollup, interface{}, error) {
	fmt.Println("loading the rollup manager info from file")
	rollupManagerPath := path.Join(baseDir, "networks", l1Network, rmAlias)
	rm, err := rollupmanager.LoadFromFile(nil, path.Join(rollupManagerPath, "rollupManager.json"))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	fmt.Println("loading RPC info")
	rpcs, err := LoadRPCs()
	if err != nil {
		return nil, nil, nil, nil, err
	}

	fmt.Println("loading the rollup info from file")
	rollupPath := path.Join(rollupManagerPath, "rollups")
	r, err := rollup.LoadFromFile(nil, path.Join(rollupPath, rAlias+".json"))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	fmt.Println("loading genesis file")
	genesisData, err := os.ReadFile(path.Join(baseDir, "genesis", r.GenesisRoot.Hex()+".json"))
	if err != nil {
		return nil, nil, nil, nil, err
	}
	var genesis interface{}
	err = json.Unmarshal(genesisData, &genesis)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return rpcs, rm, r, genesis, nil
}
