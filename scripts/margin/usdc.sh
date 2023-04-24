#!/usr/bin/env bash

set -x

grided q clp pool cusdc \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID