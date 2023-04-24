#!/bin/sh

# Unbond liquidity
gridnoded tx clp unbond-liquidity \
--from grid --keyring-backend test \
--fees 100000000000000000fury \
--symbol ceth \
--units 1000000000000000000 \
--chain-id localnet \
--broadcast-mode block \
-y