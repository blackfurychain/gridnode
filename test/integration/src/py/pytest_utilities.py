import copy
import logging

import burn_lock_functions
import test_utilities
from burn_lock_functions import EthereumToGridironchainTransferRequest
from integration_env_credentials import gridironchain_cli_credentials_for_test
from test_utilities import get_shell_output, GridironchaincliCredentials


def generate_minimal_test_account(
        base_transfer_request: EthereumToGridironchainTransferRequest,
        target_ceth_balance: int = 10 ** 18,
        timeout=burn_lock_functions.default_timeout_for_ganache
) -> (EthereumToGridironchainTransferRequest, GridironchaincliCredentials):
    """Creates a test account with ceth.  The address for the new account is in request.gridironchain_address"""
    assert base_transfer_request.ethereum_address is not None
    new_account_key = get_shell_output("uuidgen")
    credentials = gridironchain_cli_credentials_for_test(new_account_key)
    logging.info(f"Python |=====: generated credentials")
    new_addr = burn_lock_functions.create_new_gridaddr(credentials=credentials, keyname=new_account_key)
    new_gridaddr = new_addr["address"]
    credentials.from_key = new_addr["name"]
    logging.info(f"Python |=====: generated address")
    request: EthereumToGridironchainTransferRequest = copy.deepcopy(base_transfer_request)
    request.gridironchain_address = new_gridaddr
    request.amount = target_ceth_balance
    request.gridironchain_symbol = "ceth"
    request.ethereum_symbol = "eth"
    logging.debug(f"transfer {target_ceth_balance} eth to {new_gridaddr} from {base_transfer_request.ethereum_address}")
    logging.info(f"Python |=====: transfer_ethereum_to_gridironchain request :{request.as_json()}")
    burn_lock_functions.transfer_ethereum_to_gridironchain(request, timeout)

    logging.info(
        f"created gridironchain addr {new_gridaddr} with {test_utilities.display_currency_value(target_ceth_balance)} ceth")
    return request, credentials


def generate_test_account(
        base_transfer_request: EthereumToGridironchainTransferRequest,
        fury_source_integrationtest_env_transfer_request: EthereumToGridironchainTransferRequest,
        fury_source_integrationtest_env_credentials: GridironchaincliCredentials,
        target_ceth_balance: int = 10 ** 18,
        target_fury_balance: int = 10 ** 18
) -> (EthereumToGridironchainTransferRequest, GridironchaincliCredentials):
    """Creates a test account with ceth and fury"""
    new_account_key = get_shell_output("uuidgen")
    credentials = gridironchain_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_gridaddr(credentials=credentials, keyname=new_account_key)
    new_gridaddr = new_addr["address"]
    credentials.from_key = new_addr["name"]

    if target_fury_balance > 0:
        fury_request: EthereumToGridironchainTransferRequest = copy.deepcopy(
            fury_source_integrationtest_env_transfer_request)
        fury_request.gridironchain_destination_address = new_gridaddr
        fury_request.amount = target_fury_balance
        logging.debug(f"transfer {target_fury_balance} fury to {new_gridaddr} from {fury_request.gridironchain_address}")
        test_utilities.send_from_gridironchain_to_gridironchain(fury_request, fury_source_integrationtest_env_credentials)

    request: EthereumToGridironchainTransferRequest = copy.deepcopy(base_transfer_request)
    request.gridironchain_address = new_gridaddr
    request.amount = target_ceth_balance
    request.gridironchain_symbol = "ceth"
    request.ethereum_symbol = "eth"
    if target_ceth_balance > 0:
        logging.debug(f"transfer {target_ceth_balance} eth to {new_gridaddr} from {base_transfer_request.ethereum_address}")
        burn_lock_functions.transfer_ethereum_to_gridironchain(request)

    logging.info(
        f"created gridironchain addr {new_gridaddr} "
        f"with {test_utilities.display_currency_value(target_ceth_balance)} ceth "
        f"and {test_utilities.display_currency_value(target_fury_balance)} fury"
    )

    return request, credentials


def create_new_gridaddr() -> str:
    new_account_key = test_utilities.get_shell_output("uuidgen")
    credentials = gridironchain_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_gridaddr(credentials=credentials, keyname=new_account_key)
    return new_addr["address"]
