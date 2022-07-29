package TransactionAPI

import (
	"context"
	"encoding/json"
	"log"
	"mime"
	"net/http"
	"time"
)

func (t Transaction) createTransaction(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REST] Requested create transaction: %s\n", r.URL.Path)

	t.decodeData(w, r)
	log.Println("[REST] Decoded data to t type:", t)

	// Create new record of transaction in database with status "New"
	t.Status = "New"
	id, err := t.createRecord(t)
	if err != nil {
		log.Println("[REST]", err)
	}
	log.Println("[REST] New transaction created with id:", id)

	// Sending response
	js, err := json.Marshal("Ok")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		return
	}
	log.Printf("[REST] %v bytes written to ResponseWriter", write)

	// Go gRPC transaction request
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15) // Modify time there and in mock/main.go/doTransaction if you need more time to testing of canceling transaction
		received := <-t.CancelTrChan
		if received == t.Id {
			ctx.Done()
		}
		defer cancel()
		gRPCTransactionRequest(ctx, t)
	}()
}

func (t Transaction) getTransactionStatusById(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REST] Requested transaction status by id: %s\n", r.URL.Path)

	t.decodeData(w, r)

	status, err := t.readRecord(t)
	if err != nil {
		log.Println("[REST]", err)
	}

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		return
	}
	log.Printf("[REST] %v bytes written to ResponseWriter", write)
}

func (t Transaction) cancelTransaction(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REST] Requested cancel transaction: %s\n", r.URL.Path)

	t.decodeData(w, r)

	// Sending cancel signal to a cancelChannel
	t.CancelTrChan <- t.Id

	// Do response
	js, err := json.Marshal("Canceling...")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		log.Println("[REST] write error:", err)
	}
	log.Printf("[REST] %v bytes written to ResponseWriter", write)
}

func (t Transaction) getAllTransactions(w http.ResponseWriter, r *http.Request) {
	log.Printf("[REST] Requested all transactions: %s\n", r.URL.Path)

	t.decodeData(w, r)

	transactions, err := t.readRecords(t, r.URL.Path)
	if err != nil {
		log.Println("[REST]", err)
	}

	js, err := json.Marshal(transactions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		return
	}
	log.Printf("[REST] %v bytes written to ResponseWriter", write)
}

func (t *Transaction) decodeData(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	log.Printf("[REST] Data to decode: %v\n", &t)
}
