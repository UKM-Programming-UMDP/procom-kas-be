# API Documentation: Month

This endpoint used to do CRUD on month that available to pay.

## Table of Contents

- [Controllers](#controllers)
  - [Get All Month](#get-all-month)
  - [Create Month](#create-month)
  - [Update Month Visibility](#update-month-visibility)
  - [Delete Month](#delete-month)

## Controllers

- ### Get All Month
  - Method: `GET`
  - URI: `/api/month`

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success getting a month",
    "data": [          // array
      {
        "id": 1,       // int
        "year": 2023,  // int
        "month": 4,    // int
        "show": false  // boolean
      }
    ]
  }
  ```

  Possibility Fail: -

- ### Create Month
  - Method: `POST`
  - URI: `/api/month`

  Request Body:
  ```json
  {
    "year": 2023,  // int (min=2000, max=9999)
    "month": 4     // int (min=1, max=12)
  }
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 201,
    "message": "Success creating a month",
    "data": {
      "id": 1,       // int
      "year": 2023,  // int
      "month": 4,    // int
      "show": false  // boolean
    }
  }
  ```

  Possibility Fail: `month already exist (209)`

- ### Update Month Visibility
  - Method: `PUT`
  - URI: `/api/month`

  Parameter Variable:
  ```json
  year = int (min=2000, max=9999)
  month = int (min=1, max=12)
  ```

  Request Body:
  ```json
  {
    "show": true  // boolean
  }
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success updating show month",
    "data": {
      "id": 1,       // int
      "year": 2023,  // int
      "month": 4,    // int
      "show": true   // boolean
    }
  }
  ```
  Possibility Fail: `month not found (404)`

- ### Delete Month
  - Method: `DELETE`
  - URI: `/api/month`

  Parameter Variable:
  ```json
  year = int (min=2000, max=9999)
  month = int (min=1, max=12)
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success deleting a month",
    "data": null
  }
  ```
  Possibility Fail: `month not found (404)`