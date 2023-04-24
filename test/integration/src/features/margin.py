from contextlib import contextmanager
from gridtool.common import *
from gridtool.gridchain import FURY, FURY_DECIMALS
from gridtool import command, environments, project, gridchain, cosmos


def swap_pricing_formula(x, X, y, Y):
    # Pricing for swaps:
    # y = x*Y / (x + X)
    # x = y*X / (y + Y)
    # x = swap input or output (non-fury)
    # y = swap input or output (fury)
    # X = non-fury token pool balance
    # Y = fury pool balance
    # Note: a "swap" is subjected to price calculation of all swaps that happen in that block, so the final prices
    # are not known until after swaps have been processed.
    raise NotImplemented()


class TestMargin:
    def setup_method(self):
        self.cmd = command.Command()
        self.gridnoded_home_root = self.cmd.tmpdir("gridtool.tmp")
        self.cmd.rmdir(self.gridnoded_home_root)
        self.cmd.mkdir(self.gridnoded_home_root)
        self.default_pool_setup = [
            # denom,  decimals,  pool_native_amount, pool_external_amount, faucet_amount
            ["cusdc",        6,              10**25,               10**25,        10**30 ],
            ["cdash",       18,              10**25,               10**25,        10**30 ],
            ["ceth",        18,              10**25,               10**25,        10**30 ],
        ]
        project.Project(self.cmd, project_dir()).pkill()

    def teardown_method(self):
        project.Project(self.cmd, project_dir()).pkill()

    @contextmanager
    def with_test_env(self, pool_setup):
        faucet_balance = cosmos.balance_add({
            gridchain.FURY: 10**30,
            gridchain.STAKE: 10**30,
        }, {
            denom: external_amount + faucet_amount for denom, _, _, external_amount, faucet_amount in pool_setup
        })
        env = environments.GridnodedEnvironment(self.cmd, gridnoded_home_root=self.gridnoded_home_root)
        env.add_validator()
        env.init(faucet_balance=faucet_balance)
        env.start()
        pool_definitions = {denom: (decimals, native_amount, external_amount)
            for denom, decimals, native_amount, external_amount, _ in pool_setup}
        env.setup_liquidity_pools_simple(pool_definitions)

        # Enable margin on all pools
        mtp_enabled_pools = set(denom for denom, _, _, _, _ in pool_setup)
        margin_params_before = env.gridnoded.query_margin_params()
        env.gridnoded.tx_margin_update_pools(env.clp_admin, mtp_enabled_pools, [], broadcast_mode="block")
        margin_params_after = env.gridnoded.query_margin_params()
        assert len(margin_params_before["params"]["pools"]) == 0
        assert set(margin_params_after["params"]["pools"]) == mtp_enabled_pools

        yield env, env.gridnoded, pool_definitions
        env.close()

    def test_swap_fury_to_external(self):
        with self.with_test_env(self.default_pool_setup) as (env, gridnoded, pools_definitions):
            src_denom = "fury"
            dst_denom = "ceth"
            swap_amount = 10**18
            assert dst_denom in pools_definitions

            account = gridnoded.create_addr()
            env.fund(account, {FURY: 10**25 + swap_amount})
            balance_before = gridnoded.get_balance(account)
            pool_before = gridnoded.query_pools_sorted()[dst_denom]
            gridnoded.tx_clp_swap(account, src_denom, swap_amount, dst_denom, 0, broadcast_mode="block")
            balance_after = gridnoded.get_balance(account)
            pool_after = gridnoded.query_pools_sorted()[dst_denom]

            src_balance_delta = balance_after.get(src_denom, 0) - balance_before.get(src_denom, 0)
            dst_balance_delta = balance_after.get(dst_denom, 0) - balance_before.get(dst_denom, 0)

            assert balance_before.get(dst_denom, 0) == 0
            assert balance_after.get(dst_denom, 0) > 0

            # Account fury balance should change by 1 tx fee + amount swapped
            assert src_balance_delta == - (gridchain.grid_tx_fee_in_fury + swap_amount)

            # Pool native asset balance should increase by swapped amount
            assert int(pool_after["native_asset_balance"]) - int(pool_before["native_asset_balance"]) == swap_amount

            # Pool external asset balance should decrease by amount received
            assert int(pool_after["external_asset_balance"]) - int(pool_before["external_asset_balance"]) == -dst_balance_delta

            # y = x*Y / (x + X)
            # x = y*X / (y + Y)
            # x = swap input or output (non-fury)
            # y = swap input or output (fury)
            # X = non-fury token pool balance
            # Y = fury pool balance

            y = swap_amount
            X = int(pool_after[dst_denom]["external_asset_balance"])
            Y = int(pool_after[dst_denom]["native_asset_balance"])
            x = round(y*X / (y + Y))

            ok = x == dst_balance_delta
            ok_ratio = x / dst_balance_delta
            ok_difference = x - dst_balance_delta
            assert ok  # Fails, why?
            assert True  # no-op line just for setting a breakpoint

    def test_swap_external_to_fury(self):
        with self.with_test_env(self.default_pool_setup) as (env, gridnoded, pools_definitions):
            raise NotImplemented()  # TODO

    def test_swap_external_to_external(self):
        with self.with_test_env(self.default_pool_setup) as (env, gridnoded, pools_definitions):
            account = gridnoded.create_addr()
            balances = []
            pools = []

            src_denom = "cusdc"
            dst_denom = "ceth"
            swap_amount = 10**23

            assert src_denom in pools_definitions
            assert dst_denom in pools_definitions

            balances.append(gridnoded.get_balance(account))
            pools.append(gridnoded.query_pools_sorted())
            env.fund(account, {FURY: 10**20, src_denom: swap_amount})

            balance_before = gridnoded.get_balance(account)
            pools_before = gridnoded.query_pools_sorted()
            gridnoded.tx_clp_swap(account, src_denom, swap_amount, dst_denom, 0, broadcast_mode="block")
            balance_after = gridnoded.get_balance(account)
            pools_after = gridnoded.query_pools_sorted()

            fury_delta = balance_after.get(FURY, 0) - balance_before.get(FURY, 0)
            from_delta = balance_after.get(src_denom, 0) - balance_before.get(src_denom, 0)
            to_delta = balance_after.get(dst_denom, 0) - balance_before.get(dst_denom, 0)

            # Before swap, the account should have expected amount of src_denom and zero dst_denom
            assert balance_before.get(src_denom, 0) == swap_amount
            assert balance_before.get(dst_denom, 0) == 0

            # The account should pay a transaction fee in fury of 10**17
            assert fury_delta == -gridchain.grid_tx_fee_in_fury

            # Source pool's external asset should increase by swapped amount
            assert int(pools_after[src_denom]["external_asset_balance"]) - int(pools_before[src_denom]["external_asset_balance"]) == swap_amount

            # Destination pool's external asset should decrease by amount received
            assert int(pools_after[dst_denom]["external_asset_balance"]) - int(pools_before[dst_denom]["external_asset_balance"]) == -to_delta

            src_pool_native_balance_delta = int(pools_after[src_denom]["native_asset_balance"]) - int(pools_before[src_denom]["native_asset_balance"])
            dst_pool_native_balance_delta = int(pools_after[dst_denom]["native_asset_balance"]) - int(pools_before[dst_denom]["native_asset_balance"])

            # Source pool native balance should decrease
            assert src_pool_native_balance_delta < 0

            # Destination pool  native balance should increase by the same amount
            assert dst_pool_native_balance_delta == -src_pool_native_balance_delta

            pools.append(pools_before)
            pools.append(pools_after)
            balances.append(balance_before)
            balances.append(balance_after)
            assert True  # no-op line just for setting a breakpoint

    def test_swap(self):
        # self.test_swap_fury_to_external()
        # self.test_swap_external_to_external()
        self.test_margin()

    def test_margin(self):
        borrow_asset = "fury"
        collateral_asset = "cusdc"
        collateral_amount = 10**20
        leverage = 2

        with self.with_test_env(self.default_pool_setup) as (env, gridnoded, pools_definitions):
            account = gridnoded.create_addr()
            env.fund(account, {
                FURY: 10**25,
                collateral_asset: 10**25,
            })
            margin_params = gridnoded.query_margin_params()

            # gridnoded.tx_margin_whitelist(env.clp_admin, account, broadcast_mode="block")

            pool_before_open = gridnoded.query_pools_sorted()[collateral_asset]
            balance_before_open = gridnoded.get_balance(account)
            mtp_positions_before_open = gridnoded.query_margin_positions_for_address(account)
            res = gridnoded.margin_open_simple(account, borrow_asset, collateral_asset=collateral_asset,
                collateral_amount=collateral_amount, leverage=leverage, position="long")
            mtp_id = int(res["id"])
            pool_after_open = gridnoded.query_pools_sorted()[collateral_asset]
            balance_after_open = gridnoded.get_balance(account)
            mtp_positions_after_open = gridnoded.query_margin_positions_for_address(account)

            assert len(mtp_positions_before_open) == 0
            assert len(mtp_positions_after_open) == 1

            open_borrow_delta = balance_after_open.get(borrow_asset, 0) - balance_before_open.get(borrow_asset, 0)
            open_collateral_delta = balance_after_open.get(collateral_asset, 0) - balance_before_open.get(collateral_asset, 0)

            # TODO Why does the open position disappear after 4 blocks?
            # Whitelisting does not help
            for i in range(10):
                cnt = len(gridnoded.query_margin_positions_for_address(account))
                if cnt == 0:
                    break
                gridnoded.wait_for_last_transaction_to_be_mined()

            # TODO

            pool_before_close = gridnoded.query_pools_sorted()[collateral_asset]
            balance_before_close = gridnoded.get_balance(account)
            res2 = gridnoded.tx_margin_close(account, mtp_id, broadcast_mode="block")
            pool_after_close = gridnoded.query_pools_sorted()[collateral_asset]
            balance_after_close = gridnoded.get_balance(account)

            assert True
