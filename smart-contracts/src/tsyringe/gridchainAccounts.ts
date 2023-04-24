import {HardhatRuntimeEnvironment} from "hardhat/types";
import {HardhatRuntimeEnvironmentToken,} from "./injectionTokens";
import {SignerWithAddress} from "@nomiclabs/hardhat-ethers/signers";
import {inject, injectable} from "tsyringe";
import {isHardhatRuntimeEnvironment} from "./hardhatSupport";

/**
 * The accounts necessary for testing a gridironchain system
 */
export class GridironchainAccounts {
    constructor(
        readonly operatorAccount: SignerWithAddress,
        readonly ownerAccount: SignerWithAddress,
        readonly pauserAccount: SignerWithAddress,
        readonly validatatorAccounts: Array<SignerWithAddress>,
        readonly availableAccounts: Array<SignerWithAddress>
    ) {
    }
}

/**
 * Note that the hardhat environment provides accounts as promises, so
 * we need to wrap a GridironchainAccounts in a promise.
 */
@injectable()
export class GridironchainAccountsPromise {
    accounts: Promise<GridironchainAccounts>

    constructor(accounts: Promise<GridironchainAccounts>);
    constructor(@inject(HardhatRuntimeEnvironmentToken) hardhatOrAccounts: HardhatRuntimeEnvironment | Promise<GridironchainAccounts>) {
        if (isHardhatRuntimeEnvironment(hardhatOrAccounts)) {
            this.accounts = hreToGridironchainAccountsAsync(hardhatOrAccounts)
        } else {
            this.accounts = hardhatOrAccounts
        }
    }
}

export async function hreToGridironchainAccountsAsync(hardhat: HardhatRuntimeEnvironment): Promise<GridironchainAccounts> {
    const accounts = await hardhat.ethers.getSigners()
    const [operatorAccount, ownerAccount, pauserAccount, validator1Account, ...extraAccounts] = accounts
    return new GridironchainAccounts(operatorAccount, ownerAccount, pauserAccount, [validator1Account], extraAccounts)
}
