module.exports = async (cb) => {
    const Web3 = require("web3");

    const gridchainUtilities = require('./gridchainUtilities')
    const contractUtilites = require('./contractUtilities');

    const logging = gridchainUtilities.configureLogging(this);

    try {
        const argv = gridchainUtilities.processArgs(this, {
            ...gridchainUtilities.sharedYargOptions,
            ...gridchainUtilities.transactionYargOptions,
        });

        logging.info(`sendBurnTx: ${JSON.stringify(argv, undefined, 2)}`);

        const bridgeBankContract = await contractUtilites.buildContract(this, argv, logging, "BridgeBank", argv.bridgebank_address);

        const result = {};

        const transactionParameters = {
            from: argv.ethereum_address,
        }

        await contractUtilites.setAllowance(this, argv.symbol, argv.amount, argv, logging, transactionParameters);

        logging.info(`sendBurnTx ${JSON.stringify(argv)}}`);

        result.burn = await bridgeBankContract.burn(
            Web3.utils.utf8ToHex(argv.gridchain_address),
            argv.symbol,
            argv.amount,
            transactionParameters,
        );

        console.log(JSON.stringify(result, undefined, 0));
    } catch (e) {
        console.error(`sendBurnTx error: ${e} ${e.message}`);
        throw(e);
    }

    return cb();
};
