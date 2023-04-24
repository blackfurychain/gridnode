#!/usr/bin/env bash

set -x

grided tx margin dewhitelist did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92 \
  --from $GRID_ACT \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y