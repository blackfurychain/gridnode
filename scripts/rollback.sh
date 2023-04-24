#!/usr/bin/env bash

pkill gridnoded
sleep 5
gridnoded export --height -1 > exported_state.json
sleep 1
gridnoded migrate v0.38 exported_state.json --chain-id new-chain > new-genesis.json  2>&1
sleep 1
gridnoded unsafe-reset-all
sleep 1
cp new-genesis.json ~/.gridnoded/config/genesis.json
sleep 2
gridnoded start