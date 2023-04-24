#!/usr/bin/env bash

### chain init script for development purposes only ###

cd ../..
make clean install
rm -rf ~/.gridnoded
gridnoded init test --chain-id=localnet -o

echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | gridnoded keys add grid --recover --keyring-backend=test

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | gridnoded keys add akasha --recover --keyring-backend=test


gridnoded keys add mkey --multisig grid,akasha --multisig-threshold 2 --keyring-backend=test

gridnoded add-genesis-account $(gridnoded keys show grid -a --keyring-backend=test) "999000000000000000000000000000000fury,999000000000000000000000000000000stake,999000000000000000000000000000000cusdc,999000000000000000000000000000000ceth,999000000000000000000000000000000cwbtc,999000000000000000000000000000000ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2" --keyring-backend=test
gridnoded add-genesis-account $(gridnoded keys show akasha -a --keyring-backend=test) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test

gridnoded add-genesis-clp-admin $(gridnoded keys show grid -a --keyring-backend=test) --keyring-backend=test
gridnoded add-genesis-clp-admin $(gridnoded keys show akasha -a --keyring-backend=test) --keyring-backend=test

gridnoded set-genesis-oracle-admin grid --keyring-backend=test
gridnoded add-genesis-validators $(gridnoded keys show grid -a --bech val --keyring-backend=test) --keyring-backend=test

gridnoded set-genesis-whitelister-admin grid --keyring-backend=test
# gridnoded set-gen-denom-whitelist scripts/denoms.json

gridnoded gentx grid 1000000000000000000000000stake --chain-id=localnet --keyring-backend=test

echo "Collecting genesis txs..."
gridnoded collect-gentxs

echo "Validating genesis file..."
gridnoded validate-genesis
