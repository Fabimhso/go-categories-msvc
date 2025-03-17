# Go MSVC Category API

## About
This is a Go application developed for a college project made completely by me. It implements an HTTP server that exposes an API for managing categories. The API allows users to create new categories and list all existing categories.

## Features
- Create a new category
- List all created categories

## Technologies Used
- Go (Golang)
- Gin Gonic (web framework)
- JSON for data exchange
- In-Memory Repository (can be replaced with a database)

## Architecture
The application follows a modular structure based on **Clean Architecture** principles:

📂 **cmd/api/** – Contains the main API code, including server initialization and routes.  
📂 **internal/controllers/** – Controllers responsible for handling HTTP requests.  
📂 **internal/entities/** – Data model definitions.  
📂 **internal/repositories/** – Implementation of data storage (In-Memory Repository).  
📂 **internal/use-cases/** – Business logic for category management.  

## How to Run

1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Run the server:
   ```sh
   go run cmd/api/main.go
   ```
3. The server will be running at `http://localhost:8080`

## Endpoints

### Health Check
**GET** `/healthy`
- **Response (200 OK):**
  ```json
  {
    "status": "ok"
  }
  ```

### Create Category
**POST** `/categories`
- **Request:** JSON with the following fields:
  ```json
  {
    "name": "Category Name"
  }
  ```
- **Success Response (201 Created):**
  ```json
  {
    "categories": [
      {
        "id": 0,
        "name": "Docker",
        "created_at": "2025-03-17T02:28:10.9953065-03:00",
        "updated_at": "2025-03-17T02:28:10.9953065-03:00"
      }
    ],
    "success": true
  }
  ```

### List Categories
**GET** `/categories`
- **Success Response (200 OK):**
  ```json
  [
    {
      "id": 1,
      "name": "Category 1"
    },
    {
      "id": 2,
      "name": "Category 2"
    }
  ]
  ```

## Testing
You can test the API using **Postman**, **Insomnia**, or directly from the `api.http` file in VS Code.

## Contribution
If you would like to contribute to the project, feel free to open a Pull Request or report issues in the Issues tab.

