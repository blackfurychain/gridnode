#!/bin/bash


# This will run our stack from genesis
# All setup scripts will be run and once it is ready our app will be compiled and built
# You probably only need to use this script to set a new snapshot archive

killall gridnoded ebrelayer ganache-cli
sleep 5

./scripts/_grid-build.sh

yarn concurrently -k -r -s first \
  "./scripts/_eth.sh" \
  "./scripts/_grid.sh" \
  "yarn wait-on http-get://localhost:1317/node_info && ./scripts/_migrate.sh && ./scripts/_peggy.sh"
