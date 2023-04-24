#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean install
rm -rf ~/.grided
grided init test --chain-id=localnet -o

echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | grided keys add grid --recover --keyring-backend=test

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | grided keys add akasha --recover --keyring-backend=test

echo "Generating deterministic account - alice"
echo "crunch enable gauge equip sadness venture volcano capable boil pole lounge because service level giggle decide south deposit bike antique consider olympic girl butter" | grided keys add alice --recover --keyring-backend=test

grided keys add mkey --multisig grid,akasha --multisig-threshold 2 --keyring-backend=test

grided add-genesis-account $(grided keys show grid -a --keyring-backend=test) 500000000000000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink,5000000000000cusdt,90000000000000000000ibc/96D7172B711F7F925DFC7579C6CCC3C80B762187215ABD082CDE99F81153DC80 --keyring-backend=test
grided add-genesis-account $(grided keys show akasha -a --keyring-backend=test) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test
grided add-genesis-account $(grided keys show alice -a --keyring-backend=test) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test

grided add-genesis-clp-admin $(grided keys show grid -a --keyring-backend=test) --keyring-backend=test
grided add-genesis-clp-admin $(grided keys show akasha -a --keyring-backend=test) --keyring-backend=test

grided set-genesis-oracle-admin grid --keyring-backend=test
grided add-genesis-validators $(grided keys show grid -a --bech val --keyring-backend=test) --keyring-backend=test

grided set-genesis-whitelister-admin grid --keyring-backend=test
grided set-gen-denom-whitelist scripts/denoms.json

grided gentx grid 1000000000000000000000000stake --moniker grid_val --chain-id=localnet --keyring-backend=test

echo "Collecting genesis txs..."
grided collect-gentxs

echo "Validating genesis file..."
grided validate-genesis
