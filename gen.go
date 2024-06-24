//go:generate protoc --proto_path=proto --go_out=./generated/pool --go_opt=paths=source_relative --go-grpc_out=./generated/pool --go-grpc_opt=paths=source_relative pool.proto
//go:generate protoc --proto_path=proto --go_out=./generated/pool_miners --go_opt=paths=source_relative --go-grpc_out=./generated/pool_miners --go-grpc_opt=paths=source_relative pool_miners.proto
//go:generate protoc --proto_path=proto --go_out=./generated/pool_payouts --go_opt=paths=source_relative --go-grpc_out=./generated/pool_payouts --go-grpc_opt=paths=source_relative pool_payouts.proto

package proto
