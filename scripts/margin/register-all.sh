#!/usr/bin/env bash

set -x

# grided tx tokenregistry register denoms/stake.json \
#   --node ${GRIDNODE_NODE} \
#   --chain-id "${GRIDNODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block

grided tx tokenregistry register denoms/fury.json \
  --node ${GRIDNODE_NODE} \
  --chain-id "${GRIDNODE_CHAIN_ID}" \
  --from "${ADMIN_ADDRESS}" \
  --keyring-backend test \
  --gas 500000 \
  --gas-prices 0.5fury \
  -y \
  --broadcast-mode block

grided tx tokenregistry register denoms/cusdc.json \
  --node ${GRIDNODE_NODE} \
  --chain-id "${GRIDNODE_CHAIN_ID}" \
  --from "${ADMIN_ADDRESS}" \
  --keyring-backend test \
  --gas 500000 \
  --gas-prices 0.5fury \
  -y \
  --broadcast-mode block

# grided tx tokenregistry register denoms/ceth.json \
#   --node ${GRIDNODE_NODE} \
#   --chain-id "${GRIDNODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block

# grided tx tokenregistry register denoms/cwbtc.json \
#   --node ${GRIDNODE_NODE} \
#   --chain-id "${GRIDNODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block

# grided tx tokenregistry register denoms/uatom.json \
#   --node ${GRIDNODE_NODE} \
#   --chain-id "${GRIDNODE_CHAIN_ID}" \
#   --from "${ADMIN_ADDRESS}" \
#   --keyring-backend test \
#   --gas 500000 \
#   --gas-prices 0.5fury \
#   -y \
#   --broadcast-mode block