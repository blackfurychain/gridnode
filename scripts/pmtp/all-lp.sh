#!/usr/bin/env bash

set -x

grided q clp all-lp \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID