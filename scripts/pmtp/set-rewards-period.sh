#!/usr/bin/env bash

set -x

grided tx clp reward-period \
  --path ./rewards.json \
  --from=$GRID_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --gas 500000 \
  --node ${GRIDNODE_NODE} \
  --chain-id=$GRIDNODE_CHAIN_ID \
  --broadcast-mode=block \
  -y