#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean install
grided init test --chain-id=localnet

echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | grided keys add grid --recover

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | grided keys add akasha --recover


grided keys add mkey --multisig grid,akasha --multisig-threshold 2
grided add-genesis-account $(grided keys show grid -a) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink
grided add-genesis-account $(grided keys show akasha -a) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink
grided add-genesis-account $(grided keys show mkey -a) 500000000000000000000000fury

grided add-genesis-clp-admin $(grided keys show grid -a)
grided add-genesis-clp-admin $(grided keys show akasha -a)

grided add-genesis-validators $(grided keys show grid -a --bech val)

grided gentx grid 1000000000000000000000000stake --keyring-backend test

echo "Collecting genesis txs..."
grided collect-gentxs

echo "Validating genesis file..."
grided validate-genesis


mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/release-20210414000000/bin

cp $GOPATH/bin/old/grided $DAEMON_HOME/cosmovisor/genesis/bin
cp $GOPATH/bin/grided $DAEMON_HOME/cosmovisor/upgrades/release-20210414000000/bin/

#contents="$(jq '.gov.voting_params.voting_period = 10' $DAEMON_HOME/config/genesis.json)" && \
#echo "${contents}" > $DAEMON_HOME/config/genesis.json
