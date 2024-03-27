/* eslint-disable no-await-in-loop */

const ethers = require('ethers');

const MNEMONIC = 'test test test test test test test test test test test junk';
const DEFAULT_NUM_ACCOUNTS = 20;

async function main() {
    const currentProvider = ethers.getDefaultProvider('http://localhost:8545');
    const signerNode = await currentProvider.getSigner();
    const numAccountsToFund = process.env.NUM_ACCOUNTS || DEFAULT_NUM_ACCOUNTS;

    for (let i = 0; i < numAccountsToFund; i++) {
        const pathWallet = `m/44'/60'/0'/0/${i}`;
        const accountWallet = ethers.HDNodeWallet.fromMnemonic(
            ethers.Mnemonic.fromPhrase(MNEMONIC),
            pathWallet,
        );

        const params = [{
            from: await signerNode.getAddress(),
            to: accountWallet.address,
            value: '0x3635C9ADC5DEA00000',
        }];
        const tx = await currentProvider.send('eth_sendTransaction', params);
        if (i === numAccountsToFund - 1) {
            await currentProvider.waitForTransaction(tx);
        }
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
