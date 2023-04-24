#!/usr/bin/env bash

set -x

grided tx margin whitelist $(grided keys show tester1 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester2 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester3 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester4 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester5 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester6 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester7 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester8 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester9 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester10 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
grided tx margin whitelist $(grided keys show tester11 --keyring-backend=test -a) \
  --from $ADMIN_KEY \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --node ${GRIDNODE_NODE} \
  --chain-id $GRIDNODE_CHAIN_ID \
  --broadcast-mode block \
  -y
