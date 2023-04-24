#!/usr/bin/env bash

cp $GOPATH/src/new/grided $GOPATH/bin/
cosmovisor start >> gridnode.log 2>&1  &