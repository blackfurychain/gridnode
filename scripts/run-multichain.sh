#!/usr/bin/env bash

killall gridnoded 
killall hermes 

#gridnodecli rest-server &

echo "starting gridnode servers"
sleep 1
gridnoded start --home ~/.gridnode-1 --p2p.laddr 0.0.0.0:27655 --grpc.address 0.0.0.0:9090 --grpc-web.address 0.0.0.0:9093 --address tcp://0.0.0.0:27659 --rpc.laddr tcp://127.0.0.1:27665 >> abci_1.log 2>&1 &
sleep 1
gridnoded start --home ~/.gridnode-2 --p2p.laddr 0.0.0.0:27656 --grpc.address 0.0.0.0:9091 --grpc-web.address 0.0.0.0:9094 --address tcp://0.0.0.0:27660 --rpc.laddr tcp://127.0.0.1:27666 >> abci_2.log 2>&1  &
sleep 1
gridnoded start --home ~/.gridnode-3 --p2p.laddr 0.0.0.0:27657 --grpc.address 0.0.0.0:9092 --grpc-web.address 0.0.0.0:9095 --address tcp://0.0.0.0:27661 --rpc.laddr tcp://127.0.0.1:27667 >> abci_3.log 2>&1 &
sleep 5

hermes start > hermes.log 2>&1 &

#gridnoded q bank balances did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92 --node tcp://127.0.0.1:27665
#gridnoded q bank balances did:fury:g1l7hypmqk2yc334vc6vmdwzp5sdefygj23y4thn --node tcp://127.0.0.1:27666
#gridnoded q bank balances did:fury:g1l7hypmqk2yc334vc6vmdwzp5sdefygj23y4thn --node tcp://127.0.0.1:27667
#
#gridnoded tx ibc-transfer transfer transfer channel-0 did:fury:g1l7hypmqk2yc334vc6vmdwzp5sdefygj23y4thn 100fury --node tcp://127.0.0.1:27665 --chain-id=localnet-1 --from=akasha --log_level=debug --gas-prices=0.5fury --keyring-backend test  --home ~/.gridnode-1
#gridnoded tx ibc-transfer transfer transfer channel-2 did:fury:g1l7hypmqk2yc334vc6vmdwzp5sdefygj23y4thn 50ibc/E0263CEED41F926DCE9A805F0358074873E478B515A94DF202E6B69E29DA6178 --node tcp://127.0.0.1:27666 --chain-id=localnet-2 --from=akasha --log_level=debug --gas-prices=0.5fury --keyring-backend test  --home ~/.gridnode-2
#gridnoded tx ibc-transfer transfer transfer channel-0 did:fury:g1syavy2npfyt9tcncdtsdzf7kny9lh777gfgs92 50ibc/4C2B3D3B398FC7B8FFA3A96314006FF0B38E3BFC4CE90D8EE46E9EB6768A482D --node tcp://127.0.0.1:27666 --chain-id=localnet-2 --from=grid --log_level=debug --gas-prices=0.5fury --keyring-backend test  --home ~/.gridnode-2
#gridnoded tx ibc-transfer transfer transfer channel-1 did:fury:g1l7hypmqk2yc334vc6vmdwzp5sdefygj23y4thn 50ibc/5C3977A32007D22B1845B57076D0E27C3159C3067B11B9CEF6FA551D71DAEDD6 --node tcp://127.0.0.1:27667 --chain-id=localnet-3 --from=akasha --log_level=debug --gas-prices=0.5fury --keyring-backend test  --home ~/.gridnode-3
