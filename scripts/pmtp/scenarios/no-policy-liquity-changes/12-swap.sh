#!/usr/bin/env bash

set -x

grided tx clp swap \
  --from $GRID_ACT \
  --keyring-backend test \
  --sentSymbol fury \
  --receivedSymbol cusdt \
  --sentAmount 100000000000000000000000 \
  --minReceivingAmount 0 \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y