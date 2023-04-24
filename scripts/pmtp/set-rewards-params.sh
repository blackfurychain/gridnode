#!/usr/bin/env bash

set -x

grided tx clp reward-params \
  --cancelPeriod 43200 \
  --lockPeriod 100800 \
  --from=$GRID_ACT \
  --keyring-backend=test \
  --fees 100000000000000000fury \
  --gas 500000 \
  --node ${GRIDNODE_NODE} \
  --chain-id=$GRIDNODE_CHAIN_ID \
  --broadcast-mode=block \
  -y

# grided tx clp reward-params \
#   --cancelPeriod 66825 \
#   --lockPeriod 124425 \
#   --from=$GRID_ACT \
#   --keyring-backend=test \
#   --fees 100000000000000000fury \
#   --gas 500000 \
#   --node ${GRIDNODE_NODE} \
#   --chain-id=$GRIDNODE_CHAIN_ID \
#   --broadcast-mode=block \
#   -y

# grided tx clp reward-params \
#   --cancelPeriod 66825 \
#   --lockPeriod 100800 \
#   --from=$GRID_ACT \
#   --keyring-backend=test \
#   --fees 100000000000000000fury \
#   --gas 500000 \
#   --node ${GRIDNODE_NODE} \
#   --chain-id=$GRIDNODE_CHAIN_ID \
#   --broadcast-mode=block \
#   -y