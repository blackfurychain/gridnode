#!/usr/bin/env bash

echo "Creating pools ceth and cdash"
grided tx clp create-pool --from grid --symbol ceth --nativeAmount 20000000000000000000 --externalAmount 20000000000000000000  --yes

sleep 5
grided tx clp create-pool --from grid --symbol cdash --nativeAmount 20000000000000000000 --externalAmount 20000000000000000000  --yes


sleep 8
echo "Swap Native for Pegged - Sent fury Get ceth"
grided tx clp swap --from grid --sentSymbol fury --receivedSymbol ceth --sentAmount 2000000000000000000 --minReceivingAmount 0 --yes
sleep 8
echo "Swap Pegged for Native - Sent ceth Get fury"
grided tx clp swap --from grid --sentSymbol ceth --receivedSymbol fury --sentAmount 2000000000000000000 --minReceivingAmount 0 --yes
sleep 8
echo "Swap Pegged for Pegged - Sent ceth Get cdash"
grided tx clp swap --from grid --sentSymbol ceth --receivedSymbol cdash --sentAmount 2000000000000000000 --minReceivingAmount 0 --yes

grided q clp pools

