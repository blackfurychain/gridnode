#!/usr/bin/env bash

grided tx clp reward-period --path=./data/rp_24.json \
	--from $ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $GRIDCHAIN_ID \
	--node $GRIDNODE \
	--broadcast-mode block \
	--yes