#!/usr/bin/env bash

set -x

echo "$(jq '.app_state.gov.voting_params.voting_period = "60s"' $HOME/.gridnoded/config/genesis.json)" > $HOME/.gridnoded/config/genesis.json