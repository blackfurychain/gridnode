# DISPENSATION MODULE (Distribution)

## Overview
- The module allows a user to create a Distribution which can be of Type [Airdrop/LiquidityMining/ValidatorSubsidy]. 
- It accepts an input and output list .
- This transaction needs to signed by at-least all addresses of the input list ( can be set to more ,but not less)
- The module accumulates  funds from  the input address list and distributes it among the output list .
- The records are created in the same block.
- The distribution process starts in the next block with 10 distributions per block.



## Technicals 
### Data structures
 - The base-level data structure is 
```go
package records
type DistributionType int64

const Airdrop DistributionType = 1
const LiquidityMining DistributionType = 2
const ValidatorSubsidy DistributionType = 3

type Distribution struct {
    DistributionType DistributionType `json:"distribution_type"`
    DistributionName string           `json:"distribution_name"`
}
```
This is stored in the keeper with the key DistributionType_DistributionName for historical records. Therefore, the combination of type and name needs to be unique

- Distribution records are created for processing individual transfers to recipients

```go
package records

type DistributionStatus int64

const Pending DistributionStatus = 1
const Completed DistributionStatus = 2

type DistributionRecord struct {
	DistributionStatus          DistributionStatus `json:"distribution_status"`
	DistributionName            string             `json:"distribution_name"`
	DistributionType            DistributionType   `json:"distribution_type"`
	RecipientAddress            sdk.AccAddress     `json:"recipient_address"`
	Coins                       sdk.Coins          `json:"coins"`
	DistributionStartHeight     int64              `json:"distribution_start_height"`
	DistributionCompletedHeight int64              `json:"distribution_completed_height"`
}
```
This record is also stored in the keeper for historical records .

### High Level Flow
- After the sanity checks are cleared , the program iterates over the input addresses and sends all funds from these address to a module account.
- The program iterates over the output addresses and creates individual records for them in the keeper .( This design would need to be changed in the future to save gas costs)
- In case of type LiquidityMining or ValidatorSubsidy the program also checks is the associated claim for the record is present .
- The begin block iterates over these records and completes 10 records per block .
- Complete refers to sending the specified amount from the  module account to the recipient.
- In case of type LiquidityMining or ValidatorSubsidy the program also deletes the associated claim.


### User flow 
 The set of user commands to use this module 
```shell
#Multisig Key - It is a key composed of two or more keys (N) , with a signing threshold (K) ,such that the transaction needs K out of N votes to go through.

#create airdrop
#mkey        : multisig key
#ar1         : name of airdrop , needs to be unique for every airdrop. If not the tx gets rejected
#input.json  : list of funding addresses  -  Input address must be part of the multisig key
#output.json : list of airdrop receivers.

gridnoded tx dispensation create mkey ar1 input.json output.json --gas 200064128 --generate-only >> offlinetx.json

#First user signs
gridnoded tx sign --multisig $(gridnoded keys show mkey -a) --from $(gridnoded keys show grid -a)  offlinetx.json >> sig1.json

#Second user signs
gridnoded tx sign --multisig $(gridnoded keys show mkey -a) --from $(gridnoded keys show akasha -a)  offlinetx.json >> sig2.json

#Multisign created from the above signatures
gridnoded tx multisign offlinetx.json mkey sig1.json sig2.json >> signedtx.json

#transaction broadcast , distribution happens
gridnoded tx broadcast signedtx.json
```

### Events Emitted 
Transfer events are emitted for each transfer . There are two type of transfers in a distribution
- Transfer for address in the input list to the Dispensation Module Address.

```json
{
  "type": "transfer",
  "attributes": [
    {
      "key": "recipient",
      "value": "grid1zvwfuvy3nh949rn68haw78rg8jxjevgm2c820c"
    },
    {
      "key": "sender",
      "value": "grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd"
    },
    {
      "key": "amount",
      "value": "15000000000000000000fury"
    }
  ]
}
```
- Transfer of funds from the Dispensation Module Address to recipients in the output list
```json
{
  "type": "transfer",
  "attributes": [
    {
      "key": "recipient",
      "value": "grid1p6z0ze9mztfd8cx5z9g6pndmzdrtxnsfesnn97"
    },
    {
      "key": "sender",
      "value": "grid1zvwfuvy3nh949rn68haw78rg8jxjevgm2c820c"
    },
    {
      "key": "amount",
      "value": "10000000000000000000fury"
    }
  ]
}
```


- A distribution started event is emitted in the block in which the distribution is created .The distribution process starts from the next block
```json
 {
  "type": "distribution_started",
  "attributes": [
    {
      "key": "module_account",
      "value": "grid1zvwfuvy3nh949rn68haw78rg8jxjevgm2c820c"
    }
  ]
}
```


### Queries supported
```shell
#Query all distributions
gridnoded q dispensation distributions-all
#Query all distribution records by distribution name 
gridnoded q dispensation records-by-name-all ar1
#Query pending distribution records by distribution name 
gridnoded q dispensation records-by-name-pending ar1
#Query completed distribution records by distribution name
gridnoded q dispensation records-by-name-completed ar1
#Query distribution records by address
gridnoded q dispensation records-by-addr grid1cp23ye3h49nl5ty35vewrtvsgwnuczt03jwg00
```
