#!/bin/bash
pid_gridnode=$(ps aux | grep "grided start" | grep -v grep | awk '{print $2}')
# pid_rest=$(ps aux | grep "grided rest-server" | grep -v grep | awk '{print $2}')

if [[ ! -z "$pid_gridnode" ]]; then 
  kill -9 $pid_gridnode
fi

# if [[ ! -z "$pid_rest" ]]; then 
#   kill -9 $pid_rest
# fi