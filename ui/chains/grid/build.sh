#!/bin/bash

killall gridnoded

rm $(which gridnoded) 2> /dev/null || echo gridnoded not install yet ...

rm -rf ~/.gridnoded

cd ../../../ && make install 