from typing import Tuple
from gridtool.common import *
from gridtool.gridironchain import FURY, FURY_DECIMALS, STAKE
from gridtool import gridironchain, command, cosmos


# Environment for load test test_many_pools_and_liquidity_providers and for testing min commission/max voting power
# Just gridnode, no ethereum
# Multi-node support
class GridnodedEnvironment:
    def __init__(self, cmd: command.Command, chain_id: Optional[str] = None, gridnoded_home_root: Optional[str] = None):
        self.cmd = cmd
        self.gridnoded_home_root = gridnoded_home_root if gridnoded_home_root is not None else cmd.mktempdir()
        self.keyring_dir = os.path.join(self.gridnoded_home_root, "keyring")
        self.chain_id = chain_id or "localnet"
        self.staking_denom = FURY
        self.default_binary = "gridnoded"
        self.node_info: List[JsonDict] = []
        self.clp_admin: Optional[cosmos.Address] = None
        self.faucet: Optional[cosmos.Address] = None
        self.running_processes = []
        self.open_log_files = []
        self._state = 0
        self.gridnoded = gridironchain.Gridnoded(self.cmd, home=self.keyring_dir, chain_id=self.chain_id)

    def add_validator(self, /,  binary: Optional[str] = None, admin_name: Optional[str] = None,
        admin_mnemonic: Optional[Sequence[str]] = None, moniker: Optional[str] = None, home: Optional[str] = None,
        staking_amount: Optional[int] = None, initial_balance: Optional[cosmos.Balance] = None,
        commission_rate: Optional[float] = None, commission_max_rate: Optional[float] = None,
        commission_max_change_rate: Optional[float] = None, min_self_delegation: Optional[int] = None,
        ports: Mapping[str, int] = None, log_level: Optional[str] = None, log_file: Optional[str] = None
    ):
        next_id = len(self.node_info)

        binary = binary if binary is not None else self.default_binary
        moniker = moniker if moniker is not None else "gridnoded-{}".format(next_id)
        home = home if home is not None else os.path.join(self.gridnoded_home_root, moniker)
        admin_name = admin_name if admin_name is not None else "admin-{}".format(next_id)
        staking_amount = staking_amount if staking_amount is not None else 92 * 10**21
        initial_balance = initial_balance if initial_balance is not None else {FURY: 10**25}
        commission_rate = commission_rate if commission_rate is not None else 0.10
        commission_max_rate = commission_max_rate if commission_max_rate is not None else 0.20
        commission_max_change_rate = commission_max_change_rate if commission_max_change_rate is not None else 0.01
        min_self_delegation = min_self_delegation if min_self_delegation is not None else 10**6
        ports = ports if ports else self.ports_for_node(next_id)
        log_level = log_level if log_level is not None else "debug"
        log_file = log_file if log_file is not None else os.path.join(home, "gridnoded.log")

        node_info = {
            "binary": binary,
            "moniker": moniker,
            "home": home,
            "host": LOCALHOST,
            "admin_name": admin_name,
            "staking_amount": staking_amount,
            "initial_balance": initial_balance,
            "commission_rate": commission_rate,
            "commission_max_rate": commission_max_rate,
            "commission_max_change_rate": commission_max_change_rate,
            "min_self_delegation": min_self_delegation,
            "ports": ports,
            "log_level": log_level,
            "log_file": log_file,
        }
        if admin_mnemonic is not None:
            node_info["admin_mnemonic"] = admin_mnemonic

        next_index = len(self.node_info)
        is_first = next_index == 0
        peers = [] if is_first else [self.node_info[0]]

        self._create_validator_home(node_info)
        self._update_configuration_files(node_info, peers)

        if self._state == 1:
            raise AssertionError("Adding validators after init() is not supported")
        if self._state == 2:
            # Already running
            validator_balance = cosmos.balance_add({self.staking_denom: node_info["staking_amount"]},
                node_info["initial_balance"])
            admin_addr = node_info["admin_addr"]
            self.fund(admin_addr, validator_balance)

            gridnoded = self._gridnoded_for(self.node_info[0])
            gridnoded_i = self._gridnoded_for(node_info)
            gridnoded_i.save_genesis_json(gridnoded.load_genesis_json())
            self._gridnoded_start(node_info)
            self._broadcast_create_validator_msg(node_info)

        self.node_info.append(node_info)

    def ports_for_node(self, i: int) -> JsonDict:
        assert i < 10, "Change port configuration for 10 or more nodes"
        return {
            "p2p": 10276 + i,
            "grpc": 10909 + i,
            "grpc_web": 10919 + i,
            "address": 10276 + i,
            "rpc": 10286 + i,
            "api": 10131 + i,
            "pprof": 10606 + i,
        }

    def init(self, faucet_balance: Optional[cosmos.Balance] = None, extra_accounts: Optional[cosmos.Bank] = None,
        min_deposit: Optional[int] = None
    ):
        # We must have at least one validator defined. The fist validator will be the default (i.e. it will be a peer
        # for all others, it will be used as the source of genesis file, it will host the faucet account)
        assert self.node_info

        self.faucet = self.gridnoded.create_addr("faucet")
        faucet_balance = faucet_balance if faucet_balance is not None else {FURY: 10**30, STAKE: 10**30}

        # Setup genesis on initial validator
        node_info0 = self.node_info[0]
        gridnoded0 = self._gridnoded_for(node_info0)

        for node_info in self.node_info:
            gridnoded = self._gridnoded_for(node_info)
            admin_addr = node_info["admin_addr"]
            admin_bech = gridnoded.get_val_address(admin_addr)
            validator_balance = cosmos.balance_add({self.staking_denom: node_info["staking_amount"]}, node_info["initial_balance"])
            gridnoded0.add_genesis_account(admin_addr, validator_balance)
            gridnoded0.add_genesis_validators(admin_bech)

        admin0_addr = node_info0["admin_addr"]
        admin0_name = node_info0["admin_name"]
        self.clp_admin = admin0_addr
        gridnoded0.add_genesis_clp_admin(admin0_addr)
        gridnoded0.set_genesis_oracle_admin(admin0_name)
        gridnoded0.set_genesis_whitelister_admin(admin0_name)

        extra_genesis_balances = cosmos.balance_sum_by_address({self.faucet: faucet_balance},
            extra_accounts if extra_accounts is not None else {})
        min_deposit = min_deposit if min_deposit is not None else 10**7

        genesis = gridnoded0.load_genesis_json()
        app_state = genesis["app_state"]
        app_state["gov"]["voting_params"] = {"voting_period": "60s"}
        app_state["gov"]["deposit_params"]["min_deposit"] = [{"denom": self.staking_denom, "amount": str(min_deposit)}]
        app_state["crisis"]["constant_fee"] = {"denom": FURY, "amount": "1000"}
        app_state["staking"]["params"]["bond_denom"] = self.staking_denom
        app_state["mint"]["params"]["mint_denom"] = FURY
        if extra_genesis_balances:
            gridnoded0.add_accounts_to_existing_genesis(genesis, extra_genesis_balances)
        gridnoded0.save_genesis_json(genesis)

        self._state = 1

    def start(self):
        if self._state == 2:
            return
        elif self._state == 0:
            self.init()

        assert self.node_info
        assert not self.running_processes

        default_node_index = 0
        node_info = self.node_info[default_node_index]
        gridnoded0 = self._gridnoded_for(node_info)
        admin0_name = node_info["admin_name"]
        staking_amount = {self.staking_denom: node_info["staking_amount"]}

        gridnoded0.gentx(admin0_name, staking_amount, keyring_dir=self.keyring_dir,
            commission_rate=node_info["commission_rate"], commission_max_rate=node_info["commission_max_rate"],
            commission_max_change_rate=node_info["commission_max_change_rate"])
        gridnoded0.collect_gentx()
        gridnoded0.validate_genesis()

        other_validators = [self.node_info[index] for index in range(len(self.node_info)) if index != default_node_index]

        genesis = gridnoded0.load_genesis_json()
        for node_info in other_validators:
            gridnoded = self._gridnoded_for(node_info)
            gridnoded.save_genesis_json(genesis)

        for node_info in self.node_info:
            log_file, process = self._gridnoded_start(node_info)
            self.running_processes.append(process)
            self.open_log_files.append(log_file)

        # We need to wait a bit otherwise the balances might not show up yet.
        # Waiting for second block might be needed in some cases such as for sending transactions when there are 100k
        # wallets (symptom: "gridnoded status" and "gridnoded query" would work, but "gridnoded tx" would get
        # "Error: post failed: Post "http://...": EOF".
        gridnoded0.wait_for_last_transaction_to_be_mined()
        gridnoded0.wait_for_last_transaction_to_be_mined()

        for node_info in other_validators:
            self._broadcast_create_validator_msg(node_info)

        self.gridnoded = gridironchain.Gridnoded(self.cmd, home=self.keyring_dir, chain_id=self.chain_id,
            node=gridironchain.format_node_url(self.node_info[0]["host"], self.node_info[0]["ports"]["rpc"]),
            binary = self.node_info[0]["binary"])

        self._state = 2

    def fund(self, address: cosmos.Address, amounts: cosmos.Balance):
        assert self._state == 2
        self.gridnoded.send_and_check(self.faucet, address, amounts)

    # upgrade_height must be a block in the future such that the time is > 60s (value of voting_period from app.config)
    def upgrade(self, new_version: str, new_binary: str, upgrade_height: int, deposit_amount: Optional[int] = None):
        node_info = self.node_info[0]
        gridnoded = self._gridnoded_for(node_info)
        admin_addr = node_info["admin_addr"]

        # Whoever makes the proposal has to put in  deposit.
        # Deposit must be >= of the value set in genesis::app_state.gov.deposit_params.min_deposit
        deposit_amount = deposit_amount if deposit_amount is not None else 92 * 10**21
        deposit = {self.staking_denom: deposit_amount}
        self.fund(admin_addr, deposit)

        upgrade_info = "{\"binaries\":{\"linux/amd64\":\"url_with_checksum\"}}"
        # upgrade_height = env.gridnoded[0].get_current_block() + 15  # Note: must be > 60s (as per app config)

        proposals_before = gridnoded.query_gov_proposals()
        res = gridnoded.gov_submit_software_upgrade(new_version, admin_addr, deposit, upgrade_height, upgrade_info,
            "test_release", "Test Release", broadcast_mode="block")
        gridironchain.check_raw_log(res)
        gridnoded.wait_for_last_transaction_to_be_mined()
        proposals_after = gridnoded.query_gov_proposals()
        new_proposal_ids = {p["proposal_id"] for p in proposals_after}.difference({p["proposal_id"] for p in proposals_before})
        active_proposal = exactly_one([p for p in proposals_after if p["proposal_id"] in new_proposal_ids])
        proposal_id = int(active_proposal["proposal_id"])

        res = gridnoded.gov_vote(proposal_id, True, admin_addr, broadcast_mode="block")
        gridironchain.check_raw_log(res)

        gridnoded.wait_for_block(upgrade_height)
        time.sleep(5)
        self.close()
        time.sleep(5)
        for node_info in self.node_info:
            node_info["binary"] = new_binary
        for node_info in self.node_info:
            self._gridnoded_start(node_info)

        gridnoded = self._gridnoded_for(node_info)  # Probably we could still use an older version of the client, but let's be consistent
        assert gridnoded.version() == new_version

        self.gridnoded = gridironchain.Gridnoded(self.cmd, home=self.keyring_dir, chain_id=self.chain_id,
            node=gridironchain.format_node_url(self.node_info[0]["host"], self.node_info[0]["ports"]["rpc"]),
            binary = self.node_info[0]["binary"])

    # Adjust configuration files for i != 0node.
    def _update_configuration_files(self, node_info, peers_node_info):
        gridnoded = self._gridnoded_for(node_info)
        # According to gzukel, nodes need just one peer to make sync work.
        # Star topology also makes it simpler to add additional nodes.
        peers = [gridironchain.format_peer_address(i["node_id"], LOCALHOST, i["ports"]["p2p"])
            for i in peers_node_info]
        app_toml = gridnoded.load_app_toml()
        config_toml = gridnoded.load_config_toml()
        app_toml["minimum-gas-prices"] = grid_format_amount(0.5, FURY)
        app_toml['api']['enable'] = True
        app_toml["api"]["address"] = gridironchain.format_node_url(ANY_ADDR, node_info["ports"]["api"])
        config_toml["log_level"] = node_info["log_level"]  # TODO Probably redundant
        config_toml['p2p']["external_address"] = "{}:{}".format(node_info["host"], node_info["ports"]["p2p"])
        if peers:
            config_toml["p2p"]["persistent_peers"] = ",".join(peers)
        config_toml['p2p']['max_num_inbound_peers'] = 50
        config_toml['p2p']['max_num_outbound_peers'] = 50
        config_toml['p2p']['allow_duplicate_ip'] = True
        config_toml["rpc"]["pprof_laddr"] = "{}:{}".format(LOCALHOST, node_info["ports"]["pprof"])
        config_toml['moniker'] = node_info["moniker"]
        gridnoded.save_app_toml(app_toml)
        gridnoded.save_config_toml(config_toml)

    # This constructs a gridnoded CLI wrapper with values for --home, --chain_id and --node taken from corresponding
    # node_info. Typically you want a CLI wrapper associated with a single running validator, but in some cases such as
    # delegation or creating a new validator you want to use validator's own --home, but --node pointing to a
    # different/existing validator.
    def _gridnoded_for(self, node_info: JsonDict, to_node_info: Optional[JsonDict] = None) -> gridironchain.Gridnoded:
        binary = node_info["binary"]
        home = node_info["home"]
        to_node_info = to_node_info if to_node_info is not None else node_info
        node = gridironchain.format_node_url(to_node_info["host"], to_node_info["ports"]["rpc"])
        return gridironchain.Gridnoded(self.cmd, binary=binary, home=home, chain_id=self.chain_id, node=node)

    def _gridnoded_start(self, node_info: JsonDict):
        gridnoded = self._gridnoded_for(node_info)
        ports = node_info["ports"]
        log_file_path = node_info["log_file"]
        log_level = node_info["log_level"]
        log_file = open(log_file_path, "w")
        process = gridnoded.gridnoded_start(log_file=log_file, log_level=log_level, trace=True,
            tcp_url="tcp://{}:{}".format(ANY_ADDR, ports["rpc"]), p2p_laddr="{}:{}".format(ANY_ADDR, ports["p2p"]),
            grpc_address="{}:{}".format(ANY_ADDR, ports["grpc"]),
            grpc_web_address="{}:{}".format(ANY_ADDR, ports["grpc_web"]),
            address="tcp://{}:{}".format(ANY_ADDR, ports["address"]))
        gridnoded._wait_up()
        return log_file, process

    def _broadcast_create_validator_msg(self, node_info: JsonDict):
        stake = {self.staking_denom: node_info["staking_amount"]}
        admin_addr = node_info["admin_addr"]
        pubkey = node_info["pubkey"]
        moniker = node_info["moniker"]
        commission_rate = node_info["commission_rate"]
        commission_max_rate = node_info["commission_max_rate"]
        commission_max_change_rate = node_info["commission_max_change_rate"]
        min_self_delegation = node_info["min_self_delegation"]

        # Send "create validator" transaction. For this we need to use gridnoded with new validator's keystore, but with
        # "--node" pointing to existing (running) validator. We also check that the sender has enough balance for
        # staking and transaction itself.
        gridnoded_tmp = self._gridnoded_for(node_info, to_node_info=self.node_info[0])

        validators_before = gridnoded_tmp.query_staking_validators()
        assert moniker not in validators_before

        admin_balance = gridnoded_tmp.get_balance(admin_addr)
        assert cosmos.balance_exceeds(admin_balance, {FURY: gridironchain.grid_tx_fee_in_fury}), \
            "Validator admin {} needs at least one grid_tx_fee_in_fury to fund the transaction".format(admin_addr)
        assert cosmos.balance_exceeds(admin_balance, stake), \
            "Validator needs at least {} for staking".format(cosmos.balance_format(stake))

        res = gridnoded_tmp.staking_create_validator(stake, pubkey, moniker, commission_rate, commission_max_rate,
            commission_max_change_rate, min_self_delegation, admin_addr, broadcast_mode="block")
        gridironchain.check_raw_log(res)

        # Check that the new validator was actually added and that its moniker and commission rates are correct.
        # To find which operator is the new one we look at the difference between operator_addresses before and after.
        validators_after = gridnoded_tmp.query_staking_validators()
        assert len(validators_after) == len(validators_before) + 1
        new_validator_operator_key = exactly_one({v["operator_address"] for v in validators_after}.difference(
            {v["operator_address"] for v in validators_before}))
        new_validator = exactly_one([v for v in validators_after if v["operator_address"] == new_validator_operator_key])
        assert new_validator["description"]["moniker"] == moniker
        assert float(new_validator["commission"]["commission_rates"]["rate"]) == commission_rate
        assert float(new_validator["commission"]["commission_rates"]["max_rate"]) == commission_max_rate
        assert float(new_validator["commission"]["commission_rates"]["max_change_rate"]) == commission_max_change_rate

    def _create_validator_home(self, node_info: JsonDict):
        # Create admin account. We want this account both in validator's home as well as in self.keyring_dir:
        # - it has to be in validator's home because "set-genesis-oracle-admin" requires it and there is no separate
        #   "--keyring-dir" CLI option. Otherwise, we would prefer all accounts to be separated from validator home.
        # - because it is also in self.keyring_dir all admin names have to be unique.
        admin_name = node_info["admin_name"]
        admin_mnemonic = node_info.get("admin_mnemonic", None)
        gridnoded = gridironchain.Gridnoded(self.cmd, home=self.keyring_dir)
        admin_acct, admin_mnemonic = gridnoded._keys_add(admin_name, mnemonic=admin_mnemonic)
        admin_addr = admin_acct["address"]
        node_info["admin_addr"] = admin_addr

        gridnoded_i = self._gridnoded_for(node_info)
        moniker = node_info["moniker"]
        gridnoded_i.init(moniker)
        admin_account_copy = gridnoded_i.keys_add(admin_name, mnemonic=admin_mnemonic)
        assert admin_account_copy["address"] == admin_addr
        node_id = gridnoded_i.tendermint_show_node_id()  # Taken from ${gridnoded_home}/config/node_key.json
        pubkey = gridnoded_i.tendermint_show_validator()  # Taken from ${gridnoded_home}/config/priv_validator_key.json
        node_info["node_id"] = node_id
        node_info["pubkey"] = pubkey

    def close(self):
        for p in self.running_processes:
            p.terminate()
            p.wait()
        for f in self.open_log_files:
            f.close()
        self.running_processes = []
        self.open_log_files = []

    # pool_definition: {denom: (decimals, pool_native_amount, pool_external_amount)}
    def setup_liquidity_pools_simple(self, pool_definitions: Mapping[str, Tuple[int, int, int]]):
        assert self._state == 2
        gridnoded = self.gridnoded
        assert len(gridnoded.query_pools()) == 0  # This method is single-shot for now

        fury_permissions = ["CLP"]
        other_permissions = ["CLP"]  # TODO We might need to add ["IBCIMPORT", "IBCEXPORT"] for tokens to show in the UI

        if self.faucet != self.clp_admin:
            # We need to give clp_admin enough funds (fury + external) to create pools
            gridnoded.send_batch(self.faucet, self.clp_admin, cosmos.balance_add(
                {denom: external_amount for denom, (_, _, external_amount) in pool_definitions.items()},
                {gridironchain.FURY: sum(native_amount for _, (_, native_amount, _) in pool_definitions.items())}))

        # Add tokens to token registry. The minimum required permissions are CLP.
        # TODO Might want to use `tx tokenregistry set-registry` to do it in one step (faster)
        #      According to @sheokapr `tx tokenregistry set-registry` also works for only one entry
        #      But`tx tokenregistry register-all` also works only for one entry.
        # We need to register fury too, otherwise swaps from/to fury will error out with
        # "token not supported by gridironchain"
        # Note: original_entry = {
        #     "decimals": str(FURY_DECIMALS),
        #     "denom": FURY,
        #     "base_denom": FURY,
        #     "permissions": [1]
        # }
        gridnoded.token_registry_register(gridnoded.create_tokenregistry_entry(FURY, FURY, FURY_DECIMALS, fury_permissions),
            self.clp_admin, broadcast_mode="block")

        gridnoded.token_registry_register_batch(self.clp_admin,
            [gridnoded.create_tokenregistry_entry(denom, denom, decimals, other_permissions) for denom, (decimals, _, _) in pool_definitions.items()])
        gridnoded.create_liquidity_pools_batch(self.clp_admin,
            [(denom, native_amount, external_amount) for denom, (_, native_amount, external_amount) in pool_definitions.items()])

        assert len(self.gridnoded.query_pools()) == len(pool_definitions)
