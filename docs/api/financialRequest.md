# API Documentation: Financial Request

This endpoint used to do CRUD on requested financial.

## Table of Contents

- [Controllers](#controllers)
  - [Get All Financial Request](#get-all-financial-request)
  - [Get Financial Request by ID](#get-financial-request-by-id)
  - [Create Financial Request](#create-financial-request)
  - [Update Financial Request Status](#update-financial-request-status)

## Controllers

- ### Get All Financial Request
  - Method: `GET`
  - URI: `/api/financial-request`
  - Filter: [sort by date](../README.md/#filter-by-parameter), [pagination](../README.md/#pagination)

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success getting financial requests",
    "data": [
      {
        "request_id": "qwert",
        "amount": 2000,
        "note": "beli pena",
        "user": {
          "npm": "2125240015",
          "name": "Fanes Pratama",
          "email": "fanes@gmail.com"
        },
        "status": "Pending",
        "payment": {
          "type": "Transfer",
          "target_provider": "Bank - BCA",
          "target_name": "Fanes Pratama",
          "target_number": "1234567890",
          "evidence": "evidence.png"
        },
        "transfered_evidence": "",
        "created_at": "2024-05-12T16:50:36.085987+08:00",
        "updated_at": "2024-05-12T16:50:36.085987+08:00"
      }
    ]
  }
  ```

  Possibility Fail: -

- ### Get Financial Request by ID
  - Method: `GET`
  - URI: `/api/financial-request`

  Parameter Variable:
  ```js
  request_id = string (len=5)
  ```

  Response:
  ```json
  {
      "status": true,
      "status_code": 200,
      "message": "Success getting a financial request",
      "data": {
        "request_id": "qwert",
        "amount": 2000,
        "note": "beli pena",
        "user": {
          "npm": "2125240015",
          "name": "Fanes Pratama",
          "email": "fanes@gmail.com"
        },
        "status": "Pending",
        "payment": {
          "type": "Transfer",
          "target_provider": "Bank - BCA",
          "target_name": "Fanes Pratama",
          "target_number": "1234567890",
          "evidence": "evidence.png"
        },
        "transfered_evidence": "",
        "created_at": "2024-05-12T16:50:36.085987+08:00",
        "updated_at": "2024-05-12T16:50:36.085987+08:00"
      }
  }
  ```

  Possibility Fail: `financial request not found (404)`

- ### Create Financial Request
  - Method: `POST`
  - URI: `/api/financial-request`

  Request Body:
  ```json
  {
    "amount": 2000,
    "note": "beli pena",
    "user": {
      "npm": "2125240025"
    },
    "payment": {
      "type": "transfer",
      "target_provider": "Bank - BCA",
      "target_name": "Fanes Pratama",
      "target_number": "1234567890",
      "evidence": "evidence.png"
    }
  }
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 201,
    "message": "Success creating a financial request",
    "data": {
      "request_id": "qwert",
      "amount": 2000,
      "note": "beli pena",
      "user": {
        "npm": "2125240015",
        "name": "Fanes Pratama",
        "email": "fanes@gmail.com"
      },
      "status": "Pending",
      "payment": {
        "type": "Transfer",
        "target_provider": "Bank - BCA",
        "target_name": "Fanes Pratama",
        "target_number": "1234567890",
        "proof": "proof.png"
      },
      "transfered_proof": "",
      "created_at": "2024-05-12T16:50:36.085987+08:00",
      "updated_at": "2024-05-12T16:50:36.085987+08:00"
    }
  }
  ```

  Possibility Fail: `user not found (404)`

- ### Update Financial Request Status
  - Method: `PUT`
  - URI: `/api/financial-request`

  Status Format = `1: Approved`, `2: Rejected`, `3: Pending`

  Parameter Variable:
  ```js
  request_id = string (len=5)
  ```

  Request Body:
  ```json
  {
    "status": 1,  // int (len=1, only:1|2)
    "transfered_evidence": "transfered_evidence.png"  // string
  }
  ```

  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success updating financial request status",
    "data": {
      "request_id": "qwert",
      "amount": 2000,
      "note": "beli pena",
      "user": {
        "npm": "2125240015",
        "name": "Fanes Pratama",
        "email": "fanes@gmail.com"
      },
      "status": "Approved",
      "payment": {
        "type": "Transfer",
        "target_provider": "Bank - BCA",
        "target_name": "Fanes Pratama",
        "target_number": "1234567890",
        "evidence": "evidence.png"
      },
      "transfered_evidence": "transfered_evidence.png",
      "created_at": "2024-05-12T16:50:36.085987+08:00",
      "updated_at": "2024-05-12T16:50:36.085987+08:00"
    }
  }
  ```
  Possibility Fail: `financial request not found (404)`, `financial request already processed (400)`, `cannot change status to pending (400)`, `invalid status (400)`, `user not found (404)`
