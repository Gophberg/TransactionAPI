# TransactionAPI
#### User struct:
```
UserID    string
UserEmail string
Amount    int
Currency  string
```
#### Transaction struct:
```
TransactionID integer
UserID        integer
UserEmail     text
Amount        decimal.Decimal (or int+int)
Currency      text 
CreationDate  date
UpdateDate    date  
Status        text
```
#### Status struct:
```
New
Success
Failure
Error
Canceled
```
### Cycle:
```
UserByREST:
    /user
        api.TransactionCreate(*user)
        api.TransactionStatus(*user)
        api.TransactionCancel(*user)

API:
    transactionCreate(*transaction) status {
        db.Status = new
        Status = lock()
        ExtPaySys.Auth(jwt)
        ExtPaySys.TransactionRequest(*user) {
            if ExtPaySys.Status = success || failure {
                Status := "locked"
                db.Update(status)
                return "status"
            }
            ulse {
                db.Update(status)
                return "status"
            }
        }
    }
    TransactionStatus(*transaction) status {
        return ExtPaySys.CheckStatus(*user.Id)
    }
    db.SelectById(*user.Id) []transactions
    db.SelectByEmail(*user.Email) []transactions
    TransactionCancel(*user)
        if Status = locked {
            return "impossible"
        }
        ExtPaySys.TransactionCancel(*transaction.Id) 
            Error if impossible (if status is success/failure, i.e. locked)

ExtPaySys:
    Auth(jwt)
    TransactionRequest(*transactionExtPaySys) {
        defer time.Sleep(10 * time.Second)
        if TransactionCancel {
            // canceling()
            return "canceled"
        }
        return "success" || "failure"
    }
    CheckStatus(*transactionExtPaySys) {
        return "status"
    }
```
