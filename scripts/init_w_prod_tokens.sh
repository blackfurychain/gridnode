#!/usr/bin/env bash

### chain init script for development purposes only ###

make clean install
rm -rf ~/.grided
grided init test --chain-id=localnet -o

echo "Generating deterministic account - grid"
echo "race draft rival universe maid cheese steel logic crowd fork comic easy truth drift tomorrow eye buddy head time cash swing swift midnight borrow" | grided keys add grid --recover --keyring-backend=test

echo "Generating deterministic account - akasha"
echo "hand inmate canvas head lunar naive increase recycle dog ecology inhale december wide bubble hockey dice worth gravity ketchup feed balance parent secret orchard" | grided keys add akasha --recover --keyring-backend=test


grided keys add mkey --multisig grid,akasha --multisig-threshold 2 --keyring-backend=test

grided add-genesis-account $(grided keys show grid -a --keyring-backend=test) "999000000000000000000000000000000fury,999000000000000000000000000000000stake,999000000000000000000000000000000ceth,999000000000000000000000000000000cusdc,999000000000000000000000000000000cusdt,999000000000000000000000000000000ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2,999000000000000000000000000000000ibc/F279AB967042CAC10BFF70FAECB179DCE37AAAE4CD4C1BC4565C2BBC383BC0FA,999000000000000000000000000000000ibc/F141935FF02B74BDC6B8A0BD6FE86A23EE25D10E89AA0CD9158B3D92B63FDF4D" --keyring-backend=test
grided add-genesis-account $(grided keys show akasha -a --keyring-backend=test) 500000000000000000000000fury,500000000000000000000000catk,500000000000000000000000cbtk,500000000000000000000000ceth,990000000000000000000000000stake,500000000000000000000000cdash,500000000000000000000000clink --keyring-backend=test

grided add-genesis-clp-admin $(grided keys show grid -a --keyring-backend=test) --keyring-backend=test
grided add-genesis-clp-admin $(grided keys show akasha -a --keyring-backend=test) --keyring-backend=test

grided set-genesis-oracle-admin grid --keyring-backend=test
grided add-genesis-validators $(grided keys show grid -a --bech val --keyring-backend=test) --keyring-backend=test

# FIXME: commented as it overrides admin accounts list in genesis set by default
# grided set-genesis-whitelister-admin grid --keyring-backend=test
# grided set-gen-denom-whitelist scripts/denoms.json

grided gentx grid 1000000000000000000000000stake --chain-id=localnet --keyring-backend=test

echo "Collecting genesis txs..."
grided collect-gentxs

echo "Validating genesis file..."
grided validate-genesis
