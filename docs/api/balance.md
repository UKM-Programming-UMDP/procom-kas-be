# API Documentation: Balance

This endpoint used to get or update balance by admin.

## Table of Contents

- [Controllers](#controllers)
  - [Get Balance](#get-balance)
  - [Update Balance](#update-balance)

## Controllers

- ### Get Balance
  - Method: `GET`
  - URI: `/api/balance`

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success getting a balance",
    "data": {
      "balance": 20000,  // int
      "updated_at": "2024-05-12T16:19:45.042201+08:00"  // timestamp
    }
  }
  ```

  Possibility Fail: `Balance not initialized yet (500)`

- ### Update Balance
  - Method: `PUT`
  - URI: `/api/balance`

  Activity Format = `1: Add`, `2: Substract`

  Request Body:
  ```json
  {
    "user": {
      "npm": "2125240015"   // string (PK User NPM)
    },
    "amount": 5000,         // int (min=1)
    "activity": 2,          // int (len=1, only:1|2)
    "note": "Bayar parkir"  // string
  }
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success updating balance",
    "data": {
      "balance": 25000,  // int
      "updated_at": "2024-05-12T16:19:45.042201+08:00"  // timestamp
    }
  }
  ```

  Possibility Fail: `user not found (404)`, `invalid activity (400)`

