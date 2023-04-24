#!/usr/bin/env bash

gridnoded tx clp set-lppd-params --path=./data/lpd_params.json \
	--from $ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $GRIDCHAIN_ID \
	--node $GRIDNODE \
	--broadcast-mode block \
	--yes