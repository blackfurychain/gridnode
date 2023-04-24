#!/bin/sh

set -x

gridnoded tx clp add-liquidity \
  --externalAmount 488436982990 \
  --nativeAmount 96176925423929435353999282 \
  --symbol ceth \
  --from $GRID_ACT \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y