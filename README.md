### TransactionAPI 

---
To run postgres db execute `docker-compose up`  
And run `make`
---
To get status of transaction by id, do request `localhost:9000/getStatus?id=N`  
Where `N` is a number of transaction to check.  
To get all transactions by UserId, do request `localhost:9000/getAllTransactionsByUserId?id=N`  
Where `N` is a number of transactions to check.  
To get all transactions by UserEmail, do request `localhost:9000/getAllTransactionsByUserEmail?email="joe@mail.edu"`  
Note: In this point my db contains the two records. `joe@mail.edu` and `jane@mail.edu`
