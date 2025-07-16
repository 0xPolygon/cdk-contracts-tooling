# l2-sovereign-chain contracts

All the files and directories within this directory have been generated using the import-contracts command of the CLI in this repo.
The ABI and the binnaries of the smart contracts have been extracted from [agglayer-contracts repo](https://github.com/agglayer/agglayer-contracts.git), using the version feature/descentralized_aggoracle (commit 26053433f526fe07f7455d3f2209c4f7b5354dab)

The CLI command used to generate the contracts: `$ go run ./cmd import-contracts --contracts-version feature/descentralized_aggoracle --contracts-alias l2-sovereign-chain --proving-schema pp --node-version 18`
