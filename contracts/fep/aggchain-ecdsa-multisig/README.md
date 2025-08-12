# aggchain-ecdsa-multisig contracts

All the files and directories within this directory have been generated using the import-contracts command of the CLI in this repo.
The ABI and the binnaries of the smart contracts have been extracted from [agglayer-contracts repo](https://github.com/agglayer/agglayer-contracts.git), using the version feature/multisig_consensus (commit 2c2ce6cb108b472cb6583aaa6a35628123a2107a)

The CLI command used to generate the contracts: `$ go run ./cmd import-contracts -cv feature/multisig_consensus -proving-schema fep -alias aggchain-ecdsa-multisig -node 22`
