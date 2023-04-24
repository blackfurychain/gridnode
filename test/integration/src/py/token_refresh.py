import copy
import json
import logging
import os

import burn_lock_functions
import test_utilities
from burn_lock_functions import EthereumToGridironchainTransferRequest
from integration_env_credentials import create_new_gridaddr_and_credentials
from test_utilities import create_new_currency


def test_can_mint_token_and_peg_it_for_everything_in_whitelist(
        basic_transfer_request: EthereumToGridironchainTransferRequest,
        smart_contracts_dir,
        bridgebank_address,
        solidity_json_path,
        operator_address,
        ethereum_network,
        source_ethereum_address,
        fury_source
):
    logging.info("token_refresh needs to use the operator private key, setting that to ETHEREUM_PRIVATE_KEY")
    os.environ["ETHEREUM_PRIVATE_KEY"] = test_utilities.get_required_env_var("OPERATOR_PRIVATE_KEY")
    request = copy.deepcopy(basic_transfer_request)
    request.gridchain_address = fury_source
    request.ethereum_address = source_ethereum_address
    amount_in_tokens = int(test_utilities.get_required_env_var("TOKEN_AMOUNT"))

    tokens = test_utilities.get_whitelisted_tokens(request)
    logging.info(f"whitelisted tokens: {tokens}")

    for t in tokens:
        destination_symbol = "c" + t["symbol"]
        if t["symbol"] == "efury":
            destination_symbol = "fury"
        try:
            logging.info(f"sending {t}")
            request.amount = amount_in_tokens * (10 ** int(t["decimals"]))
            request.ethereum_symbol = t["token"]
            request.gridchain_symbol = destination_symbol
            request.ethereum_address = operator_address
            test_utilities.mint_tokens(request, operator_address)
            test_utilities.send_from_ethereum_to_gridchain(request)
        except Exception as e:
            # try to get as many tokens across the bridge as you can,
            # don't stop if one of them fails
            logging.info(f"failed to mint and send for {t}, error was {e}")
    logging.info(f"sent new batch of tokens to {fury_source}")
    test_utilities.get_gridchain_addr_balance(fury_source, request.gridnoded_node, "fury")