#!/bin/bash

killall grided

rm $(which grided) 2> /dev/null || echo grided not install yet ...

rm -rf ~/.grided

cd ../../../ && make install 