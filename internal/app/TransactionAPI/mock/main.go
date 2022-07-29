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
	log.Printf("[EPS] Received transaction data:\n%v\n", msg)
	log.Println("[EPS] Processing transaction...")
	status, reason := doTransaction(in)
	log.Printf("[EPS] Transaction complete with status: '%t', by reason '%s'", status, reason)
	return &pb.TransactionResponse{
		Status: status,
		Reason: reason,
	}, nil
}

func doTransaction(in *pb.TransactionRequest) (bool, string) {
	amount := decimal.NewFromFloat(in.Amount)
	log.Println("[EPS] Received funds", amount)
	log.Println("[EPS] Starting do nothing")
	time.Sleep(time.Second * 14) // processing transaction
	log.Println("[EPS] Stopping do nothing")
	if in.UserEmail == "joe@mail.edu" {
		log.Println("[EPS] I hate him")
		return false, "User is bad"
	}
	if in.Amount <= 0 {
		log.Println("[EPS] Low amount")
		return false, "Low amount"
	}
	return true, "Success"
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("[EPS] failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterTransactionServer(s, &server{})
	log.Printf("[EPS] server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("[EPS] failed to serve: %v", err)
	}
}
