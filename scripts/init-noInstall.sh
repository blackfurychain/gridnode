#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean-config
grided init test --chain-id=localnet

echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | grided keys add grid --recover

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | grided keys add akasha --recover

grided add-genesis-account $(grided keys show grid -a) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink
grided add-genesis-account $(grided keys show akasha -a) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink

grided add-genesis-clp-admin $(grided keys show grid -a)
grided add-genesis-clp-admin $(grided keys show akasha -a)

grided add-genesis-validators $(grided keys show grid -a --bech val)

grided gentx grid 1000000000000000000000000stake --keyring-backend test

echo "Collecting genesis txs..."
grided collect-gentxs

echo "Validating genesis file..."
grided validate-genesis
