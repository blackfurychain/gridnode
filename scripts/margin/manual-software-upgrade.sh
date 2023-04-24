#!/usr/bin/env bash

set -x

gridnoded tx gov submit-proposal software-upgrade "${NEW_VERSION}" \
  --from ${GRID_ACT} \
  --deposit "${DEPOSIT}" \
  --upgrade-height "${TARGET_BLOCK}" \
  --title "v${NEW_VERSION}" \
  --description "v${NEW_VERSION}" \
  --chain-id "${GRIDNODE_CHAIN_ID}" \
  --node "${GRIDNODE_NODE}" \
  --keyring-backend "test" \
  --fees 100000000000000000fury \
  --broadcast-mode=block \
  -y