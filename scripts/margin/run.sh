#!/usr/bin/env bash

set -x

killall gridnoded

cd ../..
make install
gridnoded start --trace
