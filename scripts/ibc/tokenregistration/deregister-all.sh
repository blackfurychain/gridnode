#!/bin/sh

# sh ./deregister-all.sh testnet

. ./envs/$1.sh 

TOKEN_REGISTRY_ADMIN_ADDRESS="did:fury:g1tpypxpppcf5lea47vcvgy09675nllmcucxydvu"

gridnoded tx tokenregistry deregister-all ./$GRIDCHAIN_ID/tokenregistry.json \
  --node $GRID_NODE \
  --chain-id $GRIDCHAIN_ID \
  --from $TOKEN_REGISTRY_ADMIN_ADDRESS \
  --keyring-backend $KEYRING_BACKEND \
  --gas=500000 \
  --gas-prices=0.5fury \
  -y