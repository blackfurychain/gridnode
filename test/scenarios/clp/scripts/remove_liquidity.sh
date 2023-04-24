#!/bin/sh

# Remove liquidity 
grided tx clp remove-liquidity \
--from grid --keyring-backend test \
--fees 100000000000000000fury \
--symbol ceth \
--wBasis 5000 --asymmetry 0 \
--chain-id localnet \
--broadcast-mode block \
-y