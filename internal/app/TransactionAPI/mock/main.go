package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "github.com/Gophberg/TransactionAPI/internal/app/TransactionAPI/pb"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.TransactionServer
}

func (s *server) Transaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Printf("Received: %v", in.GetUserEmail())
	return &pb.TransactionResponse{Message: "Hello " + in.GetUserEmail() + in.GetCurrency()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTransactionServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
