#!/usr/bin/env bash

pkill grided
sleep 5
grided export --height -1 > exported_state.json
sleep 1
grided migrate v0.38 exported_state.json --chain-id new-chain > new-genesis.json  2>&1
sleep 1
grided unsafe-reset-all
sleep 1
cp new-genesis.json ~/.grided/config/genesis.json
sleep 2
grided start