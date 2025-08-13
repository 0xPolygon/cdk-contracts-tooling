# aggchain-ecdsa-multisig contracts

All the files and directories within this directory have been generated using the import-contracts command of the CLI in this repo.
The ABI and the binnaries of the smart contracts have been extracted from [agglayer-contracts repo](https://github.com/agglayer/agglayer-contracts.git), using the version feature/multisig_consensus (commit 659db80acd840bf9e887c323cd81da0f4206b79d)

The CLI command used to generate the contracts: `$ go run ./cmd import-contracts -cv feature/multisig_consensus -proving-schema fep -alias aggchain-ecdsa-multisig -node 22`
