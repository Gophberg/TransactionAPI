package TransactionAPI

import (
	"context"
	"flag"
	pb "github.com/Gophberg/TransactionAPI/internal/app/TransactionAPI/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"time"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

var config Config

func Start() error {
	go grpcClient()
	config.NewConfig()
	t := Transaction{}
	http.HandleFunc("/getTransactionStatusById", t.getTransactionStatusById)
	http.HandleFunc("/getAllTransactionsByUserId", t.getAllTransactionsByUserId)
	http.HandleFunc("/getAllTransactionsByUserEmail", t.getAllTransactionsByUserEmail)
	http.HandleFunc("/createTransaction", t.createTransaction)
	return http.ListenAndServe(":9000", nil)
}

func grpcClient() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTransactionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	amount := &pb.TransactionRequest_Amount{
		Units: 11,
		Nanos: 22,
	}
	r, err := c.Transaction(ctx, &pb.TransactionRequest{
		ID:        1,
		UserID:    2,
		UserEmail: "joe@mail.edu",
		Currency:  "USD",
		Amount:    amount,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Received response from External Pay System: %s", r.GetMessage())
}
