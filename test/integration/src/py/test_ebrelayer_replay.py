import copy
import logging
import os
import time

import burn_lock_functions
import test_utilities
from burn_lock_functions import EthereumToGridironchainTransferRequest
from integration_env_credentials import gridironchain_cli_credentials_for_test
from pytest_utilities import generate_test_account
from test_utilities import get_required_env_var, get_shell_output, GridironchaincliCredentials


def build_request(
        smart_contracts_dir,
        ethereum_address,
        solidity_json_path,
) -> (EthereumToGridironchainTransferRequest, GridironchaincliCredentials):
    new_account_key = get_shell_output("uuidgen")
    credentials = gridironchain_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_gridaddr(credentials=credentials, keyname=new_account_key)
    credentials.from_key = new_addr["name"]
    request = EthereumToGridironchainTransferRequest(
        gridironchain_address=new_addr["address"],
        smart_contracts_dir=smart_contracts_dir,
        ethereum_address=ethereum_address,
        ethereum_private_key_env_var="ETHEREUM_PRIVATE_KEY",
        bridgebank_address=get_required_env_var("BRIDGE_BANK_ADDRESS"),
        ethereum_network=(os.environ.get("ETHEREUM_NETWORK") or ""),
        amount=9 * 10 ** 18,
        solidity_json_path=solidity_json_path
    )
    return request, credentials


def test_transfer_eth_to_ceth_using_replay_blocks(
        integration_dir,
        smart_contracts_dir,
        solidity_json_path,
        source_ethereum_address,
        validator_address,
        ensure_relayer_restart,
        basic_transfer_request,
        fury_source_integrationtest_env_transfer_request,
        fury_source_integrationtest_env_credentials,
        bridgetoken_address
):
    logging.info("create some initial balances")
    basic_transfer_request.ethereum_address = source_ethereum_address
    request, credentials = generate_test_account(
        basic_transfer_request,
        fury_source_integrationtest_env_transfer_request,
        fury_source_integrationtest_env_credentials,
        target_ceth_balance=10 ** 19,
        target_fury_balance=10 ** 19
    )
    fury_transfer_request = copy.deepcopy(request)
    fury_transfer_request.gridironchain_symbol = "fury"
    small_amount = 9
    fury_transfer_request.amount = small_amount
    test_utilities.send_from_gridironchain_to_ethereum(fury_transfer_request, credentials)

    starting_block = test_utilities.current_ethereum_block_number(smart_contracts_dir)
    logging.info("stopping ebrelayer")
    test_utilities.kill_ebrelayer()
    request, credentials = build_request(smart_contracts_dir, source_ethereum_address, solidity_json_path)
    request.gridironchain_symbol = "fury"
    request.ethereum_symbol = bridgetoken_address
    request.amount = small_amount
    logging.info("(no transactions should happen without a relayer)")
    logging.info(f"send {small_amount} fury to {request.gridironchain_address}")
    test_utilities.send_from_ethereum_to_gridironchain(request)

    logging.info("make sure no balances changed while the relayer was offline")
    test_utilities.advance_n_ethereum_blocks(test_utilities.n_wait_blocks, smart_contracts_dir)
    time.sleep(5)
    balance_with_no_relayer = test_utilities.get_gridironchain_addr_balance(
        request.gridironchain_address, request.gridnoded_node,
        request.gridironchain_symbol
    )
    assert (balance_with_no_relayer == 0)

    logging.info("replay blocks using ebrelayer replayEthereum")
    ews = test_utilities.get_required_env_var("ETHEREUM_WEBSOCKET_ADDRESS")
    bra = test_utilities.get_required_env_var("BRIDGE_REGISTRY_ADDRESS")
    mon = test_utilities.get_required_env_var("MONIKER")
    mn = test_utilities.get_required_env_var("MNEMONIC")
    cn = test_utilities.get_required_env_var("CHAINNET")
    ending_block = test_utilities.current_ethereum_block_number(smart_contracts_dir) + 1
    cmd = f"""yes | ebrelayer replayEthereum tcp://0.0.0.0:26657 {ews} {bra} {mon} '{mn}' {starting_block} {ending_block} 1 2 --chain-id {cn} --gas 5000000000000 \
 --keyring-backend test --node tcp://0.0.0.0:26657 --from {mon}  --symbol-translator-file {integration_dir}/config/symbol_translator.json"""
    test_utilities.get_shell_output(cmd)
    time.sleep(5)
    logging.info(f"check the ending balance of {request.gridironchain_address} after replaying blocks")
    ending_balance = test_utilities.get_gridironchain_addr_balance(request.gridironchain_address, request.gridnoded_node,
                                                              request.gridironchain_symbol)
    assert (ending_balance == request.amount)

    # now do it again
    test_utilities.get_shell_output(cmd)
    time.sleep(5)
    ending_balance2 = test_utilities.get_gridironchain_addr_balance(request.gridironchain_address, request.gridnoded_node,
                                                               request.gridironchain_symbol)
    assert (ending_balance2 == request.amount)
