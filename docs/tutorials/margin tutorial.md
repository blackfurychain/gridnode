# Gridironchain - Margin Trading Basics Tutorial

#### demo video

TBD

#### Previous tutorial

- Clp Basics: https://github.com/Gridironchain/gridnode/blob/develop/docs/tutorials/clp%20tutorial.md

#### Dependencies:

0. `git clone git@github.com:Gridironchain/gridnode.git`
1. `cd gridnode`
2. `git checkout feature/margin-1`
3. `make install`

#### What is Margin Trading or trading with margin or leverage

Margin trading refers to the use of borrowed funds from continuous liquidity pools (CLP) providers
to trade a financial asset. The margin trader can bet against an asset to go up (long) or down (short) and
relies upon the collateral amount to form the loan.
At the same time providing a continous funding interest rate to the CLP providers.

#### Setup

1. Initialize the local chain run;

```bash
./scripts/init.sh
```

2. Decrease the gouvernance voting period time;
```bash
echo "$(jq '.app_state.gov.voting_params.voting_period = "60s"' $HOME/.gridnoded/config/genesis.json)" > $HOME/.gridnoded/config/genesis.json
```

3. Start the chain;
```bash
./scripts/run.sh
```

4. Check to see you have two local accounts/keys setup;
```bash
gridnoded keys list --keyring-backend=test
```

Result:
```
- name: akasha
  type: local
  address: grid1l7hypmqk2yc334vc6vmdwzp5sdefygj2ad93p5
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A0mB4PyE5XeS3sNpFXIX536INyNoJHkMu1DEQ8FgH8Mq"}'
  mnemonic: ""
- name: mkey
  type: multi
  address: grid1kkdqp4dtqmc7wh59vchqr0zdzk8w2ydukjugkz
  pubkey: '{"@type":"/cosmos.crypto.multisig.LegacyAminoPubKey","threshold":2,"public_keys":[{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AvUEsFHbsr40nTSmWh7CWYRZHGwf4cpRLtJlaRO4VAoq"},{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A0mB4PyE5XeS3sNpFXIX536INyNoJHkMu1DEQ8FgH8Mq"}]}'
  mnemonic: ""
- name: grid
  type: local
  address: grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AvUEsFHbsr40nTSmWh7CWYRZHGwf4cpRLtJlaRO4VAoq"}'
  mnemonic: ""
```

5. Check your seed account balance/s;
```bash
gridnoded q bank balances $(gridnoded keys show grid -a --keyring-backend=test)
```
```bash
gridnoded q bank balances $(gridnoded keys show akasha -a --keyring-backend=test)
```

#### Create and query pools

note:

- the minimum threshold for native amount is 10^18 fury.
- the minimum transaction fee for these operations is 10^17 fury.

1. Create the first pool for ceth;
```bash
gridnoded tx clp create-pool \
  --from grid \
  --keyring-backend test \
  --symbol ceth \
  --nativeAmount 2000000000000000000 \
  --externalAmount 2000000000000000000 \
  --fees 100000000000000000fury \
  --chain-id localnet \
  -y
```

2. Create another pool for cdash with a different account;
```bash
gridnoded tx clp create-pool \
  --from akasha \
  --keyring-backend test \
  --symbol cdash \
  --nativeAmount 3000000000000000000 \
  --externalAmount 3000000000000000000 \
  --fees 100000000000000000fury \
  --chain-id localnet \
  -y
```

3. Query all clp pools;
```bash
gridnoded q clp pools
```

#### Enable margin on pools

The set of pools that have margin enabled is managed through governance.

The param change proposal takes the format:

```json
{
  "title": "Margin Pools Param Change",
  "description": "Update enabled margin pools",
  "changes": [
    {
      "subspace": "margin",
      "key": "Pools",
      "value": ["ceth", "cdash"]
    }
  ],
  "deposit": "10000000stake"
}
```

1. Save the proposal above within a file named `proposal.json`

2. Submit a param change proposal;
```bash
gridnoded tx gov submit-proposal param-change proposal.json --from grid --keyring-backend test --chain-id localnet -y
```

3. Vote on proposal;
```bash
gridnoded tx gov vote 1 yes --from grid --chain-id localnet --keyring-backend test -y
```

4. Query the proposal to check the proposal status has passed;
```bash
gridnoded q gov proposals --chain-id localnet
```

