#!/usr/bin/env bash

set -x

gridnoded tx margin dewhitelist grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd \
  --from $GRID_ACT \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y