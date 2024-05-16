# API Documentation: File Upload

This endpoint used to upload media such as images.

## Table of Contents

- [Controllers](#controllers)
  - [Images](#images)
    - [Get All Images](#get-all-images)
    - [Get Image](#get-image)
    - [Upload Image](#upload-image)
    - [Delete Image](#delete-image)

## Controllers

### Images

- ### Get All Images
  - Method: `GET`
  - URI: `/api/file/images`

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success get all images",
    "data": [
      {
        "url_id": "aickgrctlibriheojkudiiocizbcrmnublbnvgsx-car.png",
        "name": "car"
      }
    ]
  }
  ```

  Possibility Fail: -

- ### Get Image
  - Method: `GET`
  - URI: `/api/file/image/:imageID`

  Response: `Direct Image media`

  Possibility Fail: `image not found (404)`

- ### Upload Image
  - Method: `POST`
  - URI: `/api/file/image`

  Request Body (Form Data):
  ```json
  "file": yourImage (max=5mb)
  ```

  Response:
  ```json
  {
    "status": true,
    "status_code": 201,
    "message": "Success uploading an image",
    "data": {
      "url_id": "aickgrctlibriheojkudiiocizbcrmnublbnvgsx-car.png" // string
    }
  }
  ```
  Possibility Fail: `no image uploaded (400)`, `Image size too large (400)`

- ### Delete Image
  - Method: `DELETE`
  - URI: `/api/file/image/:imageID`

  Response:
  ```json
  {
    "status": true,
    "status_code": 200,
    "message": "Success deleting an image",
    "data": null
  }
  ```
  Possibility Fail: `image not found (404)`