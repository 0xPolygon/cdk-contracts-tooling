package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
)

type RPCs struct {
	Networks map[string]RPCConfig
}

type RPCConfig struct {
	ChainID uint64
	URL     string
}

func LoadRPCs() (*RPCs, error) {
	var r RPCs
	return &r, LoadFromFile("./rpcs.toml", &r)
}

func (r *RPCs) GetClient(network string) (*ethclient.Client, error) {
	rpcConfig, ok := r.Networks[network]
	if !ok {
		return nil, fmt.Errorf("there is no RPC config for network %s. Please add it to rpcs.toml", network)
	}
	return ethclient.Dial(rpcConfig.URL)
}

func (r *RPCs) GetChainID(network string) (uint64, error) {
	rpcConfig, ok := r.Networks[network]
	if !ok {
		return 0, fmt.Errorf("there is no RPC config for network %s. Please add it to rpcs.toml", network)
	}
	return rpcConfig.ChainID, nil
}
