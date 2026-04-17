package common

//go:generate protoc --go_out=plugins=grpc:genproto/trainer -I ../../api/protobuf ../../api/protobuf/trainer.proto
//go:generate protoc --go_out=plugins=grpc:genproto/users -I ../../api/protobuf ../../api/protobuf/users.proto
