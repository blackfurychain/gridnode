#!/usr/bin/env bash

# Use grided q account $(grided keys show grid -a) to get seq
seq=1
grided tx dispensation create Airdrop output.json --gas 90128 --from $(grided keys show grid -a) --yes --broadcast-mode async --sequence $seq --account-number 3 --chain-id localnet
seq=$((seq+1))
grided tx dispensation create ValidatorSubsidy output.json --gas 90128 --from $(grided keys show grid -a) --yes --broadcast-mode async --sequence $seq --account-number 3 --chain-id localnet
seq=$((seq+1))
grided tx dispensation create ValidatorSubsidy output.json --gas 90128 --from $(grided keys show grid -a) --yes --broadcast-mode async --sequence $seq --account-number 3 --chain-id localnet