#!/usr/bin/env bash

set -x

grided q clp pmtp-params \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID