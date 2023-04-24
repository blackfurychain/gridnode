#!/bin/bash
pid_gridnode=$(ps aux | grep "gridnoded start" | grep -v grep | awk '{print $2}')
# pid_rest=$(ps aux | grep "gridnoded rest-server" | grep -v grep | awk '{print $2}')

if [[ ! -z "$pid_gridnode" ]]; then 
  kill -9 $pid_gridnode
fi

# if [[ ! -z "$pid_rest" ]]; then 
#   kill -9 $pid_rest
# fi