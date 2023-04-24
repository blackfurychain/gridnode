#!/usr/bin/env bash

set -x

gridnoded tx margin close \
  --from $GRID_ACT \
  --id 7 \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y