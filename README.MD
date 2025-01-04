# GoLang demo app

## Description
- This application uses SQLite for as a database
- It has [welcome page at http://localhost:8080](http://localhost:8080)
- All APIs are at `/api/v1/`
- For registration use [http://localhost:8080/api/v1/registration/signup](http://localhost:8080/api/v1/registration/signup)
- For login use [http://localhost:8080/api/v1/registration/login](http://localhost:8080/api/v1/registration/login)

## How to call rest-api

### **Authentication Process**
To access the authenticated routes:
1. **Sign Up**: Use `POST /api/v1/registration/signup` to create a new user.
2. **Log In**: Use `POST /api/v1/registration/login` with your username and password to receive a `token`.
3. **Add Token to Headers**: Include the `Authorization` header with the format `Bearer <token>` in your requests to authenticated routes.

### Registration
- **Endpoint**: `POST /api/v1/registration/signup`
- **Description**: Registers a new user.
- **Request Body**:
```json
  {
    "username": "example_user",
    "password": "example_password"
  }
```
- **Response**:
```json
  {
    "message": "User created successfully",
    "user_id": 123
  }
```

### User Login
- **Endpoint**: `POST /api/v1/registration/login`
- **Description**: Logs in an existing user and returns a token.
- **Request Body**:
```json
{
  "username": "example_user",
  "password": "example_password"
}
```
- **Response**:
```json
{
  "token": "eyJh... (JWT token)"
}
```

### **1. List All Books**
- **Endpoint**: `GET /api/v1/books`
- **Description**: Returns a list of all books in the database.
- **Headers**:
```
  Authorization: Bearer <token>
```
- **Response**:
```json
[
  {
    "id": 1,
    "name": "Book Name",
    "author": "Author Name",
    "description": "Description of the book",
    "borrowed": false,
    "last_borrowed_at": null,
    "last_returned_at": null,
    "current_borrower_id": null
  }
]
```

### **2. Get Book Details**
- **Endpoint**: `GET /api/v1/books/:id`
- **Description**: Fetches details of a specific book using its `id`.
- **Parameters**:
    - `id` (path): The ID of the book.
- **Headers**:
```
  Authorization: Bearer <token>
```
- **Response**:
```json
{
  "id": 1,
  "name": "Book Name",
  "author": "Author Name",
  "description": "Description of the book",
  "borrowed": false,
  "last_borrowed_at": null,
  "last_returned_at": null,
  "current_borrower_id": null
}
```

### **3. Borrow a Book**
- **Endpoint**: `POST /api/v1/books/:id/borrow`
- **Description**: Marks a book as borrowed by the currently authenticated user.
- **Parameters**:
    - `id` (path): The ID of the book.

- **Headers**:
```
  Authorization: Bearer <token>
```
- **Response**:
```json
{
  "message": "Book borrowed successfully",
  "borrowed_at": "2023-10-01T12:34:56Z"
}
```

### **4. Return a Book**
- **Endpoint**: `POST /api/v1/books/:id/return`
- **Description**: Marks a borrowed book as returned by the currently authenticated user.
- **Parameters**:
    - `id` (path): The ID of the book.

- **Headers**:
```
  Authorization: Bearer <token>
```

- **Response**:
```json
{
  "message": "Book returned successfully",
  "returned_at": "2023-10-02T15:12:34Z"
}
```