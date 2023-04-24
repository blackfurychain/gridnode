#!/bin/zsh

proposal_count=1
gridnoded q gov proposals --node tcp://127.0.0.1:27665 >/dev/null 2>&1
if [ $? -eq 0 ]
then
  proposal_count=3
fi

#echo "vote for $proposal_count on localnet-1"
gridnoded tx gov submit-proposal update-client 07-tendermint-0 07-tendermint-2 --from grid --keyring-backend test --home ~/.gridnode-1  --node tcp://127.0.0.1:27665 --title "vote for $proposal_count" --description "vote for $proposal_count from localnet-2" --chain-id localnet-1  --deposit 100000000stake --broadcast-mode block --yes 
echo "proposal made"
gridnoded tx gov vote $proposal_count yes --chain-id localnet-1 --from grid --keyring-backend test --home ~/.gridnode-1 --node tcp://127.0.0.1:27665 --yes --broadcast-mode block 
echo "sleeping"
sleep 30
gridnoded tx gov vote $proposal_count yes --chain-id localnet-1 --from akasha --keyring-backend test --home ~/.gridnode-1 --node tcp://127.0.0.1:27665 --yes --broadcast-mode block 
echo "done"

echo "vote for $((proposal_count+1)) on localnet-1"
gridnoded tx gov submit-proposal update-client 07-tendermint-1 07-tendermint-3 --from grid --keyring-backend test --home ~/.gridnode-1  --node tcp://127.0.0.1:27665 --title "vote for $((proposal_count+1))" --description "vote for $((proposal_count+1)) from localnet-3" --chain-id localnet-1  --deposit 100000000stake --broadcast-mode block --yes 
echo "proposal made"
gridnoded tx gov vote $((proposal_count+1)) yes --chain-id localnet-1 --from grid --keyring-backend test --home ~/.gridnode-1 --node tcp://127.0.0.1:27665 --yes --broadcast-mode block 
echo "sleeping"
sleep 30
gridnoded tx gov vote $((proposal_count+1)) yes --chain-id localnet-1 --from akasha --keyring-backend test --home ~/.gridnode-1 --node tcp://127.0.0.1:27665 --yes --broadcast-mode block 
echo "done"
