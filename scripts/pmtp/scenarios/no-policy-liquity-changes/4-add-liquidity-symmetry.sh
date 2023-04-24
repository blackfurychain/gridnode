#!/usr/bin/env bash

set -x

gridnoded tx clp add-liquidity \
  --from $GRID_ACT \
  --keyring-backend test \
  --symbol cusdt \
  --nativeAmount 1000000000000000000000000 \
  --externalAmount 25378853317 \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y