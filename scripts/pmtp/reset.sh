#!/usr/bin/env bash

set -x

gridnoded tx clp pmtp-rates \
  --blockRate=0.00 \
  --runningRate=0.00 \
  --from=$GRID_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id=$GRIDNODE_CHAIN_ID \
  --broadcast-mode=block \
  -y