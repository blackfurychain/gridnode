#!/usr/bin/env bash

set -x

grided q bank balances $ADMIN_ADDRESS \
    --node ${GRIDNODE_NODE} \
    --chain-id $GRIDNODE_CHAIN_ID