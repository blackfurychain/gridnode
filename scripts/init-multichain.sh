#!/usr/bin/env bash

if ! type "hermes" > /dev/null; then
  # install foobar here
  echo "You need the hermes relayer to run this script. You can find it here https://github.com/informalsystems/ibc-rs"
  exit 0
fi

### chain init script for development purposes only ###
killall gridnoded
killall hermes
rm -rf ~/.gridnode-1
rm -rf ~/.gridnode-2
rm -rf ~/.gridnode-3
make clean install
gridnoded init test --chain-id=localnet-1 -o --home ~/.gridnode-1

echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | gridnoded keys add grid --recover --keyring-backend=test --home ~/.gridnode-1

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | gridnoded keys add akasha --recover --keyring-backend=test --home ~/.gridnode-1


gridnoded keys add mkey --multisig grid,akasha --multisig-threshold 2 --keyring-backend=test --home ~/.gridnode-1

gridnoded add-genesis-account $(gridnoded keys show grid -a --keyring-backend=test --home ~/.gridnode-1) 50000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test --home ~/.gridnode-1
gridnoded add-genesis-account $(gridnoded keys show akasha -a --keyring-backend=test --home ~/.gridnode-1) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test --home ~/.gridnode-1

gridnoded add-genesis-clp-admin $(gridnoded keys show grid -a --keyring-backend=test --home ~/.gridnode-1) --keyring-backend=test --home ~/.gridnode-1
gridnoded add-genesis-clp-admin $(gridnoded keys show akasha -a --keyring-backend=test --home ~/.gridnode-1 ) --keyring-backend=test --home ~/.gridnode-1
gridnoded set-genesis-whitelister-admin $(gridnoded keys show grid -a --keyring-backend=test --home ~/.gridnode-1) --keyring-backend=test --home ~/.gridnode-1
gridnoded set-gen-denom-whitelist scripts/denoms.json --home ~/.gridnode-1

gridnoded add-genesis-validators $(gridnoded keys show grid -a --bech val --keyring-backend=test --home ~/.gridnode-1) --keyring-backend=test --home ~/.gridnode-1

gridnoded gentx grid 1000000000000000000000000stake --keyring-backend=test --home ~/.gridnode-1 --chain-id=localnet-1

echo "Collecting genesis txs..."
gridnoded collect-gentxs --home ~/.gridnode-1

echo "Validating genesis file..."
gridnoded validate-genesis --home ~/.gridnode-1



gridnoded init test --chain-id=localnet-2 -o --home ~/.gridnode-2


echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | gridnoded keys add grid --recover --keyring-backend=test --home ~/.gridnode-2

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | gridnoded keys add akasha --recover --keyring-backend=test --home ~/.gridnode-2


gridnoded keys add mkey --multisig grid,akasha --multisig-threshold 2 --keyring-backend=test --home ~/.gridnode-2

gridnoded add-genesis-account $(gridnoded keys show grid -a --keyring-backend=test --home ~/.gridnode-2 ) 50000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test --home ~/.gridnode-2
gridnoded add-genesis-account $(gridnoded keys show akasha -a --keyring-backend=test --home ~/.gridnode-2) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test --home ~/.gridnode-2

gridnoded add-genesis-clp-admin $(gridnoded keys show grid -a --keyring-backend=test --home ~/.gridnode-2 ) --keyring-backend=test --home ~/.gridnode-2
gridnoded add-genesis-clp-admin $(gridnoded keys show akasha -a --keyring-backend=test --home ~/.gridnode-2) --keyring-backend=test --home ~/.gridnode-2
gridnoded set-genesis-whitelister-admin $(gridnoded keys show grid -a --keyring-backend=test --home ~/.gridnode-2) --keyring-backend=test --home ~/.gridnode-2
gridnoded set-gen-denom-whitelist scripts/denoms.json --home ~/.gridnode-2
gridnoded add-genesis-validators $(gridnoded keys show grid -a --bech val --keyring-backend=test --home ~/.gridnode-2 ) --keyring-backend=test --home ~/.gridnode-2

gridnoded gentx grid 1000000000000000000000000stake --chain-id=localnet --keyring-backend=test --home ~/.gridnode-2 --chain-id=localnet-2


echo "Collecting genesis txs..."
gridnoded collect-gentxs --home ~/.gridnode-2

echo "Validating genesis file..."
gridnoded validate-genesis --home ~/.gridnode-2



gridnoded init test --chain-id=localnet-3 -o --home ~/.gridnode-3


echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | gridnoded keys add grid --recover --keyring-backend=test --home ~/.gridnode-3

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | gridnoded keys add akasha --recover --keyring-backend=test --home ~/.gridnode-3


gridnoded keys add mkey --multisig grid,akasha --multisig-threshold 2 --keyring-backend=test --home ~/.gridnode-3

gridnoded add-genesis-account $(gridnoded keys show grid -a --keyring-backend=test --home ~/.gridnode-3 ) 50000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test --home ~/.gridnode-3
gridnoded add-genesis-account $(gridnoded keys show akasha -a --keyring-backend=test --home ~/.gridnode-3) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test --home ~/.gridnode-3

