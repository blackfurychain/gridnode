#!/bin/bash 

addr=$1
shift

gridnoded q auth account ${addr:=${VALIDATOR1_ADDR}} -o json | jq
