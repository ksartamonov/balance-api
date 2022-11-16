# balance-api

Implementation of a microservice for working with user's balance. 

## Usage

Microservice accepts requests and sends responses in JSON. Examples of using the API are described below:

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

In case of successful money transfer:
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

