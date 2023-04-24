height=$(gridnoded --home $CHAINDIR/.gridnoded q block | jq -r .block.header.height)
seq $height | parallel -k gridnoded --home $CHAINDIR/.gridnoded q block {}
