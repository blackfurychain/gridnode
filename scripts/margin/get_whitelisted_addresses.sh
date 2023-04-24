#!/usr/bin/env bash

set -x

gridnoded q margin whitelist \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID