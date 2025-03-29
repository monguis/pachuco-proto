protoc \
 --go_out=. --go_opt=paths=source_relative \
 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 --grpc-gateway_out=. --grpc-gateway_opt paths=source_relative \
 dice-throws_protos/dice-throws.proto

protoc \
 --go_out=. --go_opt=paths=source_relative \
 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
 --grpc-gateway_out=. --grpc-gateway_opt paths=source_relative \
 user_protos/users.proto
