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
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("[gRPC] connection error: %v", err)
		}
	}(conn)
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
	log.Printf("[gRPC] Response from EPS. Status '%t', reason '%s'", r.GetStatus(), r.GetReason())

	// Create DB record with status what returned by EPS
	t.Status = convertStatus(r.GetStatus())
	id, err := t.CreateTransaction(t)
	if err != nil {
		log.Println("[gRPC]", err)
	}
	log.Println("[gRPC] New transaction created with id:", id)

}

func convertStatus(s bool) string {
	if s == true {
		return "Success"
	}
	return "Fail"
}
