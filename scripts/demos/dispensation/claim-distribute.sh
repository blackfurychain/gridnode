


# Multisig Key - It is a key composed of two or more keys (N) , with a signing threshold (K) ,such that the transaction needs K out of N votes to go through.
gridnoded tx dispensation claim ValidatorSubsidy --from akasha --keyring-backend test --yes --chain-id localnet -o json
gridnoded tx dispensation claim ValidatorSubsidy --from grid --keyring-backend test --yes --chain-id localnet -o json
# create airdrop
# mkey = multisig key
# ar1 = name for airdrop , needs to be unique for every airdrop . If not the tx gets rejected
# input.json list of funding addresses  -  Input address must be part of the multisig key
# output.json list of airdrop receivers.
sleep 8
gridnoded q dispensation claims-by-type ValidatorSubsidy -o json
sleep 8
gridnoded tx dispensation create ValidatorSubsidy output.json grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd --gas 200064128 --from grid1syavy2npfyt9tcncdtsdzf7kny9lh777yqc2nd --keyring-backend test --fees 100000fury --yes --chain-id gridchain-devnet-042 --node tcp://rpc-devnet-042.gridchain.finance:80

sleep 8
gridnoded q dispensation distributions-all --chain-id localnet -o json
#gridnoded q dispensation records-by-name-all ar1 >> all.json
#gridnoded q dispensation records-by-name-pending ar1 >> pending.json
#gridnoded q dispensation records-by-name-completed ar1 >> completed.json
#gridnoded q dispensation records-by-addr grid1cp23ye3h49nl5ty35vewrtvsgwnuczt03jwg00


