# balance-api

Implementation of a microservice for working with user's balance. 

## Usage

Microservice accepts requests and sends responses in JSON. Examples of using the API are described below:
To start the service (Faced some problems with docker(Connection Refused)):
1. Cloning the repository:
```sh
git lone https://github.com/ksartamonov/balance-api
```
2. Configure Postgres database in file pkg/config/config.go 
3. Run SQL-scripts to create tables(for testing also) build/database/sql/startup.sql
4. Launch the program:
```sh
go run main.go
```

### Accrual of money to the balance

With this __POST__-request, it is possible to add money to the balance:
```http request
http://localhost:8080/billing-api/ProfitMoney
```
Body example:
```JSON
    {
        "id" : 1,
        "money" : 100
    }
```

### Reserving money
To reserve user's money for a transaction you can use a __POST__-request:
```http request
http://localhost:8080/billing-api/ReserveMoney
```

In case there is not enough money on the user's balance, the result will be:
```
Error! Not enough money!
```

In case of successful money reservation:
```
Success! Money is reserved!
```

Body example:
```JSON
    {
        "id" : 1,
        "money" : 50
    }
```

### Money Transfer
Money transfer can be made via the __POST__-request:
```http request
http://localhost:8080/billing-api/TransferMoney
```

Body example:
```JSON
    {
        "SenderId" : 1,
        "ReceiverId" : 2,
        "Money" : 30
    }
```

In case there is not enough money on the user's balance, the result will be:
```
Error! Not enough money!
```

In case of successful money reservation:
```
Success! Money transferred!
```

### Getting user's balance
Using this __POST__-request we can get user's balance:
```http request
http://localhost:8080/billing-api/CheckBalance
```

Body example:
```JSON
    {
        "id" : 1
    }
```
Reply will look like:
```JSON
{
    "balance": 1000,
    "id": 1
}
```
### Getting user's operations report
Via this __POST__-request we can get a CSV-report containing all the operations of a specific user:
```http request
http://localhost:8080/billing-api/GetUserReport
```

Body example:
```JSON
    {
        "Id" : 2,
        "Month" : 11,
        "Year" : 2022
    }
```
