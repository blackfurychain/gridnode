#!/usr/bin/env bash

set -x

gridnoded tx clp add-liquidity \
  --from $GRID_ACT \
  --keyring-backend test \
  --symbol ceth \
  --nativeAmount 0 \
  --externalAmount 455000000000000 \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y