Result:
```
pagination:
  next_key: null
  total: "0"
proposals:
- content:
    '@type': /cosmos.params.v1beta1.ParameterChangeProposal
    changes:
    - key: Pools
      subspace: margin
      value: |-
        [
                "ceth",
                "cdash"
              ]
    description: Update enabled margin pools
    title: Margin Pools Param Change
  deposit_end_time: "2022-02-09T18:50:23.040643413Z"
  final_tally_result:
    abstain: "0"
    "no": "0"
    no_with_veto: "0"
    "yes": "1000000000000000000000000"
  proposal_id: "1"
  status: PROPOSAL_STATUS_PASSED
  submit_time: "2022-02-07T18:50:23.040643413Z"
  total_deposit:
  - amount: "10000000"
    denom: stake
  voting_end_time: "2022-02-07T18:51:23.040643413Z"
  voting_start_time: "2022-02-07T18:50:23.040643413Z"
```

5. Verify that the margin param has changed;
```bash
gridnoded q params subspace margin Pools --chain-id localnet
```

Result:
```
key: Pools
subspace: margin
value: '["ceth","cdash"]'
```

#### Create and query margin trading positions (MTP)

1. Create margin long position against ceth;
```bash
gridnoded tx margin open \
  --from grid \
  --keyring-backend test \
  --borrow_asset ceth \
  --collateral_asset fury \
  --collateral_amount 1000 \
  --position long \
  --chain-id localnet \
  -y
```

Result:
```
code: 0
codespace: ""
data: ""
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 08CF11E1DECB1FF9340933D2D178DC6EBE2EE7114825FA2955C54972845C6E59
```

2. Add up to an existing margin position by creating a second margin position for the same asset ceth;
```bash
gridnoded tx margin open \
  --from grid \
  --keyring-backend test \
  --borrow_asset ceth \
  --collateral_asset fury \
  --collateral_amount 500 \
  --position long \
  --chain-id localnet \
  -y
```

Result:
```
code: 0
codespace: ""
data: ""
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 97E7A90E3DB3956101F8C226AC8F369F7C403956C84A4830103EAB3A286701B6
```

3. Query all the existing margin positions (same asset ceth);
```bash
gridnoded q margin positions-for-address $(gridnoded keys show grid -a --keyring-backend=test)
```

Result:
```
mtps:
- address: grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd
  collateral_amount: "1000"
  collateral_asset: fury
  custody_amount: "4000"
  custody_asset: ceth
  id: "1"
  leverage: "1"
  liabilities_i: "0"
  liabilities_p: "1000"
  mtp_health: "0.100000000000000000"
  position: LONG
- address: grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd
  collateral_amount: "500"
  collateral_asset: fury
  custody_amount: "2000"
  custody_asset: ceth
  id: "2"
  leverage: "1"
  liabilities_i: "0"
  liabilities_p: "500"
  mtp_health: "0.100000000000000000"
  position: LONG
```

#### Reduce size and close existing margin positions

1. Reduce the size of an existing margin position for ceth by closing one of the existing MTPs;
```bash
gridnoded tx margin close \
  --from grid \
  --keyring-backend test \
  --id 2 \
  --chain-id localnet \
  -y
```

Result:
```
code: 0
codespace: ""
data: ""
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 083CCA8E8C0E6E60A83A53764CD15031F1794AE79A00D6CD1F9E60E43601A39C
```

2. Query remaining margin positions for ceth;
```bash
gridnoded q margin positions-for-address $(gridnoded keys show grid -a --keyring-backend=test)
```

Result:
```
mtps:
- address: grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd
  collateral_amount: "1000"
  collateral_asset: fury
  custody_amount: "4000"
  custody_asset: ceth
  id: "1"
  leverage: "1"
  liabilities_i: "0"
  liabilities_p: "1000"
  mtp_health: "0.100000000000000000"
  position: LONG
```

3. Close the margin long position entirely for ceth;
```bash
gridnoded tx margin close \
  --from grid \
  --keyring-backend test \
  --id 1 \
  --chain-id localnet \
  -y
```

Result:
```
code: 0
codespace: ""
data: ""
gas_used: "0"
gas_wanted: "0"
height: "0"
info: ""
logs: []
raw_log: '[]'
timestamp: ""
tx: null
txhash: 110C1CF8DDE40A1554D500AE584CBF8875209908A8D7792256EF9486B2F84B70
```

4. Query existing margin positions (none);
```bash
gridnoded q margin positions-for-address $(gridnoded keys show grid -a --keyring-backend=test)
```

Result:
```
mtps: []
```
