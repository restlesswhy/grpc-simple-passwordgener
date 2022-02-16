prot:
	protoc --go_out=. --go_opt=paths=source_relative proto/passwordservice.proto

buf: prot
	protoc --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/passwordservice.proto