#!/usr/bin/env bash

set -x

gridnoded tx margin update-pools ./pools.json \
  --closed-pools ./closed-pools.json \
  --from=$GRID_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --gas 500000 \
  --node ${GRIDNODE_NODE} \
  --chain-id=$GRIDNODE_CHAIN_ID \
  --broadcast-mode=block \
  -y