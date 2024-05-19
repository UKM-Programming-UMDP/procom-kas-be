# API Documentation: Balance History

This endpoint used to get all of history when there is changes on balance.

## Table of Contents

- [Controllers](#controllers)
  - [Get History](#get-history)

## Controllers

- ### Get History
  - Method: `GET`
  - URI: `/api/balance/history`
  - Filter: [sort by date](../README.md/#filter-by-parameter), [pagination](../README.md/#pagination)

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success getting balance history",
    "data": [
      {
        "amount": 10000,
        "activity": "Add",
        "note": "[kas pay]",
        "user": {
          "npm": "2125240015",
          "name": "Fanes Pratama"
        },
        "created_at": "2024-05-12T10:09:37.88331+08:00"
      }
    ]
  }
  ```

  Possibility Fail: `Balance not initialized yet (500)`
