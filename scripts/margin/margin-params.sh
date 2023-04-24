#!/usr/bin/env bash

set -x

gridnoded q margin params \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID