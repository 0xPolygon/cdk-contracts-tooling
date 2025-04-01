# l2-sovereign-chain contracts

All the files and directories within this directory have been generated using the import-contracts command of the CLI in this repo.
The ABI and the binnaries of the smart contracts have been extracted from [agg-contracts-internal repo](https://github.com/agglayer/agg-contracts-internal.git), using the version v10.0.0-rc.4 (commit 2bdd64e6cbec65f0b2da4e47068fdef46445b21e)

The CLI command used to generate the contracts: `$ go run ./cmd import-contracts -cv v10.0.0-rc.4 -proving-schema pp -alias l2-sovereign-chain -node 20`
