# Banking Backend Service

This project is a simple banking backend service implemented in Go. It allows users to perform basic banking operations such as depositing and withdrawing money from accounts.

## Features

- Create and manage accounts
- Deposit money into accounts
- Withdraw money from accounts
- Track transactions for each account

## requirements
- golang 1.23

## Clone
```sh
   git clone git@github.com:lonmarsDev/banking-backend.git
   cd banking-backend
```

## how to run api server
on root project directory do
```sh
go run ./cmd/main.go
```

log will see like this
```sh
2024/08/31 01:14:32 INFO Starting server on :8383
```

## Testing api

Fist of all need to create account
```sh
curl --location --request POST 'http://localhost:8383/create-account' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "marlon pamisa"
}'
```

Deposit
```sh
curl --location --request POST 'http://localhost:8383/deposit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id" : 1,
    "amount" : 300
}'
```

Withdrawal
```sh
curl --location --request POST 'http://localhost:8383/withdraw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id" : 1,
    "amount" : 300
}'
```

View Balance
```sh
curl --location --request GET 'http://localhost:8383/view-balance?id=1'
```

Get SOA
```sh
curl --location --request GET 'http://localhost:8383/get-soa?id=1'
```

### Contact
For more information, please contact marlonpamisa@gmail.com.