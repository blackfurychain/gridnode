#!/usr/bin/env bash

set -x

grided q clp params \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID