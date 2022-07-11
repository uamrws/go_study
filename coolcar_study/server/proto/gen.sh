protoc -I . --go-grpc_out ./gen/go --go-grpc_opt paths=source_relative trip.proto
protoc -I . --go_out ./gen/go --go_opt paths=source_relative trip.proto
protoc -I . --grpc-gateway_out ./gen/go --grpc-gateway_opt paths=source_relative --grpc-gateway_opt grpc_api_configuration=trip.yaml trip.proto