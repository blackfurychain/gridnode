# Ceth Receiver Account Setting
When a Gridironchain account sends a lock/burn transaction, the relayers need to respond to that and send a transaction to an Ethereum smart contract.
When that happens, Gridironchain needs to charge a transaction fee. The cEth in the lock/burn transaction will go to a defined account.
That account can be changed using gridnoded commands, and the cEth in that account can be sent to other accounts.

## Set the account
At runtime, we can update the account using transaction.
The transaction is privileged, and only the admin account can set it.
The admin account is the oracle admin account. The command is as follows:

```bash
gridnoded tx ethbridge update_ceth_receiver_account $oracle_admin_address $ceth_receiver_account --node tcp://rpc.gridchain.finance:80 --keyring-backend=file --chain-id=gridchain --from=$oracle_admin_moniker --fees=100000fury
```

## Rescue the Ceth
Before setting the account, some cEth can be locked in the ethbridge module, so we need a method to rescue the cEth.
Similar to the account setting method, the transaction is privileged and only the admin account can call it.
It will transfer the cEth from ethbridge module to an specific account.

```bash
gridnoded tx ethbridge rescue_ceth $oracle_admin_address $ceth_receiver_account $ceth_amount --node tcp://rpc.gridchain.finance:80 --keyring-backend=file --chain-id=gridchain --from=$oracle_admin_moniker --fees=100000fury
```
