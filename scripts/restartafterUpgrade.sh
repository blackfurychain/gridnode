#!/usr/bin/env bash

cp $GOPATH/src/new/gridnoded $GOPATH/bin/
cosmovisor start >> gridnode.log 2>&1  &