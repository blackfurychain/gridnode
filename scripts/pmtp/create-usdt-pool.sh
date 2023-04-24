#!/usr/bin/env bash

set -x

grided tx clp create-pool \
  --from $GRID_ACT \
  --keyring-backend test \
  --symbol cusdt \
  --nativeAmount 1550459183129248235861408 \
  --externalAmount 174248776094 \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y