#!/usr/bin/env bash

set -x

gridnoded tx gov vote 2 yes \
  --from ${GRID_ACT} \
  --keyring-backend test \
  --chain-id="${GRIDNODE_CHAIN_ID}" \
  --node="${GRIDNODE_NODE}" \
  --fees=100000000000000000fury \
  --broadcast-mode=block \
  -y