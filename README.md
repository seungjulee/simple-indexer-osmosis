Prerequisite
1. sqlite
2. twirp
3. protobuf

Installation


```sh
curl --request "POST" \
--header "Content-Type: application/json" \
--data '{"address": "1F7249F418B90714BF52797336B771B5AD467533"}' \
http://localhost:8080/twirp/rpc.SimpleOsmosisExplorer/GetBlocksByProposer

{"blocks":[{"block_id":"","hash":"23C9F0C97350B5F241668872F38745D98155733D27AC0B80A71709D2FA987F13","height":"10724728","proposer_address":"1F7249F418B90714BF52797336B771B5AD467533","time":"2023-07-27T15:51:11.958921266Z","txs":"","num_txs":"5"}]}
```