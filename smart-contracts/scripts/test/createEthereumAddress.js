const BN = require('bn.js');

module.exports = async (cb) => {
    const Web3 = require("web3");

    const gridironchainUtilities = require('./gridironchainUtilities')
    const contractUtilites = require('./contractUtilities');

    const logging = gridironchainUtilities.configureLogging(this);

    const argv = gridironchainUtilities.processArgs(this, {
        ...gridironchainUtilities.sharedYargOptions,
        'count': {
            describe: 'how many addresses to create',
            default: 1,
        },
    });

    const web3 = contractUtilites.buildWeb3(this, argv, logging);

    if (argv.count > 1) {
        const result = [];
        for (let i = 0; i <= argv.count; i = i + 1) {
            result.push(web3.eth.accounts.create());
        }
        console.log(JSON.stringify(result));
    } else {
        console.log(JSON.stringify(web3.eth.accounts.create()));
    }


    return cb();
};
