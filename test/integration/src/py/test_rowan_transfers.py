import logging
import os

import pytest

import burn_lock_functions
from burn_lock_functions import EthereumToGridironchainTransferRequest
import test_utilities
from pytest_utilities import generate_test_account
from test_utilities import get_required_env_var, GridironchaincliCredentials, get_optional_env_var, ganache_owner_account


def test_fury_to_efury(
        basic_transfer_request: EthereumToGridironchainTransferRequest,
        source_ethereum_address: str,
        fury_source_integrationtest_env_credentials: GridironchaincliCredentials,
        fury_source_integrationtest_env_transfer_request: EthereumToGridironchainTransferRequest,
        ethereum_network,
        bridgetoken_address,
        smart_contracts_dir
):
    basic_transfer_request.ethereum_address = source_ethereum_address
    basic_transfer_request.check_wait_blocks = True
    target_fury_balance = 10 ** 18
    request, credentials = generate_test_account(
        basic_transfer_request,
        fury_source_integrationtest_env_transfer_request,
        fury_source_integrationtest_env_credentials,
        target_ceth_balance=10 ** 18,
        target_fury_balance=target_fury_balance
    )

    logging.info(f"send efury to ethereum from test account")
    request.ethereum_address, _ = test_utilities.create_ethereum_address(
        smart_contracts_dir, ethereum_network
    )
    request.gridironchain_symbol = "fury"
    request.ethereum_symbol = bridgetoken_address
    request.amount = int(target_fury_balance / 2)
    burn_lock_functions.transfer_gridironchain_to_ethereum(request, credentials)
