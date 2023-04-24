#!/usr/bin/env bash

set -x

gridnoded tx clp pmtp-params \
  --pmtp_start=22811 \
  --pmtp_end=224410 \
  --epochLength=14400 \
  --rGov=0.05 \
  --from=$GRID_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id=$GRIDNODE_CHAIN_ID \
  --broadcast-mode=block \
  -y