const BN = require('bn.js');

module.exports = async (cb) => {
    const Web3 = require("web3");

    const gridironchainUtilities = require('./gridironchainUtilities')
    const contractUtilites = require('./contractUtilities');

    const logging = gridironchainUtilities.configureLogging(this);

    const argv = gridironchainUtilities.processArgs(this, {
        ...gridironchainUtilities.sharedYargOptions,
        ...gridironchainUtilities.bridgeBankAddressYargOptions,
        ...gridironchainUtilities.symbolYargOption,
        ...gridironchainUtilities.amountYargOption,
        'limit_amount': {
            describe: 'an amount',
            demandOption: true
        },
        'operator_address': {
            type: "string",
            demandOption: true,
        },
        'token_name': {
            type: "string",
            demandOption: true,
        },
        'decimals': {
            type: "number",
            demandOption: true,
        },
    });

    const amount = new BN(argv.amount, 10);
    const limitAmount = new BN(argv.limit_amount, 10);

    const standardOptions = {
        from: argv.operator_address
    }

    const newTokenBuilder = await contractUtilites.buildBaseContract(this, argv, logging, "GridironchainTestToken");
    const newToken = await newTokenBuilder.new(argv.token_name, argv.symbol, argv.decimals, standardOptions);

    const bridgeBankContract = await contractUtilites.buildContract(this, argv, logging, "BridgeBank", argv.bridgebank_address);

    const operator_address = argv.operator_address;

    const updateWhiteListResult = await bridgeBankContract.updateEthWhiteList(newToken.address, true, standardOptions);

    const token_destination = argv.operator_address;

    await newToken.mint(token_destination, amount, standardOptions);

    await newToken.approve(bridgeBankContract.address, amount.toString(), standardOptions);

    const result = {
        destination: token_destination,
        "amount": amount.toString(),
        "newtoken_address": newToken.address,
        "newtoken_symbol": argv.symbol,
    }
    console.log(JSON.stringify(result));

    return cb();
};
