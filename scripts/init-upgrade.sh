#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean install
gridnoded init test --chain-id=localnet

echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | gridnoded keys add grid --recover

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | gridnoded keys add akasha --recover


gridnoded keys add mkey --multisig grid,akasha --multisig-threshold 2
gridnoded add-genesis-account $(gridnoded keys show grid -a) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink
gridnoded add-genesis-account $(gridnoded keys show akasha -a) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink
gridnoded add-genesis-account $(gridnoded keys show mkey -a) 500000000000000000000000fury

gridnoded add-genesis-clp-admin $(gridnoded keys show grid -a)
gridnoded add-genesis-clp-admin $(gridnoded keys show akasha -a)

gridnoded add-genesis-validators $(gridnoded keys show grid -a --bech val)

gridnoded gentx grid 1000000000000000000000000stake --keyring-backend test

echo "Collecting genesis txs..."
gridnoded collect-gentxs

echo "Validating genesis file..."
gridnoded validate-genesis


mkdir -p $DAEMON_HOME/cosmovisor/genesis/bin
mkdir -p $DAEMON_HOME/cosmovisor/upgrades/release-20210414000000/bin

cp $GOPATH/bin/old/gridnoded $DAEMON_HOME/cosmovisor/genesis/bin
cp $GOPATH/bin/gridnoded $DAEMON_HOME/cosmovisor/upgrades/release-20210414000000/bin/

#contents="$(jq '.gov.voting_params.voting_period = 10' $DAEMON_HOME/config/genesis.json)" && \
#echo "${contents}" > $DAEMON_HOME/config/genesis.json
