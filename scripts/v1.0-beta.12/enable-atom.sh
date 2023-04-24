#!/usr/bin/env bash

grided tx tokenregistry register ./data/atom_all_permissions.json \
	--from $ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $GRIDCHAIN_ID \
	--node $GRIDNODE \
	--broadcast-mode block \
	--yes