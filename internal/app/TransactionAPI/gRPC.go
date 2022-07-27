package TransactionAPI

import (
	"context"
	"flag"
	"github.com/Gophberg/TransactionAPI/internal/app/TransactionAPI/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func gRPCTransactionRequest(ctx context.Context, t Transaction) {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("[gRPC] did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewTransactionClient(conn)

	amount, _ := t.Amount.Float64()
	r, err := c.Transaction(ctx, &pb.TransactionRequest{
		ID:        t.Id,
		UserID:    t.UserID,
		UserEmail: t.UserEmail,
		Currency:  t.Currency,
		Amount:    amount,
	})
	if err != nil {
		log.Fatalf("[gRPC] could not transact: %v", err)
	}
	log.Printf("[gRPC] Received response from External Pay System: %t", r.GetStatus())
}
