#!/usr/bin/env bash

set -x

grided tx clp remove-liquidity-units \
  --from $GRID_ACT \
  --keyring-backend test \
  --symbol cusdt \
  --withdrawUnits 10000000000 \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y