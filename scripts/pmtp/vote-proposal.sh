#!/usr/bin/env bash

set -x

grided tx gov vote 1 yes \
    --from $GRID_ACT \
    --keyring-backend test \
    --node ${GRIDNODE_NODE} \
    --chain-id $GRIDNODE_CHAIN_ID \
    --fees 100000000000000000fury \
    --broadcast-mode block \
    --trace \
    -y