package main

import (
	"awsproj/internal/hashes/hashesDelivery"
	"awsproj/internal/hashes/hashesUseCase"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = ":8080"
)
func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	hasherUC := hashesDelivery.NewHashesDelivery(&hashesUseCase.HashesUseCase{})

	hashesDelivery.RegisterHasherServer(s, hasherUC)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
