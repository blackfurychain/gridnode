#!/bin/sh

set -x

gridnoded tx clp reward-params \
  --lockPeriod 0 \
  --cancelPeriod 0 \
  --from $GRID_ACT \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y