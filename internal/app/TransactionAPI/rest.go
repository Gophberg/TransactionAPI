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
	log.Printf("Requested create transaction: %s\n", r.URL.Path)

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
		return
	}

	id, err := t.CreateTransaction(t)
	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		return
	}
	log.Printf("%v bytes written to ResponseWriter", write)

	// Go gRPC transaction request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	gRPCTransactionRequest(ctx, t)

}

func (t Transaction) getTransactionStatusById(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requested transaction status by id: %s\n", r.URL.Path)

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
		return
	}

	status, err := t.GetTransactionStatusById(t)
	if err != nil {
		log.Println(err)
	}

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	write, err := w.Write(js)
	if err != nil {
		return
	}
	log.Printf("%v bytes written to ResponseWriter", write)
}

func (t Transaction) getAllTransactionsByUserId(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requested all transactions status by UserId: %s\n", r.URL.Path)

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
		return
	}

	transactions, err := t.GetAllTransactionsByUserId(t)
	if err != nil {
		log.Println(err)
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
	log.Printf("%v bytes written to ResponseWriter", write)
}

func (t Transaction) getAllTransactionsByUserEmail(w http.ResponseWriter, r *http.Request) {
	log.Printf("Requested all transactions status by UserEmail: %s\n", r.URL.Path)

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
		return
	}

	transactions, err := t.GetAllTransactionsByUserEmail(t)
	if err != nil {
		log.Println(err)
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
	log.Printf("%v bytes written to ResponseWriter", write)
}
