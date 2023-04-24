#!/usr/bin/env bash

set -x

ACCOUNT_NUMBER=$(grided q auth account $ADMIN_ADDRESS \
    --node ${GRIDNODE_NODE} \
    --chain-id $GRIDNODE_CHAIN_ID \
    --output json \
    | jq -r ".account_number")
SEQUENCE=$(grided q auth account $ADMIN_ADDRESS \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --output json \
  | jq -r ".sequence")
for i in {0..12244}; do
  echo "tx ${i}"
  grided tx clp add_liquidity \
    --from=$GRID_ACT \
    --keyring-backend=test \
    --externalAmount=${EXTERNAL_AMOUNT} \
    --nativeAmount=${NATIVE_AMOUNT} \
    --symbol=${SYMBOL} \
    --fees=100000000000000000fury \
    --gas=500000 \
    --node=${GRIDNODE_NODE} \
    --chain-id=${GRIDNODE_CHAIN_ID} \
    --broadcast-mode=async \
    --account-number=${ACCOUNT_NUMBER} \
    --sequence=$(($SEQUENCE + $i)) \
    -y
  done