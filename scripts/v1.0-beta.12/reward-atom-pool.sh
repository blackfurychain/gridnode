#!/usr/bin/env bash

gridnoded tx clp reward-period --path=./data/atom_rewards_fix.json \
	--from $ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $GRIDCHAIN_ID \
	--node $GRIDNODE \
	--broadcast-mode block \
	--yes