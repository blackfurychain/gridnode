#!/usr/bin/env bash

set -x

gridnoded q clp all-lp \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID