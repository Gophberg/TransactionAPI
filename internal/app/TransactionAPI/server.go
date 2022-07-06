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
	go grpcserver()
	config.NewConfig()
	t := Transaction{}
	http.HandleFunc("/getTransactionStatusById", t.getTransactionStatusById)
	http.HandleFunc("/getAllTransactionsByUserId", t.getAllTransactionsByUserId)
	http.HandleFunc("/getAllTransactionsByUserEmail", t.getAllTransactionsByUserEmail)
	http.HandleFunc("/createTransaction", t.createTransaction)
	return http.ListenAndServe(":9000", nil)
}

func grpcserver() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTransactionClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Transaction(ctx, &pb.TransactionRequest{UserEmail: "hello@joe.edu"})
	//r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())

}
