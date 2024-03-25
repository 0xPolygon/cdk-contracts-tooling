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

### Examples

#### Generate genesis file for a node (`cdk-validium-node` or `zkevm-node`)

1. Import the rollup manager: `go run ./cmd import-rm -l1 sepolia -addr 0x32d33d5137a7cffb54c5bf8371172bcec5f310ff -alias cardona`. In this example:
    - `-l1 sepolia` is the L1 network, make sure that your `wallets.toml` file contains a valid RPC for the L1 network this CDK belongs to
    - `-addr 0x32d33d5137a7cffb54c5bf8371172bcec5f310ff` is the L1 address where the rollup manager is deployed
    - `-alias cardona` is an arbitrary name. You can really name this however you want, but this is going to be used later to reference this rollup manager deployment
2. Import the rollup: `go run ./cmd import-r -l1 sepolia -rm cardona -r API3 -chainid 879490799`. In this example:
    - `-l1 sepolia` is the L1 network, make sure that your `wallets.toml` file contains a valid RPC for the L1 network this CDK belongs to
    - `-rm cardona` rollup manager referenced by the alias used in the previous step
    - `-chainid 879490799` is the chain ID of the rollup (L2 Chain ID)
    - `-alias API3` is an arbitrary name. You can really name this however you want, but this is going to be used later to reference this rollup deployment
3. Generate the file: `go run ./cmd genesis -l1 sepolia -rm cardona -r API3 -output API3.json`. In this example:
    - `-l1 sepolia` is the L1 network, make sure that your `wallets.toml` file contains a valid RPC for the L1 network this CDK belongs to
    - `-rm cardona` rollup manager referenced by the alias used in the previous step
    - `-r API3` rollup referenced by the alias used in the previous step
    - `-output API3.json` file where the genesis will be stored

Note that step 1 only needs to be done once, if there are multiple CDKs attaced to the same rollup manager, with a single run it will be enough