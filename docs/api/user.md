# API Documentation: User

This endpoint used to do CRUD on user.

## Table of Contents

- [Controllers](#controllers)
  - [Get All User](#get-all-user)
  - [Get User By NPM](#get-user-by-npm)
  - [Create User](#create-user)
  - [Update User](#update-user)
  - [Delete User](#delete-user)

## Controllers

- ### Get All User
  - Method: `GET`
  - URI: `/api/users`
  - Filter: [sort by date](../README.md/#filter-by-parameter), [sort by name](../README.md/#filter-by-parameter), [pagination](../README.md/#pagination)

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success getting users",
    "data": [                        // array
      {
        "npm": "2125240015",         // string
        "name": "Fanes Pratama",     // string
        "email": "fanes@gmail.com",  // string
        "kas_payed": 20000,          // int
        "month_start_pay": {
          "id": 3                    // int (PK Month ID)
        }
      }
    ]
  }
  ```

  Possibility Fail: -

- ### Get User By NPM
  - Method: `GET`
  - URI: `/api/users/details`

  Parameter Variable:
  ```js
  npm = string (len=10)
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success getting a user",
    "data": {
      "npm": "2125240015",         // string
      "name": "Fanes Pratama",     // string
      "email": "fanes@gmail.com",  // string
      "kas_payed": 20000,          // int
      "month_start_pay": {
        "id": 3                    // int (PK Month ID)
      }
    }
  }
  ```

  Possibility Fail: `user not found (404)`

- ### Create User
  - Method: `POST`
  - URI: `/api/users`

  Request Body:
  ```json
  {
    "npm": "2125240015",         // string (len=10)
    "name": "Fanes Pratama",     // string (min=3, max=255)
    "email": "fanes@gmail.com",  // string (email)
    "kas_payed": 0,              // int (min=0)
    "month_start_pay": {
      "id": 2                    // int (PK Month ID)
    }
  }
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 201,
    "message": "Success creating a user",
    "data": {
      "npm": "2125240015",         // string
      "name": "Fanes Pratama",     // string
      "email": "fanes@gmail.com",  // string
      "kas_payed": 0,              // int
      "month_start_pay": {
        "id": 2                    // int (PK Month ID)
      }
    }
  }
  ```

  Possibility Fail: `user already exist (409)`, `email already exists (409)`, `month not found (404)`

- ### Update User Visibility
  - Method: `PUT`
  - URI: `/api/users`

  Parameter Variable:
  ```js
  npm = string (len=10)
  ```

  Request Body:
  ```json
  {
    "name": "Fanes Pratama",     // string (min=3, max=255)
    "email": "fanes@gmail.com",  // string (email)
    "kas_payed": 50000,          // int (min=0)
    "month_start_pay": {
      "id": 2                    // string (PK Month ID)
    }
  }
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success updating a user",
    "data": {
      "npm": "2125240015",         // string
      "name": "Fanes Pratama",     // string
      "email": "fanes@gmail.com",  // string
      "kas_payed": 50000,          // int
      "month_start_pay": {
        "id": 3                    // int (PK Month ID)
      }
    }
  }
  ```
  Possibility Fail: `user not found (404)`, `email already exists (409)`, `month not found (404)`

- ### Delete User
  - Method: `DELETE`
  - URI: `/api/users`

  Parameter Variable:
  ```js
  npm = string (len=10)
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success deleting a user",
    "data": null
  }
  ```
  Possibility Fail: `user not found (404)`