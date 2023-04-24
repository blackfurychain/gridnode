#!/usr/bin/env bash

gridnoded tx margin update-pools ./data/temp_pools.json \
	--closed-pools ./data/closed_pools.json \
  --from=$ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $GRIDCHAIN_ID \
	--node $GRIDNODE \
	--broadcast-mode block \
	--yes

gridnoded tx margin whitelist grid1mwmrarhynjuau437d07p42803rntfxqjun3pfu \
  --from=$ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $GRIDCHAIN_ID \
	--node $GRIDNODE \
	--broadcast-mode block \
	--yes