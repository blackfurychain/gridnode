#!/usr/bin/env bash

echo "Creating pools ceth and cdash"
gridnoded tx clp create-pool --from grid --symbol ceth --nativeAmount 1000000000000000000 --externalAmount 1000000000000000000  --yes --chain-id=localnet --keyring-backend=test

sleep 5
gridnoded tx clp create-pool --from grid --symbol cdash --nativeAmount 1000000000000000000 --externalAmount 1000000000000000000  --yes --chain-id=localnet --keyring-backend=test

echo "Query all pools"
sleep 8
gridnoded query clp pools

echo "Query specific pool"
sleep 8
gridnoded query clp pool ceth

echo "Query Liquidity Provider / Pool creator is the first lp for the pool"
sleep 8
gridnoded query clp lp ceth $(gridnoded keys show grid -a )

echo "adding more liquidity"
sleep 8
gridnoded tx clp add-liquidity --from grid --symbol ceth --nativeAmount 10000000000000000000 --externalAmount 10000000000000000000 --yes --chain-id=localnet --keyring-backend=test

echo "swap"
sleep 8
gridnoded tx clp swap --from grid --sentSymbol ceth --receivedSymbol cdash --sentAmount 1000000000000000000 --minReceivingAmount 0 --yes

echo "removing Liquidity"
sleep 8
gridnoded tx clp remove-liquidity --from grid --symbol ceth --wBasis 5001 --asymmetry -1 --yes

echo "removing more Liquidity"
sleep 8
gridnoded tx clp remove-liquidity --from grid --symbol ceth --wBasis 5001 --asymmetry -1 --yes

echo "decommission pool"
sleep 8
gridnoded tx clp decommission-pool --from grid --symbol ceth --yes

echo "gridnoded query clp pools "
echo "Should list only 1 pool , user has been added as admin"



