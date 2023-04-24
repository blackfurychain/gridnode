# Minimum swap fees

This is a proposal to set a minimum swap fee.

## Current swap behaviour

Due to ratio shifting the current swap behaviour depends on whether the user is buying or selling Fury.

In the following formulas:

```
X - input depth (balance + liabilities)
Y - output depth (balance + liabilities)
x - input amount
y - output amount
r - current ratio shifting running rate
f - swap fee rate, this must satisfy 0 =< f =< 1
```

### Swapping to Fury:

```
y = (1 - f) * x * Y / ((x + X) * (1 + r))
fee = f * x * Y / ((x + X) * (1 + r))
```

Equivalently this can be written as:

```
raw_XYK_output = x * Y / (x + X)
adjusted_output = raw_XYK_output / (1 + r)

(1) fee = f * adjusted_output
y = adjusted_output - fee
```

### Swapping from Fury:

```
y = (1 - f) * (1 + r) * x * Y / (x + X)
fee = f * (1 + r) * x * Y / (x + X)
```

Similar to the case of swapping to fury, this can be written as:

```
raw_XYK_output = x * Y / (x + X)
adjusted_output = raw_XYK_output * (1 + r)

(2) fee = f * adjusted_output
y = adjusted_output - fee
```

## Proposed Change

Apply a minimum fee when swapping.

The fee calculation in equation (1) and (2) becomes:

```
fee = min(max(f * adjusted_output, min_fee), adjusted_output)
```

Where `min_fee` is a minimum fee parameter for the token being bought, which is set via an admin key. See CLI
section for more details.

The min function is required to ensure that the fee is not greater than the adjusted output.

If a `min-fee` has not been set for a token then it defaults to zero.

## Events

There are no new events or updates to existing events.

## CLI

The CLI option for querying the swap fee rate (`gridnoded q clp swap-fee-rate`) and setting the swap fee
rate (`gridnoded tx clp set-swap-fee-rate`), must be renamed to `gridnoded q clp swap-fee-params`
and `gridnoded tx clp set-swap-fee-params` and updated to include the min fee.

### Setting

The CLI should validate that the min fees are valid cosmos Uint256.

```bash
gridnoded tx clp set-swap-fee-params \
  --from grid \
  --path ./swap-fee-params.json \
  --keyring-backend test \
  --fees 100000000000000000fury \
  --chain-id localnet \
  -y
```

```json
{
	"swap_fee_rate": "0.003",
	"token_params": [{
			"asset": "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2",
			"min_swap_fee": "12"
		},
		{
			"asset": "cusdc",
			"min_swap_fee": "800"
		},
		{
			"asset": "fury",
			"min_swap_fee": "12"
		}
	]
}
```

### Querying

```bash
gridnoded q clp swap-fee-params --output json
```

```json
{
	"swap_fee_rate": "0.003000000000000000",
	"token_params": [{
			"asset": "ibc/27394FB092D2ECCD56123C74F36E4C1F926001CEADA9CA97EA622B25F41E5EB2",
			"min_swap_fee": "12"
		},
		{
			"asset": "cusdc",
			"min_swap_fee": "800"
		},
		{
			"asset": "fury",
			"min_swap_fee": "12"
		}
	]
}
```
