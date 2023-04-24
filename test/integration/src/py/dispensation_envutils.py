import logging
import os
import time
import json
import pytest
import string
import random

import burn_lock_functions
from burn_lock_functions import EthereumToGridironchainTransferRequest
import test_utilities
from pytest_utilities import generate_test_account
from integration_env_credentials import gridchain_cli_credentials_for_test
from test_utilities import get_required_env_var, GridironchaincliCredentials, get_optional_env_var, ganache_owner_account, \
    get_shell_output_json, get_shell_output, detect_errors_in_gridnoded_output, get_transaction_result, amount_in_wei


# CODE TO GENERATE RANDOM STRING FOR DISPENSATION NAME AS DISPENSATION NAME IS A UNIQUE KEY
def id_generator(size=6, chars=string.ascii_uppercase + string.digits):
    return ''.join(random.choice(chars) for _ in range(size))


# CODE TO GENERATE NEW ADDRESS AND ADD IT IN THE KEYRING
def create_new_gridaddr_and_key():
    new_account_key = test_utilities.get_shell_output("uuidgen").replace("-", "")
    credentials = gridchain_cli_credentials_for_test(new_account_key)
    new_addr = burn_lock_functions.create_new_gridaddr(credentials=credentials, keyname=new_account_key)
    return new_addr["address"], new_addr["name"]


