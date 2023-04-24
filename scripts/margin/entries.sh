#!/usr/bin/env bash

set -x

grided q tokenregistry entries \
    --node ${GRIDNODE_NODE} \
    --chain-id $GRIDNODE_CHAIN_ID | jq