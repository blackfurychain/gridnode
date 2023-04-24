#!/bin/sh
. ./envs/$1.sh 

# sh ./register-one.sh testnet ixo


TOKEN_REGISTRY_ADMIN_ADDRESS="grid1tpypxpppcf5lea47vcvgy09675nllmcucxydvu"

gridnoded tx tokenregistry register ./$GRIDCHAIN_ID/$2.json \
  --node $GRID_NODE \
  --chain-id $GRIDCHAIN_ID \
  --from $TOKEN_REGISTRY_ADMIN_ADDRESS \
  --keyring-backend $KEYRING_BACKEND \
  --gas=500000 \
  --gas-prices=0.5fury \
  -y