# CODE TO SEND SOME SAMPLE TOKEN TO NEW ADDRESS
def send_sample_fury(from_address, to_address, amount, keyring_backend, chain_id, offline):
    logging.debug(f"transfer_fury")
    gridchain_fees_entry = f"--fees 100000000000000000fury"
    keyring_backend_entry = f"--keyring-backend {keyring_backend}"
    cmd = " ".join([
        "gridnoded tx bank send",
        f"{from_address}",
        f"{to_address}",
        f"{amount}",
        keyring_backend_entry,
        gridchain_fees_entry,
        f"--chain-id {chain_id}",
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    assert (json_str.get("code", 0) == 0)
    return json_str


# CODE TO QUERY BLOCK FOR NEW DISPENSATION TXN
def query_block_claim(txn_hash):
    cmd = " ".join([
        "gridnoded query tx",
        f"{txn_hash}",
        "--chain-id localnet",
        "-o json"
    ])
    json_str = get_shell_output_json(cmd)
    return json_str


# CODE TO CHECK ACCOUNT BALANCE
def balance_check(address, currency):
    logging.debug(f"check_balance")
    cmd = " ".join([
        "gridnoded query bank balances",
        f"{address}",
        f"--denom {currency}",
        f"-o json"
    ])
    json_str = get_shell_output_json(cmd)
    balance = json_str['amount']
    return balance


#CODE TO CREATE A CLI TO CREATE A SINGLE_KEY ONLINE DISPENSATION TXN
def create_online_singlekey_txn(
        claimType,
        signing_address,
        chain_id
):
    logging.debug(f"create_online_dispensation")
    gridchain_gas_entry = f"--gas auto"
    gridchain_fees_entry = f"--fees 50000fury"
    keyring_backend_entry = f"--keyring-backend test"
    output = 'output.json'
    cmd = " ".join([
        "gridnoded tx dispensation create",
        f"{claimType}",
        output,
        gridchain_gas_entry,
        gridchain_fees_entry,
        f"--from {signing_address}",
        f"--chain-id {chain_id}",
        keyring_backend_entry,
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    assert (json_str.get("code", 0) == 0)
    txn = json_str["txhash"]
    return txn


#CODE TO CREATE A CLI TO CREATE A SINGLE_KEY ONLINE DISPENSATION TXN WITH AN ASYNC FLAG
def create_online_singlekey_async_txn(
        claimType,
        signing_address,
        chain_id
):
    logging.debug(f"create_online_dispensation")
    gridchain_fees_entry = f"--fees 150000fury"
    gridchain_gas_entry = f"--gas auto --gas-adjustment=1.5"
    keyring_backend_entry = f"--keyring-backend test"
    output = 'output.json'
    cmd = " ".join([
        "gridnoded tx dispensation create",
        f"{claimType}",
        output,
        gridchain_fees_entry,
        gridchain_gas_entry,
        f"--from {signing_address}",
        f"--chain-id {chain_id}",
        keyring_backend_entry,
        f"--broadcast-mode async",
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    assert (json_str.get("code", 0) == 0)
    txn = json_str["txhash"]
    return txn


#CODE TO CREATE A CLI TO CREATE A SINGLE_KEY OFFLINE DISPENSATION TXN
def create_offline_singlekey_txn(
        claimType,
        signing_address,
        chain_id,
    ):
    logging.debug(f"create_unsigned_offline_dispensation_txn")
    gridchain_fees_entry = f"--fees 150000fury"
    gridchain_gas_entry = f"--gas auto --gas-adjustment=1.5"
    output = 'output.json'
    cmd = " ".join([
        "gridnoded tx dispensation create",
        f"{claimType}",
        output,
        f"--from {signing_address}",
        f"--chain-id {chain_id}",
        gridchain_fees_entry,
        gridchain_gas_entry,
        f"--generate-only", 
        f"--yes -o json"  
    ])
    json_str = get_shell_output_json(cmd)
    assert(json_str.get("code", 0) == 0)
    return json_str


#CODE TO SIGN DISPENSATION TXN BY A USER
def sign_txn(signingaddress, file):
    keyring_backend_entry = f"--keyring-backend test"
    cmd = " ".join([
        "gridnoded tx sign",
        f"--from {signingaddress}",
        f"{file}",
        keyring_backend_entry,
        "--chain-id localnet",
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    return json_str


#CODE TO BROADCAST SINGLE SIGNED TXN ON BLOCK
def broadcast_txn(file_path):
    cmd = " ".join([
        "gridnoded tx broadcast",
        f"{file_path}",
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    txn = json_str["txhash"]
    return txn


#CODE TO BROADCAST SINGLE SIGNED TXN ON BLOCK WITH AN ASYNC FLAG
def broadcast_async_txn(file_path):
    cmd = " ".join([
        "gridnoded tx broadcast",
        f"{file_path}",
        f"--broadcast-mode async",
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    txn = json_str["txhash"]
    return txn


def create_online_singlekey_txn_with_runner(
        claimType,
        runner_address,
        distributor_address,
        chain_id
):
    logging.debug(f"create_online_dispensation")
    gridchain_fees_entry = f"--fees 150000fury"
    gridchain_gas_entry = f"--gas auto --gas-adjustment=1.5"
    keyring_backend_entry = f"--keyring-backend test"
    output = 'output.json'
    cmd = " ".join([
        "gridnoded tx dispensation create",
        f"{claimType}",
        output,
        runner_address,
        gridchain_fees_entry,
        gridchain_gas_entry,
        f"--from {distributor_address}",
        f"--chain-id {chain_id}",
        keyring_backend_entry,
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    assert (json_str.get("code", 0) == 0)
    txn = json_str["txhash"]
    return txn


#CODE TO CREATE A CLI TO CREATE A SINGLE_KEY OFFLINE DISPENSATION TXN
def create_offline_singlekey_txn_with_runner(
        claimType,
        runner_address,
        distributor_address,
        chain_id
    ):
    logging.debug(f"create_unsigned_offline_dispensation_txn")
    gridchain_fees_entry = f"--fees 150000fury"
    gridchain_gas_entry = f"--gas auto --gas-adjustment=1.5"
    keyring_backend_entry = f"--keyring-backend test"
    output = 'output.json'
    cmd = " ".join([
        "gridnoded tx dispensation create",
        f"{claimType}",
        output,
        runner_address,
        f"--from {distributor_address}",
        f"--chain-id {chain_id}",
        gridchain_fees_entry,
        gridchain_gas_entry,
        keyring_backend_entry,
        f"--generate-only", 
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    assert(json_str.get("code", 0) == 0)
    return json_str


#CODE TO EXECUTE RUN DISPENSATION CLI
def run_dispensation(
        distribution_name,
        claimType,
        distribution_count,
        runner_address,
        chain_id
    ):
    logging.debug(f"RUN DISPENSATION CLI LOGGING")
    gridchain_gas_entry = f"--gas auto --gas-adjustment=1.5"
    gridchain_fees_entry = f"--fees 200000fury"
    keyring_backend_entry = f"--keyring-backend test"
    cmd = " ".join([
        "gridnoded tx dispensation run",
        distribution_name,
        f"{claimType}",
        distribution_count,
        f"--from {runner_address}",
        f"--chain-id {chain_id}",
        gridchain_gas_entry,
        gridchain_fees_entry,
        keyring_backend_entry,
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    assert(json_str.get("code", 0) == 0)
    txn = json_str["txhash"]
    return txn


#CODE TO QUERY A NEW CLAIM 
def query_created_claim(claimType):
    cmd = " ".join([
        "gridnoded q dispensation claims-by-type",
        f"{claimType}",
        "--chain-id localnet",
        f"-o json"
    ])
    json_str = get_shell_output_json(cmd)
    return json_str


#CODE TO CREATE A NEW CLAIM
def create_claim(
        gridchain_address,
        claimType,
        keyring_backend,
        chain_id
    ):
    logging.debug(f"create_claim")
    keyring_backend_entry = f"--keyring-backend {keyring_backend}"
    gridchain_gas_entry = f"--gas auto --gas-adjustment=1.5"
    gridchain_fees_entry = f"--fees 100000000000000000fury"
    cmd = " ".join([
        "gridnoded tx dispensation claim",
        f"{claimType}",
        f"--from {gridchain_address}",
        gridchain_fees_entry,
        gridchain_gas_entry,
        f"--chain-id {chain_id}",
        keyring_backend_entry,
        f"--yes -o json"
    ])
    json_str = get_shell_output_json(cmd)
    assert(json_str.get("code", 0) == 0)
    txn = json_str["txhash"]
    return txn
