package TransactionAPI

import (
	"encoding/json"
	"fmt"
	"log"
	"mime"
	"net/http"
	"strconv"
	"strings"
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
	//var rt RequestTask
	if err := dec.Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := t.CreateTransaction(t)
	if err != nil {
		log.Println(err)
	}
	//js, err := json.Marshal(ResponseId{Id: id})
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(js)

	b, err := fmt.Fprint(w, err)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%d bytes writed. With id: %d", b, id)
}

func (t Transaction) getTransactionStatusById(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?id=")
	id, err := strconv.Atoi(splitRoute[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Requested status of transaction id: %d\n", id)

	status, err := t.GetTransactionStatusById(id)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, status)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getAllTransactionsByUserId(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?id=")
	id, err := strconv.Atoi(splitRoute[1])
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Requested all transactions of UserId: %d\n", id)
	transactions, err := t.GetAllTransactionsByUserId(id)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, transactions)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}

func (t Transaction) getAllTransactionsByUserEmail(w http.ResponseWriter, r *http.Request) {
	splitRoute := strings.Split(r.URL.String(), "?email=")
	email := strings.Trim(splitRoute[1], "%22")
	fmt.Printf("Requested all transactions of UserEmail: %v\n", email)
	transactions, err := t.GetAllTransactionsByUserEmail(email)
	if err != nil {
		log.Println(err)
	}
	b, err := fmt.Fprint(w, transactions)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(b)
}
