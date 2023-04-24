#!/bin/sh

CHAINNET0=gridironchain-ibc-0
CHAINNET1=gridironchain-ibc-1
IPADDR0=192.168.65.2
IPADDR1=192.168.65.3
IPADDR2=192.168.65.4
SUBNET=192.168.65.1/24

CHAINNET0=${CHAINNET0} \
CHAINNET1=${CHAINNET1} \
IPADDR0=${IPADDR0} \
IPADDR1=${IPADDR1} \
IPADDR2=${IPADDR2} \
SUBNET=${SUBNET} \
MNEMONIC="${MNEMONIC}" docker compose $1