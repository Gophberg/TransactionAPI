package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/Gophberg/TransactionAPI/internal/app/TransactionAPI/pb"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
	"time"
)

var (
	port = flag.Int("port", 50051, "The server port")
	//sleepDuration int = 20
)

type server struct {
	pb.TransactionServer
}

func (s *server) Transaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	ctx.Deadline()
	msg := fmt.Sprintf("UserID: %s\nUser Email: %s\nCurrency: %s\nAmount: %v\n",
		strconv.Itoa(int(in.GetUserID())),
		in.GetUserEmail(),
		in.GetCurrency(),
		in.GetAmount(),
	)
	log.Printf("Received transaction data:\n%v\n", msg)
	log.Println("Processing transaction...")
	status := doTransaction(in)
	log.Println("Transaction complete with status", status)
	return &pb.TransactionResponse{Status: status}, nil
}

func doTransaction(in *pb.TransactionRequest) bool {
	amount := decimal.NewFromFloat(in.Amount)
	log.Println("Received funds", amount)
	time.Sleep(time.Second * 5) // processing transaction
	if in.UserEmail == "joe@mail.edu" {
		log.Println("I hate him")
		return false
	}
	if in.Amount <= 0 {
		log.Println("Low amount")
		return false
	}
	return true
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
