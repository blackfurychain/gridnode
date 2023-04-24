#!/usr/bin/env bash

set -x

gridnoded q clp params \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID