#!/bin/sh

# Create fury/ceth; 
gridnoded tx clp create-pool \
--symbol ceth \
--nativeAmount 2000000000000000000 \
--externalAmount 2000000000000000000 \
--from grid --keyring-backend test \
--fees 100000000000000000fury \
--chain-id localnet \
--broadcast-mode block \
-y