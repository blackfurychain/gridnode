#!/usr/bin/env bash

set -x

gridnoded q clp pools \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID