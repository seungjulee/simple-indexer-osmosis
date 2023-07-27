export GOBIN=$(shell pwd)/bin
export PATH := $(shell pwd)/bin:${PATH}

gen:
	go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

	protoc --proto_path=. --twirp_out=. --go_out=. rpc/*.proto

run_indexer:
	go run ./cmd/indexer/main.go

run_server:
	go run ./cmd/server/main.go

run_example_get_block_by_proposer:
	# {"blocks":[{"block_id":"","hash":"23C9F0C97350B5F241668872F38745D98155733D27AC0B80A71709D2FA987F13","height":"10724728","proposer_address":"1F7249F418B90714BF52797336B771B5AD467533","time":"2023-07-27T15:51:11.958921266Z","txs":"","num_txs":"5"}]}
	curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"address": "1F7249F418B90714BF52797336B771B5AD467533"}' \
    http://localhost:8080/twirp/rpc.SimpleOsmosisExplorer/GetBlocksByProposer

