#!/usr/bin/env bash

set -x

gridnoded q clp pmtp-params \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID