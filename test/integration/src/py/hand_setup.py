import logging
import os
import sys

import burn_lock_functions
import test_utilities
from burn_lock_functions import EthereumToGridironchainTransferRequest
from integration_env_credentials import gridchain_cli_credentials_for_test
from test_utilities import get_required_env_var, GridironchaincliCredentials, get_optional_env_var, \
    ganache_owner_account

logging.basicConfig(
    level=logging.DEBUG,
    format="%(asctime)s [%(levelname)s] %(message)s",
    handlers=[logging.StreamHandler(sys.stdout)]
)

logging.debug("starting")

smart_contracts_dir = get_required_env_var("SMART_CONTRACTS_DIR")

ethereum_address = get_optional_env_var(
    "ETHEREUM_ADDRESS",
    ganache_owner_account(smart_contracts_dir)
)


def build_request() -> (EthereumToGridironchainTransferRequest, GridironchaincliCredentials):
    new_account_key = 'user1'
    credentials = gridchain_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_gridaddr(credentials=credentials, keyname=new_account_key)
    credentials.from_key = new_addr["name"]
    request = EthereumToGridironchainTransferRequest(
        gridchain_address=new_addr["address"],
        smart_contracts_dir=smart_contracts_dir,
        ethereum_address=ethereum_address,
        ethereum_private_key_env_var="ETHEREUM_PRIVATE_KEY",
        bridgebank_address=get_required_env_var("BRIDGE_BANK_ADDRESS"),
        ethereum_network=(os.environ.get("ETHEREUM_NETWORK") or ""),
        amount=9 * 10 ** 18,
        ceth_amount=2 * (10 ** 16)
    )
    return request, credentials


# if there's an existing user1 key, just remove it.  Otherwise, adding a duplicate key will just hang
try:
    test_utilities.get_shell_output(f"gridnoded keys delete user1 --home /home/vagrant/.gridnoded --keyring-backend test -o json")
except:
    logging.debug("no key to delete, this is normal in a fresh environment")
request, credentials = build_request()
burn_lock_functions.transfer_ethereum_to_gridchain(request)
test_utilities.get_gridchain_addr_balance(request.gridchain_address, request.gridnoded_node, "ceth")
logging.info(f"created account for key {credentials.from_key}")
