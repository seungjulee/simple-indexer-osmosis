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
	curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"address": "1F7249F418B90714BF52797336B771B5AD467533"}' \
    http://localhost:8080/twirp/rpc.SimpleOsmosisExplorer/GetBlocksByProposer

run_example_get_number_of_txs_in_last_n_blocks:
	curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"n": 4}' \
    http://localhost:8080/twirp/rpc.SimpleOsmosisExplorer/GetNumberOfTXsInLastNBlocks

run_example_get_top_n_peers_by_score_in_last_n_blocks:
	curl --request "POST" \
    --header "Content-Type: application/json" \
    --data '{"n_peer": 3, "n_block": 3}' \
    http://localhost:8080/twirp/rpc.SimpleOsmosisExplorer/GetTopNPeersByScoreInLastNBlocks

