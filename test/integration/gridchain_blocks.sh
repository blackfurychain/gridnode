height=$(grided --home $CHAINDIR/.grided q block | jq -r .block.header.height)
seq $height | parallel -k grided --home $CHAINDIR/.grided q block {}
