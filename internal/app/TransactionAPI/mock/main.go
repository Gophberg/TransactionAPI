package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Gophberg/TransactionAPI/internal/app/TransactionAPI/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

var (
	port              = flag.Int("port", 50051, "The server port")
	sleepDuration int = 20
)

type server struct {
	pb.TransactionServer
}

func (s *server) Transaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	msg, err := fmt.Printf("UserID: %s\nUser Email: %s\nCurrency: %s\nAmount: %d.%d\n",
		strconv.Itoa(int(in.GetUserID())),
		in.GetUserEmail(),
		in.GetCurrency(),
		in.GetAmount().Units,
		in.GetAmount().Nanos,
	)
	if err != nil {
		log.Println(err)
	}
	log.Println(msg)
	//str := strconv.Itoa(msg)
	log.Println("Doing some work")
	//time.Sleep(time.Second * 1)
	log.Println("Some work is done")
	//return &pb.TransactionResponse{Message: "Received request " + str}, nil
	return &pb.TransactionResponse{Message: "Hello " + in.GetUserEmail() + strconv.Itoa(int(in.GetAmount().Units))}, nil
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
