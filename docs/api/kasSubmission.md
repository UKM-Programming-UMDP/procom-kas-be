# API Documentation: Kas Submission

This endpoint used to do CRUD on submitted kas payment.

## Table of Contents

- [Controllers](#controllers)
  - [Get All Kas Submission](#get-all-kas-submission)
  - [Get Kas Submission by ID](#get-kas-submission-by-id)
  - [Create Kas Submission](#create-kas-submission)
  - [Update Kas Submission Status](#update-kas-submission-status)

## Controllers

- ### Get All Kas Submission
  - Method: `GET`
  - URI: `/api/kas`
  - Filter: [sort by date](../README.md/#filter-by-parameter), [pagination](../README.md/#pagination)

  Response:
  ```json
  {
      "status": true,
      "status_code": 200,
      "message": "Success getting kas submissions",
      "data": [                            // array
        {
          "submission_id": "qwert",        // string
          "user": {
            "npm": "2125240015",           // string
            "name": "Fanes Pratama",       // string
            "email": "fanes@mail.com",     // string (email)
            "kas_payed": 20000,            // int
            "month_start_pay": {
              "id": 53                     // int (PK Month ID)
            }
          },
          "payed_amount": 10000,           // int
          "status": "Pending",             // string
          "note": "oke",                   // string
          "evidence": "urlRandomId.a.png", // string (URL ID)
          "submitted_at": "2024-05-09T21:58:52.235233+08:00", // timestamp
          "updated_at": "2024-05-09T21:58:52.235233+08:00"    // timestamp
        }
      ]
  }
  ```

  Possibility Fail: -

- ### Get Kas Submission by ID
  - Method: `GET`
  - URI: `/api/kas/details`

  Parameter Variable:
  ```js
  submission_id = string (len=5)
  ```

  Response:
  ```json
  {
      "status": true,
      "status_code": 200,
      "message": "Success getting a kas submission",
      "data": {
        "submission_id": "qwert",        // string
        "user": {
          "npm": "2125240015",           // string
          "name": "Fanes Pratama",       // string
          "email": "fanes@mail.com",     // string
          "kas_payed": 20000,            // int
          "month_start_pay": {
            "id": 53                     // int (PK Month ID)
          }
        },
        "payed_amount": 10000,           // int
        "status": "Pending",             // string
        "note": "oke",                   // string
        "evidence": "urlRandomId.a.png", // string (URL ID)
        "submitted_at": "2024-05-09T21:58:52.235233+08:00", // timestamp
        "updated_at": "2024-05-09T21:58:52.235233+08:00"    // timestamp
      }
  }
  ```

  Possibility Fail: `kas submission not found (404)`

- ### Create Kas Submission
  - Method: `POST`
  - URI: `/api/kas`

  Request Body:
  ```json
  {
    "user": {
      "npm": "2125240015"               // string (PK User NPM, len=10)
    },
    "payed_amount": 10000,              // int (min=1)
    "note": "test",                     // string
    "evidence": "urlRandomId-a.png"     // string (URL ID)
  }
  ```

  Response:
  ```json
  {
      "status": true,
      "status_code": 200,
      "message": "Success creating a kas submission",
      "data": {
        "submission_id": "qwert",        // string
        "user": {
          "npm": "2125240015",           // string
          "name": "Fanes Pratama",       // string
          "email": "fanes@mail.com",     // string
          "kas_payed": 20000,            // int
          "month_start_pay": {
            "id": 53                     // int (PK Month ID)
          }
        },
        "payed_amount": 10000,           // int
        "status": "Pending",             // string
        "note": "oke",                   // string
        "evidence": "urlRandomId.a.png", // string (URL ID)
        "submitted_at": "2024-05-09T21:58:52.235233+08:00", // timestamp
        "updated_at": "2024-05-09T21:58:52.235233+08:00"    // timestamp
      }
  }
  ```

  Possibility Fail: `user not found (404)`

- ### Update Kas Submission Status
  - Method: `PUT`
  - URI: `/api/kas`

  Status Format = `1: Approved`, `2: Rejected`, `3: Pending`

  Parameter Variable:
  ```js
  submission_id = string (len=5)
  ```

  Request Body:
  ```json
  {
    "status": 1 // int (len=1, only:1|2)
  }
  ```

  Response:
  ```json
  {
      "status": true,
      "status_code": 200,
      "message": "Success updating kas submission status",
      "data": {
        "submission_id": "qwert",        // string
        "user": {
          "npm": "2125240015",           // string
          "name": "Fanes Pratama",       // string
          "email": "fanes@mail.com",     // string
          "kas_payed": 20000,            // int
          "month_start_pay": {
            "id": 53                     // int (PK Month ID)
          }
        },
        "payed_amount": 10000,           // int
        "status": "Approved",            // string
        "note": "oke",                   // string
        "evidence": "urlRandomId.a.png", // string (URL ID)
        "submitted_at": "2024-05-09T21:58:52.235233+08:00", // timestamp
        "updated_at": "2024-05-09T21:58:52.235233+08:00"    // timestamp
      }
  }
  ```
  Possibility Fail: `kas submission not found (404)`, `kas submission already processed (400)`, `cannot change status to pending (400)`, `invalid status (400)`, `user not found (404)`
