#!/bin/bash 

addr=$1
shift

grided q auth account ${addr:=${VALIDATOR1_ADDR}} -o json | jq
