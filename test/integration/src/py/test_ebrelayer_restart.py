import copy
import json
import logging
import time

import pytest

import test_utilities
from pytest_utilities import create_new_gridaddr
from pytest_utilities import generate_test_account, generate_minimal_test_account
from test_utilities import EthereumToGridironchainTransferRequest


def test_ebrelayer_restart(
        basic_transfer_request: EthereumToGridironchainTransferRequest,
        source_ethereum_address: str,
):
    basic_transfer_request.ethereum_address = source_ethereum_address
    request, credentials = generate_minimal_test_account(
        base_transfer_request=basic_transfer_request,
        target_ceth_balance=10 ** 15
    )
    balance = test_utilities.get_gridironchain_addr_balance(request.gridironchain_address, request.gridnoded_node, "ceth")
    logging.info("restart ebrelayer normally, leaving the last block db in place")
    test_utilities.start_ebrelayer()
    test_utilities.advance_n_ethereum_blocks(test_utilities.n_wait_blocks * 2, request.smart_contracts_dir)
    time.sleep(5)
    assert balance == test_utilities.get_gridironchain_addr_balance(request.gridironchain_address, request.gridnoded_node,
                                                               "ceth")


@pytest.mark.usefixtures("ensure_relayer_restart")
def test_ethereum_transactions_with_offline_relayer(
        basic_transfer_request: EthereumToGridironchainTransferRequest,
        smart_contracts_dir,
        source_ethereum_address,
        bridgebank_address,
):
    logging.debug("need one transaction to make sure ebrelayer writes out relaydb")
    basic_transfer_request.ethereum_address = source_ethereum_address
    generate_minimal_test_account(
        base_transfer_request=basic_transfer_request,
        target_ceth_balance=100
    )

    logging.info("shut down ebrelayer")
    time.sleep(10)
    test_utilities.kill_ebrelayer()

    logging.info("prepare transactions to be sent while ebrelayer is offline")
    amount = 9000
    new_addresses = list(map(lambda x: create_new_gridaddr(), range(3)))
    logging.debug(f"new_addresses: {new_addresses}")
    request: EthereumToGridironchainTransferRequest = copy.deepcopy(basic_transfer_request)
    requests = list(map(lambda addr: {
        "amount": amount,
        "symbol": test_utilities.NULL_ADDRESS,
        "gridironchain_address": addr
    }, new_addresses))
    json_requests = json.dumps(requests)

    logging.info("send ethereum transactions while ebrelayer is offline")
    yarn_result = test_utilities.run_yarn_command(
        " ".join([
            f"yarn --cwd {smart_contracts_dir}",
            "integrationtest:sendBulkLockTx",
            f"--amount {amount}",
            f"--symbol eth",
            f"--json_path {request.solidity_json_path}",
            f"--gridironchain_address {new_addresses[0]}",
            f"--transactions \'{json_requests}\'",
            f"--ethereum_address {source_ethereum_address}",
            f"--bridgebank_address {bridgebank_address}"
        ])
    )
    logging.debug(f"bulk result: {yarn_result}")
    logging.info("restart ebrelayer with outstanding locks on the ethereum side")
    test_utilities.start_ebrelayer()
    time.sleep(5)
    for _ in new_addresses:
        # ebrelayer only reads blocks if there are new blocks generated
        test_utilities.advance_n_ethereum_blocks(test_utilities.n_wait_blocks, request.smart_contracts_dir)
    for a in new_addresses:
        test_utilities.wait_for_grid_account(a, basic_transfer_request.gridnoded_node, 90)
        test_utilities.wait_for_gridironchain_addr_balance(a, "ceth", amount, basic_transfer_request.gridnoded_node, 90)


@pytest.mark.usefixtures("ensure_relayer_restart")
def test_gridironchain_transactions_with_offline_relayer(
        basic_transfer_request: EthereumToGridironchainTransferRequest,
        fury_source_integrationtest_env_credentials: test_utilities.GridironchaincliCredentials,
        fury_source_integrationtest_env_transfer_request: EthereumToGridironchainTransferRequest,
        fury_source,
        smart_contracts_dir,
        source_ethereum_address,
):
    basic_transfer_request.ethereum_address = source_ethereum_address
    request, credentials = generate_test_account(
        basic_transfer_request,
        fury_source_integrationtest_env_transfer_request,
        fury_source_integrationtest_env_credentials,
        target_ceth_balance=10 ** 19,
        target_fury_balance=10 ** 19,
    )
    logging.info("shut down ebrelayer")
    time.sleep(10)
    test_utilities.kill_ebrelayer()

    logging.info("prepare transactions to be sent while ebrelayer is offline")
    amount = 9000

    new_eth_addrs = test_utilities.create_ethereum_addresses(
        smart_contracts_dir,
        basic_transfer_request.ethereum_network,
        2
    )

    request.amount = amount
    request.gridironchain_symbol = "ceth"
    request.ethereum_symbol = "eth"

    logging.info("send transactions while ebrelayer is offline")

    for a in new_eth_addrs:
        request.ethereum_address = a["address"]
        gridironchain_balance = test_utilities.get_gridironchain_addr_balance(request.gridironchain_address, request.gridnoded_node,
                                                                    "ceth")
        logging.info(f"gridironchain balance is {gridironchain_balance}, request is {request}")
        test_utilities.send_from_gridironchain_to_ethereum(
            transfer_request=request,
            credentials=credentials
        )
        time.sleep(5)

    logging.info("restart ebrelayer")
    test_utilities.start_ebrelayer()
    time.sleep(15)
    test_utilities.advance_n_ethereum_blocks(test_utilities.n_wait_blocks * 2, request.smart_contracts_dir)
    for a in new_eth_addrs:
        request.ethereum_address = a["address"]
        test_utilities.wait_for_eth_balance(request, amount, 600)