gridnoded add-genesis-clp-admin $(gridnoded keys show grid -a --keyring-backend=test --home ~/.gridnode-3 ) --keyring-backend=test --home ~/.gridnode-3
gridnoded add-genesis-clp-admin $(gridnoded keys show akasha -a --keyring-backend=test --home ~/.gridnode-3) --keyring-backend=test --home ~/.gridnode-3
gridnoded set-genesis-whitelister-admin $(gridnoded keys show grid -a --keyring-backend=test --home ~/.gridnode-3) --keyring-backend=test --home ~/.gridnode-3
gridnoded set-gen-denom-whitelist scripts/denoms.json --home ~/.gridnode-3
gridnoded add-genesis-validators $(gridnoded keys show grid -a --bech val --keyring-backend=test --home ~/.gridnode-3 ) --keyring-backend=test --home ~/.gridnode-3

gridnoded gentx grid 1000000000000000000000000stake --chain-id=localnet-3 --keyring-backend=test --home ~/.gridnode-3 --chain-id=localnet-3

echo "Collecting genesis txs..."
gridnoded collect-gentxs --home ~/.gridnode-3

echo "Validating genesis file..."
gridnoded validate-genesis --home ~/.gridnode-3

rm -rf abci_*.log
rm -rf hermes.log
rm -rf ~/.hermes

echo "Chainging voting period to 60 seconds"
sed -i -s 's/        "voting_period": "172800s"/        "voting_period": "60s"/g' ~/.gridnode-1/config/genesis.json
sed -i -s 's/        "voting_period": "172800s"/        "voting_period": "60s"/g' ~/.gridnode-2/config/genesis.json
sed -i -s 's/        "voting_period": "172800s"/        "voting_period": "60s"/g' ~/.gridnode-3/config/genesis.json

echo "Starting gridnoded's"

sleep 1
gridnoded start --home ~/.gridnode-1 --p2p.laddr 0.0.0.0:27655  --grpc.address 0.0.0.0:9090 --grpc-web.address 0.0.0.0:9093 --address tcp://0.0.0.0:27659 --rpc.laddr tcp://127.0.0.1:27665 >> abci_1.log 2>&1  &
sleep 1
gridnoded start --home ~/.gridnode-2 --p2p.laddr 0.0.0.0:27656  --grpc.address 0.0.0.0:9091 --grpc-web.address 0.0.0.0:9094 --address tcp://0.0.0.0:27660 --rpc.laddr tcp://127.0.0.1:27666 >> abci_2.log 2>&1  &
sleep 1
gridnoded start --home ~/.gridnode-3 --p2p.laddr 0.0.0.0:27657  --grpc.address 0.0.0.0:9092 --grpc-web.address 0.0.0.0:9095 --address tcp://0.0.0.0:27661 --rpc.laddr tcp://127.0.0.1:27667 >> abci_3.log 2>&1 &
sleep 10

echo "updating token registries with IBC paths"
echo "doing localnet-1"
gridnoded tx tokenregistry register scripts/fury-localnet-1-localnet-2.json --node tcp://127.0.0.1:27665 --keyring-backend test --chain-id localnet-1 --from grid --gas 200000 --gas-prices 0.5fury --home ~/.gridnode-1 --yes
sleep 10
gridnoded tx tokenregistry register scripts/fury-localnet-1-localnet-3.json --node tcp://127.0.0.1:27665 --keyring-backend test --chain-id localnet-1 --from grid --gas 200000 --gas-prices 0.5fury --home ~/.gridnode-1 --yes
echo ""
sleep 10

echo "Doing localnet-2"
gridnoded tx tokenregistry register scripts/fury-localnet-2-localnet-1.json --node tcp://127.0.0.1:27666 --keyring-backend test --chain-id localnet-2 --from grid --gas 200000 --gas-prices 0.5fury --home ~/.gridnode-2 --yes
sleep 10
gridnoded tx tokenregistry register scripts/fury-localnet-2-localnet-3.json --node tcp://127.0.0.1:27666 --keyring-backend test --chain-id localnet-2 --from grid --gas 200000 --gas-prices 0.5fury --home ~/.gridnode-2 --yes
echo ""
sleep 10

echo "Doing localnet-3"
gridnoded tx tokenregistry register scripts/fury-localnet-3-localnet-1.json --node tcp://127.0.0.1:27667 --keyring-backend test --chain-id localnet-3 --from grid --gas 200000 --gas-prices 0.5fury --home ~/.gridnode-3 --yes
sleep 10
gridnoded tx tokenregistry register scripts/fury-localnet-3-localnet-2.json --node tcp://127.0.0.1:27667 --keyring-backend test --chain-id localnet-3 --from grid --gas 200000 --gas-prices 0.5fury --home ~/.gridnode-3 --yes
echo ""

sleep 10

echo "Setting hermes"
# copy hermes config to the hermes directory
mkdir ~/.hermes
cp scripts/hermes_config.toml ~/.hermes/config.toml

hermes keys restore -m "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" localnet-1 --name grid
hermes keys restore -m "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" localnet-2 --name grid
hermes keys restore -m "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" localnet-3 --name grid

# create hermes channels
echo "Creating localnet-1 to localnet-2"
hermes create channel localnet-1 localnet-2 --port-a transfer --port-b transfer -o unordered
sleep 1
echo "Creating localnet-2 to localnet-3"
hermes create channel localnet-2 localnet-3 --port-a transfer --port-b transfer -o unordered
sleep 1
echo "Creating localnet-1 to localnet-3"
hermes create channel localnet-1 localnet-3 --port-a transfer --port-b transfer -o unordered
sleep 1

# start hermes
hermes start > hermes.log 2>&1 &

echo "Sleeping to let hermes boot"
sleep 10
