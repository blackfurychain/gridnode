#!/usr/bin/env bash

set -x

grided q gov proposals \
    --node ${GRIDNODE_NODE} \
    --chain-id $GRIDNODE_CHAIN_ID