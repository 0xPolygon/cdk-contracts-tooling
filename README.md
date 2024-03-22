# cdk-contracts-tooling

Tools to interact with CDK smart contracts and Go bindings

## Requirements

Not all the commands have the same requirements, but here is the complete list of requirements (to be able to run all the commands):

- Go
- NVM installed under $NVM_DIR
- abigen

## Usage

From the root of this repo run `go run ./cmd`. This will print a help message explaining how to use all the commands

### Wallets

Some commands require using private keys. There are multiple ways to import private keys, but all of them require to:

1. Copy paste `wallets.example.toml` into `wallets.toml`
2. Edit the new file to include your wallets (the file has comments to make the needed changes comprehensible)

### RPCs

Some commands require using RPC endpoints. There are multiple ways to import this endpoints, but all of them require to:

1. Copy paste `rpcs.example.toml` into `rpcs.toml`
2. Edit the new file to include your RPC configurations (the file has comments to make the needed changes comprehensible)