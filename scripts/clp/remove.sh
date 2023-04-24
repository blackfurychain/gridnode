#!/bin/sh

set -x

grided tx clp remove-liquidity-units \
  --withdrawUnits 1 \
  --symbol ceth \
  --from $GRID_ACT \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y