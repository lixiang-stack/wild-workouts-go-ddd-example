package common

//go:generate protoc --go_out=genproto/trainer --go_opt=paths=source_relative --go-grpc_out=genproto/trainer --go-grpc_opt=paths=source_relative -I ../../api/protobuf -I /usr/local/include ../../api/protobuf/trainer.proto
//go:generate protoc --go_out=genproto/users --go_opt=paths=source_relative --go-grpc_out=genproto/users --go-grpc_opt=paths=source_relative -I ../../api/protobuf -I /usr/local/include ../../api/protobuf/users.proto
