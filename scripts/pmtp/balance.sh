#!/usr/bin/env bash

set -x

gridnoded q bank balances $ADMIN_ADDRESS \
    --node ${GRIDNODE_NODE} \
    --chain-id $GRIDNODE_CHAIN_ID