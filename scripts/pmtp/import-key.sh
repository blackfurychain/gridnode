#!/usr/bin/env bash

set -x

echo ${ADMIN_MNEMONIC} | gridnoded keys add ${GRID_ACT} --recover --keyring-backend=test