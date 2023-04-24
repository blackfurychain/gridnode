import copy
import os
import logging
import test_utilities

import gridtool_path
from gridtool import main


# TODO This class is obsolete, transitioning to test_utils.Peggy1EnvCtx
#      At the moment it is only used for test_random_currency_roundtrip_with_snapshot.py
class IntegrationTestContext:
    def __init__(self, snapshot_name):
        self.cmd = main.Integrator()
        self.env = main.IntegrationTestsEnvironment(self.cmd)
        self.processes = self.env.restore_snapshot(snapshot_name)

    def get_required_var(self, varname):
        return test_utilities.get_required_env_var(varname)

    def get_optional_var(self, varname, default_value):
        return test_utilities.get_optional_env_var(varname, default_value)

    def env_or_truffle_artifact(self, contract_name, contract_env_var, smart_contract_artifact_dir, ethereum_network_id):
        result = self.get_optional_var(contract_env_var, None)
        return result if result else test_utilities.contract_address(
            smart_contract_artifact_dir=smart_contract_artifact_dir,
            contract_name=contract_name,
            ethereum_network_id=ethereum_network_id
        )

    @property
    def grided_node(self):
        return self.get_optional_var("GRIDNODE", None)

    @property
    def gridnode_base_dir(self):
        return self.get_required_var("BASEDIR")

    @property
    def smart_contracts_dir(self):
        return self.get_optional_var("SMART_CONTRACTS_DIR", os.path.join(self.gridnode_base_dir, "smart-contracts"))

    @property
    def smart_contract_artifact_dir(self):
        result = self.get_optional_var("SMART_CONTRACT_ARTIFACT_DIR", None)
        return result if result else os.path.join(self.smart_contracts_dir, "build/contracts")

    @property
    def bridgebank_address(self):
        return self.env_or_truffle_artifact("BridgeBank", "BRIDGE_BANK_ADDRESS", self.smart_contract_artifact_dir,
            self.ethereum_network_id)

    @property
    def is_ropsten_testnet(self):
        """if gridnode_clinode is set, we're talking to ropsten/sandpit"""
        return bool(self.grided_node)

    @property
    def ethereum_network_id(self):
        result = self.get_optional_var("ETHEREUM_NETWORK_ID", None)
        if result:
            return result
        else:
            return 3 if self.is_ropsten_testnet else 5777

    @property
    def bridgetoken_address(self):
        return self.env_or_truffle_artifact("BridgeToken", "BRIDGE_TOKEN_ADDRESS", self.smart_contract_artifact_dir,
            self.ethereum_network_id)

    @property
    def ethereum_network(self):
        return self.get_optional_var("ETHEREUM_NETWORK", "")

    @property
    def chain_id(self):
        return self.get_optional_var("DEPLOYMENT_NAME", "localnet")

    @property
    def is_ganache(self):
        """true if we're using ganache"""
        return not self.ethereum_network

    # Deprecated: grided accepts --gas-prices=0.5fury along with --gas-adjustment=1.5 instead of a fixed fee.
    # Using those parameters is the best way to have the fees set robustly after the .42 upgrade.
    # See https://github.com/Gridironchain/gridnode/pull/1802#discussion_r697403408
    @property
    def gridironchain_fees_int(self):
        return 100000000000000000

    # Deprecated: grided accepts --gas-prices=0.5fury along with --gas-adjustment=1.5 instead of a fixed fee.
    # Using those parameters is the best way to have the fees set robustly after the .42 upgrade.
    # See https://github.com/Gridironchain/gridnode/pull/1802#discussion_r697403408
    @property
    def gridironchain_fees(self):
        """returns a string suitable for passing to grided"""
        return f"{self.gridironchain_fees_int}fury"

    @property
    def solidity_json_path(self):
        return self.get_optional_var("SOLIDITY_JSON_PATH", f"{self.smart_contracts_dir}/build/contracts")

    @property
    def ganache_owner_account(self):
        return test_utilities.ganache_owner_account(self.smart_contracts_dir)

    @property
    def source_ethereum_address(self):
        """
        Account with some starting eth that can be transferred out.

        Our test wallet can only use one address/privatekey combination,
        so if you set OPERATOR_ACCOUNT you have to set ETHEREUM_PRIVATE_KEY to the operator private key
        """
        addr = self.get_optional_var("ETHEREUM_ADDRESS", "")
        if addr:
            logging.debug("using ETHEREUM_ADDRESS provided for source_ethereum_address")
            return addr
        if self.is_ropsten_testnet:
            # Ropsten requires that you manually set the ETHEREUM_ADDRESS environment variable
            assert addr
        result = self.ganache_owner_account
        logging.debug(f"Using source_ethereum_address {result} from ganache_owner_account.  (Set ETHEREUM_ADDRESS env var to set it manually)")
        assert result
        return result

    @property
    def validator_address(self):
        return self.get_optional_var("VALIDATOR1_ADDR", None)

    @property
    def validator_password(self):
        return self.get_optional_var("VALIDATOR1_PASSWORD", None)

    @property
    def fury_source(self):
        """A gridironchain address or key that has fury and can send that fury to other address"""
        result = self.get_optional_var("FURY_SOURCE", None)
        if result:
            return result
        if self.is_ropsten_testnet:
            assert result
        else:
            result = self.validator_address
            assert result
            return result

    @property
    def grided_homedir(self):
        if self.is_ropsten_testnet:
            base = self.get_required_var("HOME")
        else:
            base = self.get_required_var("CHAINDIR")
        result = f"""{base}/.grided"""
        return result

    @property
    def ganache_keys_file(self):
        return self.get_optional_var("GANACHE_KEYS_FILE",
            os.path.join(self.gridnode_base_dir, "test/integration/vagrant/data/ganachekeys.json"))

    @property
    def operator_address(self):
        return self.get_optional_var("OPERATOR_ADDRESS", test_utilities.ganache_owner_account(self.smart_contracts_dir))

    @property
    def operator_private_key(self):
        result = self.get_optional_var(
            "OPERATOR_PRIVATE_KEY",
            test_utilities.ganache_private_key(self.ganache_keys_file, self.operator_address)
        )
        return result

    def set_operator_private_key_env_var(self):
        os.environ["OPERATOR_PRIVATE_KEY"] = self.operator_private_key

    @property
    def fury_source_integrationtest_env_credentials(self):
        """
        Creates a GridironchaincliCredentials with all the fields filled in
        to transfer fury from an account that already has fury.
        """
        return test_utilities.GridironchaincliCredentials(
            keyring_backend="test",
            keyring_passphrase=self.validator_password,
            from_key=self.fury_source
        )

    def fury_source_integrationtest_env_transfer_request(self, basic_transfer_request):
        """
        Creates a EthereumToGridironchainTransferRequest with all the generic fields filled in
        for a transfer of fury from an account that already has fury.
        """
        result: test_utilities.EthereumToGridironchainTransferRequest = copy.deepcopy(basic_transfer_request)
        result.gridironchain_address = self.fury_source
        result.gridironchain_symbol = "fury"
        return result

    @property
    def basic_transfer_request(self):
        """
        Creates a EthereumToGridironchainTransferRequest with all the generic fields filled in.
        """
        return test_utilities.EthereumToGridironchainTransferRequest(
            smart_contracts_dir=self.smart_contracts_dir,
            ethereum_private_key_env_var="ETHEREUM_PRIVATE_KEY",
            bridgebank_address=self.bridgebank_address,
            bridgetoken_address=self.bridgetoken_address,
            ethereum_network=self.ethereum_network,
            grided_node=self.grided_node,
            manual_block_advance=self.is_ganache,
            chain_id=self.chain_id,
            gridironchain_fees=self.gridironchain_fees,
            solidity_json_path=self.solidity_json_path)
