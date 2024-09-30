# Setup the project
### install libs:
```console
go mod tidy
```
### run(windows):
```console
go build ./cmd/main.go
.\main.exe 
```
### run(linux):
Use MakeFile
### Test:
```console
go test .\app\... -cover -v
```
---
# Product API Usage

## Endpoint: 
### POST `/products`

### Description:
This endpoint allows you to create a new product

### Request Body Structure:
```json
{
  "SKU": "string"
}
```

- The SKU field is required and must be a string.

### Responses:

#### 1. **Success:**
- **Status Code:** 200
- **Response Body:** 
- A single integer representing the new ID.
    ```json
    1
    ```
    
#### 2. **Invalid Body Structure:**
- **Status Code:** 400
- **Response Body:** 
    ```json
    {
      "code": 400,
      "message": "invalid body structure"
    }
    ```

#### 3. **Internal Server Error:**
- **Status Code:** 500
- **Response Body:** 
    ```json
    {
      "code": 500,
      "message": "unable to save product"
    }
    ```
---
### GET `/product_registration/:id`

### Description:
This endpoint retrieves the product registration details based on the provided `id`.

### Path Parameter:
- `id`: Must be an integer representing the product registration ID.

### Responses:

#### 1. **Invalid ID (non-integer):**
- **Status Code:** 400
- **Response Body:** 
    ```json
    {
      "code": 400,
      "message": "int id expected in path"
    }
    ```

#### 2. **Product Registration Not Found:**
- **Status Code:** 404
- **Response Body:** 
    ```json
    {
      "code": 404,
      "message": "not found"
    }
    ```

#### 3. **Success:**
- **Status Code:** 200
- **Response Body Structure:**

```json
{
  "id": 123,
  "purchase_date": "2023-01-01T00:00:00Z",
  "expiry_at": "2025-01-01T00:00:00Z",
  "product": {
    "SKU": "ABC123"
  },
  "serial_code": "XYZ987",
  "additional_product_registrations": [
    {
      "id": 456,
      "purchase_date": "2023-02-01T00:00:00Z",
      "expiry_at": "2025-02-01T00:00:00Z",
      "product": {
        "SKU": "DEF456"
      },
      "serial_code": "LMN654"
    }
  ]
}
```
- The response contains detailed product registration information, including associated additional registrations(bundle).

---
### POST `/profiles`

### Description:
This endpoint allows you to create a new profile by sending a POST request with the specified body structure.

### Request Body Structure:
```json
{
  "email": "string",
  "firstname": "string",
  "lastname": "string"
}
```

- All fields are required and must be strings.

### Responses:

#### 1. **Invalid Body Structure:**
- **Status Code:** 400
- **Response Body:** 
    ```json
    {
      "code": 400,
      "message": "invalid body structure"
    }
    ```

#### 2. **Error on Save:**
- **Status Code:** 500
- **Response Body:** 
    ```json
    {
      "code": 500,
      "message": "unable to save profile"
    }
    ```

#### 3. **Success:**
- **Status Code:** 200
- **Response Body:** 
    - A single integer representing the new ID.
  ```json
  1
  ```
 ---
### GET `/profiles`

### Description:
This endpoint retrieves all profiles along with their respective product registrations. Each product registration includes associated bundles (additional product registrations).

### Response Structure:

#### Profile Structure:
```json
[{
  "id": 1,
  "email": "user@example.com",
  "firstname": "John",
  "lastname": "Doe",
  "product_registrations": [
    {
      "id": 123,
      "purchase_date": "2023-01-01T00:00:00Z",
      "expiry_at": "2025-01-01T00:00:00Z",
      "product": {
        "SKU": "ABC123"
      },
      "serial_code": "XYZ987",
      "additional_product_registrations": [
        {
          "id": 456,
          "purchase_date": "2023-02-01T00:00:00Z",
          "expiry_at": "2025-02-01T00:00:00Z",
          "product": {
            "SKU": "DEF456"
          },
          "serial_code": "LMN654"
        }
      ]
    }
  ]
}
]
```

---
### POST `/profiles/:profile/product_registrations`

### Description:
This endpoint allows you to add product registrations to a specific profile. The request body should contain details about the product registration and any associated bundles.

### Path Parameter:
- `profile`: The ID of the profile to which the product registrations will be added.

### Request Body Structure:
```json
{
  "purchase_date": "2023-01-01T00:00:00Z",
  "expiry_at": "2025-01-01T00:00:00Z",
  "product": {
    "SKU": "ABC123"
  },
  "serial_code": "XYZ987",
  "additional_product_registrations": [
    {
      "purchase_date": "2023-02-01T00:00:00Z",
      "expiry_at": "2025-02-01T00:00:00Z",
      "product": {
        "SKU": "DEF456"
      },
      "serial_code": "LMN654"
    }
  ]
}
```
#### 1. **Invalid Body Structure:**
- **Status Code:** 400
- **Response Body:** 
    ```json
    {
      "code": 400,
      "message": "invalid body structure"
    }
    ```

#### 2. **Error on Save:**
- **Status Code:** 500
- **Response Body:** 
    ```json
    {
      "code": 500,
      "message": "unable to save"
    }
    ```

#### 3. **Success:**
- **Status Code:** 200
- **Response Body:** 
    - A single integer representing the new ID for the root object.
  ```json
  1
  ```

---
### GET `/profiles/:profile/product_registrations`

### Description:
This endpoint retrieves all product registrations associated with a specific profile identified by `profile`.

### Path Parameter:
- `profile`: The ID of the profile for which product registrations are to be retrieved.

### Response Structure:

#### Product Registration Structure:
```json
{
  "id": 123,
  "purchase_date": "2023-01-01T00:00:00Z",
  "expiry_at": "2025-01-01T00:00:00Z",
  "product": {
    "SKU": "ABC123"
  },
  "serial_code": "XYZ987",
  "additional_product_registrations": [
    {
      "id": 456,
      "purchase_date": "2023-02-01T00:00:00Z",
      "expiry_at": "2025-02-01T00:00:00Z",
      "product": {
        "SKU": "DEF456"
      },
      "serial_code": "LMN654"
    }
  ]
}
```