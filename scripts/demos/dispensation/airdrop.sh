


# Multisig Key - It is a key composed of two or more keys (N) , with a signing threshold (K) ,such that the transaction needs K out of N votes to go through.

# create airdrop
# mkey = multisig key
# ar1 = name for airdrop , needs to be unique for every airdrop . If not the tx gets rejected
# input.json list of funding addresses  -  Input address must be part of the multisig key
# output.json list of airdrop receivers.
grided tx dispensation create ValidatorSubsidy output.json did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92 --from did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92 --yes --fees 150000fury --chain-id=localnet --keyring-backend=test
grided tx dispensation run 29_did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92 ValidatorSubsidy--from did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92 --yes --fees 150000fury --chain-id=localnet --keyring-backend=test
sleep 8
grided q dispensation distributions-all -chain-id localnet
#grided q dispensation records-by-name-all ar1 >> all.json
#grided q dispensation records-by-name-pending ar1 >> pending.json
#grided q dispensation records-by-name-completed ar1 >> completed.json
#grided q dispensation records-by-addr did:fury:g1cp23ye3h49nl5ty35vewrtvsgwnuczt03jwg00

grided tx dispensation create Airdrop output.json --gas 90128 --from $(grided keys show grid -a) --yes --broadcast-mode async --sequence 26 --account-number 3 --chain-id localnet
grided tx dispensation create Airdrop output.json --gas 90128 --from $(grided keys show grid -a) --yes --broadcast-mode async --sequence 27 --account-number 3 --chain-id localnet
grided tx dispensation run 25_did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92 ValidatorSubsidy --from did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92 --yes --gas auto --gas-adjustment=1.5 --gas-prices 1.0fury --chain-id=localnet --keyring-backend=test



