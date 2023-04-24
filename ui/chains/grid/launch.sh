#!/usr/bin/env bash

. ../credentials.sh

rm -rf ~/.gridnoded

gridnoded init test --chain-id=gridchain-local
cp ./app.toml ~/.gridnoded/config

echo "Generating deterministic account - ${SHADOWFIEND_NAME}"
echo "${SHADOWFIEND_MNEMONIC}" | gridnoded keys add ${SHADOWFIEND_NAME}  --keyring-backend=test --recover

echo "Generating deterministic account - ${AKASHA_NAME}"
echo "${AKASHA_MNEMONIC}" | gridnoded keys add ${AKASHA_NAME}  --keyring-backend=test --recover

echo "Generating deterministic account - ${JUNIPER_NAME}"
echo "${JUNIPER_MNEMONIC}" | gridnoded keys add ${JUNIPER_NAME} --keyring-backend=test --recover

gridnoded add-genesis-account $(gridnoded keys show ${SHADOWFIEND_NAME} -a --keyring-backend=test) 100000000000000000000000000000fury,100000000000000000000000000000catk,100000000000000000000000000000cbtk,100000000000000000000000000000ceth,100000000000000000000000000000cusdc,100000000000000000000000000000clink,100000000000000000000000000stake
gridnoded add-genesis-account $(gridnoded keys show ${AKASHA_NAME} -a --keyring-backend=test) 100000000000000000000000000000fury,100000000000000000000000000000catk,100000000000000000000000000000cbtk,100000000000000000000000000000ceth,100000000000000000000000000000cusdc,100000000000000000000000000000clink,100000000000000000000000000stake
gridnoded add-genesis-account $(gridnoded keys show ${JUNIPER_NAME} -a --keyring-backend=test) 10000000000000000000000fury,10000000000000000000000cusdc,100000000000000000000clink,100000000000000000000ceth

gridnoded add-genesis-validators $(gridnoded keys show ${SHADOWFIEND_NAME} -a --bech val --keyring-backend=test)

gridnoded gentx ${SHADOWFIEND_NAME} 1000000000000000000000000stake --chain-id=gridchain-local --keyring-backend test

echo "Collecting genesis txs..."
gridnoded collect-gentxs

echo "Validating genesis file..."
gridnoded validate-genesis

echo "Starting test chain"

./start.sh
