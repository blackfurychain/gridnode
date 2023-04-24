#!/usr/bin/env bash

set -x

grided tx margin whitelist $ADMIN_ADDRESS \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y