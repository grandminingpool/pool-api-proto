//go:generate protoc --proto_path=pool_api --go_out=./pool_api --go_opt=paths=source_relative --go-grpc_out=./pool_api --go-grpc_opt=paths=source_relative pool_api.proto

package proto
