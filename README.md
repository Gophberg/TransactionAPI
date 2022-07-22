## TransactionAPI
---
---
Execute:
- `docker-compose up` to run postgres db
- `make gen` to generate gRPC code
- `make eps` to run mock of external pay system
- `make` to run main server
---
To get status of transaction by id, do request  
`curl -X POST http://localhost:9000/getTransactionStatusById -H "Content-Type: application/json" --data '{"id": N}'`
Where `N` is a number of transaction to check.  
  
To get all transactions by UserId, do request  
`curl -X POST http://localhost:9000/getAllTransactionsByUserId -H "Content-Type: application/json" --data '{"userid": N}'`  
Where `N` is a number of transactions to check.  
  
To get all transactions by UserEmail, do request  
`curl -X POST http://localhost:9000/getAllTransactionsByUserEmail -H "Content-Type: application/json" --data '{"useremail": "joe@mail.edu"}'`  
Note: In this point my db contains the two records. `joe@mail.edu` and `jane@mail.edu`
  
To create transaction, do request  
`curl -X POST http://localhost:9000/createTransaction -H "Content-Type: application/json" --data '{"userid": 100, "useremail": "maks@mail.edu", "amount": 333.11, "currency": "usd", "datecreated": "2022-07-02T13:57:00Z", "dateupdated": "2022-07-02T13:57:02Z"}'`  
You can to post your credentials in this request
