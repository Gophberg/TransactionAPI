# TransactionAPI
#### External DB:
```
User:
    ID      integer
    e-mail  text
```
#### Internal DB:
```
States:
    Save    boolean
    Update  boolean
```
#### Transactions:
```
TransactionID integer
UserID        integer
UserEmail     text
Amount        float
Currency      integer (need to redefinition or use type text) 
CreationDate  date
UpdateDate    date  
Status        integer (need to redefinition or use type text)
```
#### Transactions Statuses:
```
New
Success
Failure
Error
Canceled
```
### API functions:
```
Transaction:
    Create:
        {
            UserID,
            UserEmail,
            Amount,
            Currency
        }

    Status:
        Update
    Check by TransactionID
    SELECT * FROM Transactions WHERE UserID = "some id"
    SELECT * FROM Transactions WHERE UserID = "some id"
    Cancel transaction for an UserID
        Error if impossible (if status is success/failure, i.e. locked)
```
### Transaction Cycle:
```
User:
    Create:
        Status := New
        Create:
            {
                UserID,
                UserEmail,
                Amount,
                Currency
            }
API:
    Status: Lock
        Success/Failure: random
    To DB: state record
        Error: any one by random
```
### Server cycle:
```
1. db init
2. listen and serve
   1. /create
   2.
3. responce
```
