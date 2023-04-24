#!/bin/zsh

echo "vote for $1"
gridnoded tx gov submit-proposal update-client 07-tendermint-1 07-tendermint-0 --from grid --keyring-backend test --home ~/.gridnode-3  --node tcp://127.0.0.1:27667 --title "vote for $1" --description "vote for $1" --chain-id localnet-3  --deposit 100000000stake --broadcast-mode block --yes 
echo "proposal made"
gridnoded tx gov vote $1 yes --chain-id localnet-3 --from grid --keyring-backend test --home ~/.gridnode-3 --node tcp://127.0.0.1:27667 --yes --broadcast-mode block 
echo "sleeping"
sleep 30
gridnoded tx gov vote $1 yes --chain-id localnet-3 --from akasha --keyring-backend test --home ~/.gridnode-3 --node tcp://127.0.0.1:27667 --yes --broadcast-mode block 
echo "done"
