gen:
	@protoc \
	--proto_path=pb "pb/math.proto" \
	--go_out=pb/generated --go_opt=paths=source_relative \
	--go-grpc_out=pb/generated \
	--go-grpc_opt=paths=source_relative