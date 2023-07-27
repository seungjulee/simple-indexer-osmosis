gen:
	protoc --proto_path=. --twirp_out=. --go_out=. rpc/*.proto