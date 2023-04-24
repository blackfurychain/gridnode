#!/bin/sh

# sh ./deregister-one.sh testnet ixo

. ./envs/$1.sh 

TOKEN_REGISTRY_ADMIN_ADDRESS="grid1tpypxpppcf5lea47vcvgy09675nllmcucxydvu"

gridnoded tx tokenregistry deregister $2 \
  --node $GRID_NODE \
  --chain-id $GRIDCHAIN_ID \
  --from $TOKEN_REGISTRY_ADMIN_ADDRESS \
  --keyring-backend $KEYRING_BACKEND \
  --gas=500000 \
  --gas-prices=0.5fury \
  -y