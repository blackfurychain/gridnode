#!/bin/sh

# sh ./deregister-all.sh testnet

. ./envs/$1.sh 

mkdir -p ./$GRIDCHAIN_ID
rm -f ./$GRIDCHAIN_ID/temp.json
rm -f ./$GRIDCHAIN_ID/temp2.json
rm -f ./$GRIDCHAIN_ID/tokenregistry.json

gridnoded q tokenregistry add-all ./$GRIDCHAIN_ID/registry.json | jq > $GRIDCHAIN_ID/temp.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/temp.json ./$GRIDCHAIN_ID/cosmos.json | jq > $GRIDCHAIN_ID/tokenregistry.json
rm ./$GRIDCHAIN_ID/temp.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/tokenregistry.json ./$GRIDCHAIN_ID/akash.json | jq > $GRIDCHAIN_ID/temp.json
rm ./$GRIDCHAIN_ID/tokenregistry.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/temp.json ./$GRIDCHAIN_ID/sentinel.json | jq > $GRIDCHAIN_ID/tokenregistry.json
rm ./$GRIDCHAIN_ID/temp.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/tokenregistry.json ./$GRIDCHAIN_ID/iris.json | jq > $GRIDCHAIN_ID/temp.json
rm ./$GRIDCHAIN_ID/tokenregistry.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/temp.json ./$GRIDCHAIN_ID/persistence.json | jq > $GRIDCHAIN_ID/tokenregistry.json
rm ./$GRIDCHAIN_ID/temp.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/tokenregistry.json ./$GRIDCHAIN_ID/crypto-org.json | jq > $GRIDCHAIN_ID/temp.json
rm ./$GRIDCHAIN_ID/tokenregistry.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/temp.json ./$GRIDCHAIN_ID/regen.json | jq > $GRIDCHAIN_ID/tokenregistry.json
rm ./$GRIDCHAIN_ID/temp.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/tokenregistry.json ./$GRIDCHAIN_ID/terra.json | jq > $GRIDCHAIN_ID/temp.json
rm ./$GRIDCHAIN_ID/tokenregistry.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/temp.json ./$GRIDCHAIN_ID/osmosis.json | jq > $GRIDCHAIN_ID/tokenregistry.json
rm ./$GRIDCHAIN_ID/temp.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/tokenregistry.json ./$GRIDCHAIN_ID/juno.json | jq > $GRIDCHAIN_ID/temp.json
rm ./$GRIDCHAIN_ID/tokenregistry.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/temp.json ./$GRIDCHAIN_ID/ixo.json | jq > $GRIDCHAIN_ID/tokenregistry.json
rm ./$GRIDCHAIN_ID/temp.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/tokenregistry.json ./$GRIDCHAIN_ID/emoney.json | jq > $GRIDCHAIN_ID/temp.json
rm ./$GRIDCHAIN_ID/tokenregistry.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/temp.json ./$GRIDCHAIN_ID/likecoin.json | jq > $GRIDCHAIN_ID/tokenregistry.json
rm ./$GRIDCHAIN_ID/temp.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/tokenregistry.json ./$GRIDCHAIN_ID/bitsong.json | jq > $GRIDCHAIN_ID/temp.json
rm ./$GRIDCHAIN_ID/tokenregistry.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/temp.json ./$GRIDCHAIN_ID/band.json | jq > $GRIDCHAIN_ID/tokenregistry.json
rm ./$GRIDCHAIN_ID/temp.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/tokenregistry.json ./$GRIDCHAIN_ID/emoney-eeur.json | jq > $GRIDCHAIN_ID/temp.json
rm ./$GRIDCHAIN_ID/tokenregistry.json
gridnoded q tokenregistry add ./$GRIDCHAIN_ID/temp.json ./$GRIDCHAIN_ID/terra-uusd.json | jq > $GRIDCHAIN_ID/tokenregistry.json
rm ./$GRIDCHAIN_ID/temp.json