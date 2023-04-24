#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean install
gridnoded init test --chain-id=localnet

echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | gridnoded keys add grid --recover

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | gridnoded keys add akasha --recover

gridnoded add-genesis-account $(gridnoded keys show grid -a) 16205782692902021002506278400fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink,899999867990000000000000000000cacoin
gridnoded add-genesis-account $(gridnoded keys show akasha -a) 5000000000000003407464fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink,8999998679900000000000000000000cacoin

gridnoded add-genesis-clp-admin $(gridnoded keys show grid -a)
gridnoded add-genesis-clp-admin $(gridnoded keys show akasha -a)

gridnoded  add-genesis-validators $(gridnoded keys show grid -a --bech val)

gridnoded gentx grid 1000000000000000000000000stake --keyring-backend test

echo "Collecting genesis txs..."
gridnoded collect-gentxs

echo "Validating genesis file..."
gridnoded validate-genesis
