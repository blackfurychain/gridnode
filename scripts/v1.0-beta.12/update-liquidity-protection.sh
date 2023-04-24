#!/usr/bin/env bash

grided tx clp liquidity-protection-params --isActive=true \
	--maxFuryLiquidityThreshold=43815115800 \
  --maxFuryLiquidityThresholdAsset=cusdc \
  --epochLength=14400 \
	--from $ADMIN_KEY \
	--gas=500000 \
	--gas-prices=0.5fury \
	--chain-id $GRIDCHAIN_ID \
	--node $GRIDNODE \
	--broadcast-mode block \
	--yes