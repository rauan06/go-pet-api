# Assignment 2

## Setup
```bash
go mod tidy
```

## Initialization
```bash
go run cmd/api/main.go
```

## API Deployment
🚀 **Live API available at:**  
[![Render](https://img.shields.io/badge/deployment-Render-blue)](https://go-pet-api.onrender.com)

You can access the live API at:  
🔗 [Go Pet API on Render](https://go-pet-api.onrender.com)

## API Documentation
📖 **Interactive API Docs available on SwaggerHub:**  
[![SwaggerHub](https://img.shields.io/badge/documentation-SwaggerHub-blue)](https://app.swaggerhub.com/apis/hobby-ec8/books_weather/1.0.0)

You can view and test the API interactively using **SwaggerHub**:  
🔗 [Click here to open Swagger Documentation](https://app.swaggerhub.com/apis/hobby-ec8/books_weather/1.0.0)

---

## API Endpoints

### 📚 Get Books
```http
GET /books HTTP/1.1
Host: go-pet-api.onrender.com
```

### ➕ Add a Book
```http
POST /books HTTP/1.1
Host: go-pet-api.onrender.com
Content-Type: application/json

{
  "title": "New Book",
  "author": "John Doe",
  "year": 2024,
  "genre": "Fiction"
}
```

### ✏️ Update a Book
```http
PUT /book/{id} HTTP/1.1
Host: go-pet-api.onrender.com
Content-Type: application/json

{
  "title": "Updated Book",
  "author": "Jane Doe",
  "year": 2025,
  "genre": "Non-fiction"
}
```

### 🗑️ Delete a Book
```http
DELETE /book/{id} HTTP/1.1
Host: go-pet-api.onrender.com
```

### 🌦️ Get Weather for a City
```http
GET /weather/{city} HTTP/1.1
Host: go-pet-api.onrender.com
```

---

## Additional Notes
- The API follows RESTful principles.
- To interact with the API, visit **SwaggerHub** and use the "Try it out" button.

---
