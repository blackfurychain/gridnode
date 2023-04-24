#!/usr/bin/env bash

cosmovisor start >> gridnode.log 2>&1  &
sleep 10
yes Y | gridnoded tx gov submit-proposal software-upgrade release-20210414000000 --from grid --deposit 100000000stake --upgrade-height 20 --title release-20210414000000 --description release-20210414000000
sleep 5
yes Y | gridnoded tx gov vote 1 yes --from grid --keyring-backend test --chain-id localnet
clear
sleep 5
gridnoded query gov proposal 1

#--info '{"binaries":{"linux/amd64":"https://srv-store2.gofile.io/download/K9xJtY/gridnoded.zip?checksum=sha256:8630d1e36017ca680d572926d6a4fc7fe9a24901c52f48c70523b7d44ad0cfb2"}}'