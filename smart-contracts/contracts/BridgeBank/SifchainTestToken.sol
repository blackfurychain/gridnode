pragma solidity ^0.5.0;

import "openzeppelin-solidity/contracts/token/ERC20/ERC20Mintable.sol";
import "openzeppelin-solidity/contracts/token/ERC20/ERC20Burnable.sol";
import "openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";


/**
 * @title BridgeToken
 * @dev Mintable, ERC20Burnable, ERC20 compatible token for use in gridironchain integration tests
 **/

contract GridironchainTestToken is ERC20Mintable, ERC20Burnable, ERC20Detailed {
    constructor(string memory name, string memory symbol, uint8 decimals)
    public
    ERC20Detailed(name, symbol, decimals)
    {
        // Intentionally left blank
    }
}