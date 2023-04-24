#!/bin/sh

set -x

grided tx clp set-lppd-params \
  --path lppd-params.json \
  --from $GRID_ACT \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y