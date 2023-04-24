#!/usr/bin/env bash

set -x

gridnoded tx clp set-symmetry-threshold \
  --threshold=0.000000005 \
  --from=$GRID_ACT \
  --keyring-backend=test \
  --fees=100000000000000000fury \
  --gas=500000 \
  --node=${GRIDNODE_NODE} \
  --chain-id=$GRIDNODE_CHAIN_ID \
  --broadcast-mode=block \
  -y