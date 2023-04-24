#!/bin/bash

#
# Gridnode entrypoint.
#

set -x

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

. $(SCRIPT_DIR)/vagrantenv.sh
. ${TEST_INTEGRATION_DIR}/shell_utilities.sh

# need a new account to be the administrator
adminuser=$(yes | gridnoded keys add gridnodeadmin --keyring-backend test --output json 2>&1 | jq -r .address)
#{"name":"fnord","type":"local","address":"grid10ckfjtdmk9zkcs9fhl0h260xsj6kvg7esmyqrw","pubkey":"gridpub1addwnpepqtd7ysjyu9aynhemqe9sanmlest8y6dvg24aqzknfmp2ppp7cmxlkc7y8lz","mnemonic":"exact below syrup slender party witness already lamp inform dash impose ginger sauce shift tag humble awkward spawn blue flower lab census gold girl"}
set_persistant_env_var GRIDCHAIN_ADMIN_ACCOUNT $adminuser $envexportfile
gridnoded add-genesis-account $adminuser 100000000000000000000fury --home $CHAINDIR/.gridnoded
gridnoded set-genesis-oracle-admin $adminuser --home $CHAINDIR/.gridnoded
gridnoded set-genesis-whitelister-admin $adminuser --home $CHAINDIR/.gridnoded
gridnoded set-gen-denom-whitelist $SCRIPT_DIR/whitelisted-denoms.json --home $CHAINDIR/.gridnoded
gridnoded start --minimum-gas-prices 0.5fury --rpc.laddr tcp://0.0.0.0:26657 --home $CHAINDIR/.gridnoded
