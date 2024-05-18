# API Documentation

This is documentation related to the request and response formats as well as the routes of this project.

## Table of Contents

- Base API Response
  - [On Success](#on-success)
  - [On Validation Error](#on-validation-error)
  - [On Internal Error](#on-internal-error)
- Filter by Parameter
  - [Sort by Date](#sort-by-date)
  - [Sort by Name](#sort-by-name)
- Pagination
  - [Pagination](#pagination)
- Endpoint
  - [Balance](./api/balance.md)
  - [Balance History](./api/balanceHistory.md)
  - [User](./api/user.md)
  - [Kas Submission](./api/kasSubmission.md)
  - [Financial Request](./api/financialRequest.md)
  - [Month](./api/month.md)
  - [File Upload](./api/fileUpload.md)

## Base API Response
- #### **On Success**
  ```json
  {
    "status": true,           // boolean
    "status_code": 200,       // int
    "message": "Success ...", // string
    "data": {}                // object | array
  }
  ```

- #### **On Validation Error**
  ```json
  {
    "status": false,                        // boolean
    "status_code": 400,                     // int
    "message": "Invalid request body",      // string
    "errors": [                             // array | null
      {
        "field": "year",                    // string
        "message": "This field is required"     // string
      }
    ]
  }
  ```

- #### **On Internal Error**
  ```json
  {
    "status": false,         // boolean
    "status_code": 500,      // int
    "message": "Failed ...", // string
    "errors": null           // null
  }
  ```

## Filter by Parameter
### **Sort by Date**
Acceptable param variable is `sort` and `order_by`:
- The **sort** variable can be `created_at` or `updated_at`
- The **order_by** variable can be `asc` or `desc`

Example:
- Newest Created: `?sort=created_at&order_by=desc`
- Oldest Created: `?sort=created_at&order_by=asc`
- Newest Updated: `?sort=updated_at&order_by=desc`
- Oldest Updated: `?sort=updated_at&order_by=asc`

Note: this filter is optional. If fetching without this filter, it will be shows by Oldest Created.

### **Sort by Name**
Acceptable param variable is `sort` and `order_by`:
- The **sort** variable can be only `name`
  - The field that will be sorted already fixed, for example: user will be sort by Name field.
- The **order_by** variable can be `asc` or `desc`

Example:
- A-Z: `?sort=name&order_by=asc`
- Z-A: `?sort=name&order_by=desc`

Note: this filter is optional. If fetching without this filter, it will be shows by Oldest Created.

## Pagination
Acceptable param variable for pagination is `limit` and `page`:
- The **limit** variable is total of maximum data that will be fetched
- The **page** variable is to determines the page of data that will be fetched

If **limit** and **page** used for the fetch, there will be extra property as base API response like this:
```json
{
  "status": true,           // boolean
  "status_code": 200,       // int
  "message": "Success ...", // string
  "data": {},               // object | array
  "pagination": {
    "page": 1,              // int
    "limit": 5,             // int
    "total_items": 10,      // int
    "total_pages": 2        // int
  }
}
```
The `total_items` and `total_pages` will be automatically generated.

Note: this filter is optional. If fetching without this pagination, it will be showing entire data.