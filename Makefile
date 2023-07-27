export GOBIN=$(shell pwd)/bin
export PATH := $(shell pwd)/bin:${PATH}

gen:
	go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

	protoc --proto_path=. --twirp_out=. --go_out=. rpc/*.proto