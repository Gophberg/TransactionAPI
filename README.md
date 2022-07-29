## TransactionAPI

---

This is a test task for the company that closed that vacancy. 
But I still develop it for myself.
You can see all steps of developing by commits history.  
In the half of work I added the two `.todo` files with tasks.   
The `todo.todo` file contents the global tasks.  
And in the `workSchedulesReports.todo` file I add the tasks for today and tasks that appeared during the work. 

---

## How to use

*Execute:*
- `docker-compose up` to run postgres db
- `make gen` to generate gRPC code  
    Note: This step is not required. Proto files is currently generated. But you can add some changes for tests.
- `make eps` to run mock of external pay system (EPS)
- `make` to run main server
 
---

### Creating the transaction

*Do request in another terminal*  
`curl -X POST http://localhost:9000/createTransaction -H "Content-Type: application/json" --data '{"userid": 100, "useremail": "maks@mail.edu", "amount": 333.11, "currency": "usd"}'`  

You can to post your credentials in this request.  
If you try to create transaction with UserEmail `"joe@mail.edu"`.
EPS will return `"Fail"` with reason `"I hate him"` (:smile:)   
Or if you'll send zero or negative amount value.
When EPS will return `"Fail"` with appropriate reason.    
After requesting createTransaction,
API will create new record to database with status `"New"` and sending data to EPS.  
Further EPS will return status and API do create new database record with appropriate reason.

### Get status of transaction by ID

*Do request*  
`curl -X POST http://localhost:9000/getTransactionStatusById -H "Content-Type: application/json" --data '{"id": N}'`  

Where `N` is a number of transaction ID to check.  
  
### Get all transactions by UserId

*Do request*    
`curl -X POST http://localhost:9000/getAllTransactionsByUserId -H "Content-Type: application/json" --data '{"userid": N}'`    

Where `N` is a number of UserID to check.  

### Get all transactions by UserEmail
  
*Do request*  
`curl -X POST http://localhost:9000/getAllTransactionsByUserEmail -H "Content-Type: application/json" --data '{"useremail": "joe@mail.edu"}'`    

Note: In this point my db contains the two records `joe@mail.edu` and `jane@mail.edu` that added by `createtable.db` migrations file.
