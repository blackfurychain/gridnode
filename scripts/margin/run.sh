#!/usr/bin/env bash

set -x

killall grided

cd ../..
make install
grided start --trace
