package server

import (
	fmt "fmt"
	net "net"

	grpc "google.golang.org/grpc"

	protoAddress "github.com/jcsw/address-grpc-api/pkg/proto/address"

	address "github.com/jcsw/address-grpc-service/pkg/address"
	log "github.com/jcsw/address-grpc-service/pkg/system/log"
	properties "github.com/jcsw/address-grpc-service/pkg/system/properties"
)

// GrpcServer represents the gRPC server
type GrpcServer struct {
	grpc *grpc.Server
}

// Start start a gRPC server and waits for connection
func (s *GrpcServer) Start() {
	log.Info("p=server f=Start m=grpc_server_starting")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", properties.Values.ServerPort))
	if err != nil {
		log.Fatal("p=server f=Start m=failed_to_listen err=%v", err)
	}

	s.grpc = grpc.NewServer()

	protoAddress.RegisterAddressServiceServer(s.grpc, &address.Service{})

	log.Info("p=server f=Start m=grpc_server_listen_at_port_%d", properties.Values.ServerPort)
	if err := s.grpc.Serve(lis); err != nil {
		log.Fatal("p=server f=Start m=failed_to_create_server err=%v", err)
	}
}

// Stop stop a gRPC server
func (s *GrpcServer) Stop() {
	log.Info("p=server f=Stop m=grpc_server_stopping")

	s.grpc.Stop()
}
