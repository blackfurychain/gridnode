#!/usr/bin/env bash

# Use gridnoded q account $(gridnoded keys show grid -a) to get seq
seq=1
gridnoded tx dispensation create Airdrop output.json --gas 90128 --from $(gridnoded keys show grid -a) --yes --broadcast-mode async --sequence $seq --account-number 3 --chain-id localnet
seq=$((seq+1))
gridnoded tx dispensation create ValidatorSubsidy output.json --gas 90128 --from $(gridnoded keys show grid -a) --yes --broadcast-mode async --sequence $seq --account-number 3 --chain-id localnet
seq=$((seq+1))
gridnoded tx dispensation create ValidatorSubsidy output.json --gas 90128 --from $(gridnoded keys show grid -a) --yes --broadcast-mode async --sequence $seq --account-number 3 --chain-id localnet