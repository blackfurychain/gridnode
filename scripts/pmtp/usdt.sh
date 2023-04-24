#!/usr/bin/env bash

set -x

grided q clp pool cusdt \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID