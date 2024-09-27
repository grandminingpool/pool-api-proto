//go:generate protoc --proto_path=proto/utils --go_out=./generated/utils/filters --go_opt=paths=source_relative filters.proto
//go:generate protoc --proto_path=proto/utils --go_out=./generated/utils/sorts --go_opt=paths=source_relative sorts.proto
//go:generate protoc --proto_path=proto/utils --go_out=./generated/utils/pagination --go_opt=paths=source_relative pagination.proto
//go:generate protoc --proto_path=proto --go_out=./generated/pool --go_opt=paths=source_relative --go-grpc_out=./generated/pool --go-grpc_opt=paths=source_relative pool.proto
//go:generate protoc --proto_path=proto --go_out=./generated/pool_miners --go_opt=paths=source_relative --go-grpc_out=./generated/pool_miners --go-grpc_opt=paths=source_relative pool_miners.proto
//go:generate protoc --proto_path=proto --go_out=./generated/pool_payouts --go_opt=paths=source_relative --go-grpc_out=./generated/pool_payouts --go-grpc_opt=paths=source_relative pool_payouts.proto
//go:generate protoc --proto_path=proto --go_out=./generated/charts --go_opt=paths=source_relative --go-grpc_out=./generated/charts --go-grpc_opt=paths=source_relative charts.proto

package proto
