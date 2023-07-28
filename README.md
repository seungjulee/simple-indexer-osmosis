# Overview

## Architecture

## Demo

## Prerequisite
1. sqlite
2. twirp
3. protobuf

## Installation


## Running it locally


## TODO
### Infra
1. Use Websocket to subscribe to new blocks and transactions
2. If any block height is missing, fetch and index the missing block

### Feature
1. Separate out the schema for transaction. Include more properties
2. Better scoring system for net info. Look into https://github.com/nervosnetwork/rfcs/blob/master/rfcs/0007-scoring-system-and-network-security/0007-scoring-system-and-network-security.md https://github.com/bro-n-bro/bro_rating/blob/main/README.md

## Development

### Indexer

### RPC / API Server

### Database

## Example Query
```sh
curl --request "POST" \
--header "Content-Type: application/json" \
--data '{"address": "1F7249F418B90714BF52797336B771B5AD467533"}' \
http://localhost:8080/twirp/rpc.SimpleOsmosisExplorer/GetBlocksByProposer

{"blocks":[{"block_id":"","hash":"23C9F0C97350B5F241668872F38745D98155733D27AC0B80A71709D2FA987F13","height":"10724728","proposer_address":"1F7249F418B90714BF52797336B771B5AD467533","time":"2023-07-27T15:51:11.958921266Z","txs":"","num_txs":"5"}]}
```