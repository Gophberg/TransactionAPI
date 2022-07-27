## TransactionAPI
---
---
Execute:
- `docker-compose up` to run postgres db
- `make gen` to generate gRPC code
- `make eps` to run mock of external pay system (EPS)
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
If you try to create transaction with UserEmail "joe@mail.edu". 
EPS will return "Fail" with reason "I hate him" (:smile:) 
Or if you'll send zero or negative amount value. 
When EPS will return "Fail" with appropriate reason.  
After requesting createTransaction, 
API will create new record to database with status "New" and sending data to EPS.
Further EPS will return status and API do create new database record with appropriate reason. 
