#!/usr/bin/env bash

set -x

gridnoded q margin \
  positions-for-address $ADMIN_ADDRESS \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID