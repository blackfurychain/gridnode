#!/usr/bin/env bash

. ../credentials.sh

rm -rf ~/.grided

grided init test --chain-id=gridironchain-local
cp ./app.toml ~/.grided/config

echo "Generating deterministic account - ${SHADOWFIEND_NAME}"
echo "${SHADOWFIEND_MNEMONIC}" | grided keys add ${SHADOWFIEND_NAME}  --keyring-backend=test --recover

echo "Generating deterministic account - ${AKASHA_NAME}"
echo "${AKASHA_MNEMONIC}" | grided keys add ${AKASHA_NAME}  --keyring-backend=test --recover

echo "Generating deterministic account - ${JUNIPER_NAME}"
echo "${JUNIPER_MNEMONIC}" | grided keys add ${JUNIPER_NAME} --keyring-backend=test --recover

grided add-genesis-account $(grided keys show ${SHADOWFIEND_NAME} -a --keyring-backend=test) 100000000000000000000000000000fury,100000000000000000000000000000catk,100000000000000000000000000000cbtk,100000000000000000000000000000ceth,100000000000000000000000000000cusdc,100000000000000000000000000000clink,100000000000000000000000000stake
grided add-genesis-account $(grided keys show ${AKASHA_NAME} -a --keyring-backend=test) 100000000000000000000000000000fury,100000000000000000000000000000catk,100000000000000000000000000000cbtk,100000000000000000000000000000ceth,100000000000000000000000000000cusdc,100000000000000000000000000000clink,100000000000000000000000000stake
grided add-genesis-account $(grided keys show ${JUNIPER_NAME} -a --keyring-backend=test) 10000000000000000000000fury,10000000000000000000000cusdc,100000000000000000000clink,100000000000000000000ceth

grided add-genesis-validators $(grided keys show ${SHADOWFIEND_NAME} -a --bech val --keyring-backend=test)

grided gentx ${SHADOWFIEND_NAME} 1000000000000000000000000stake --chain-id=gridironchain-local --keyring-backend test

echo "Collecting genesis txs..."
grided collect-gentxs

echo "Validating genesis file..."
grided validate-genesis

echo "Starting test chain"

./start.sh
