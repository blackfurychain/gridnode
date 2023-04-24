#!/bin/sh
. ./envs/$1.sh 

# sh ./register-all.sh testnet


TOKEN_REGISTRY_ADMIN_ADDRESS="did:fury:g1tpypxpppcf5lea47vcvgy09675nllmcucxydvu"

gridnoded tx tokenregistry register-all ./$GRIDCHAIN_ID/tokenregistry.json \
  --node $GRID_NODE \
  --chain-id $GRIDCHAIN_ID \
  --from $TOKEN_REGISTRY_ADMIN_ADDRESS \
  --keyring-backend $KEYRING_BACKEND \
  --gas=500000 \
  --gas-prices=0.5fury \
  -y