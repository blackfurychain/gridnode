


# Multisig Key - It is a key composed of two or more keys (N) , with a signing threshold (K) ,such that the transaction needs K out of N votes to go through.

# create airdrop
# mkey = multisig key
# ar1 = name for airdrop , needs to be unique for every airdrop . If not the tx gets rejected
# input.json list of funding addresses  -  Input address must be part of the multisig key
# output.json list of airdrop receivers.
gridnoded tx dispensation create ValidatorSubsidy output.json did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777vuvqz8 --from did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777vuvqz8 --yes --fees 150000fury --chain-id=localnet --keyring-backend=test
gridnoded tx dispensation run 29_did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777vuvqz8 ValidatorSubsidy--from did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777vuvqz8 --yes --fees 150000fury --chain-id=localnet --keyring-backend=test
sleep 8
gridnoded q dispensation distributions-all -chain-id localnet
#gridnoded q dispensation records-by-name-all ar1 >> all.json
#gridnoded q dispensation records-by-name-pending ar1 >> pending.json
#gridnoded q dispensation records-by-name-completed ar1 >> completed.json
#gridnoded q dispensation records-by-addr did:fury:g1cp23ye3h49nl5ty35vewrtvsgwnuczt03jwg00

gridnoded tx dispensation create Airdrop output.json --gas 90128 --from $(gridnoded keys show grid -a) --yes --broadcast-mode async --sequence 26 --account-number 3 --chain-id localnet
gridnoded tx dispensation create Airdrop output.json --gas 90128 --from $(gridnoded keys show grid -a) --yes --broadcast-mode async --sequence 27 --account-number 3 --chain-id localnet
gridnoded tx dispensation run 25_did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777vuvqz8 ValidatorSubsidy --from did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777vuvqz8 --yes --gas auto --gas-adjustment=1.5 --gas-prices 1.0fury --chain-id=localnet --keyring-backend=test



