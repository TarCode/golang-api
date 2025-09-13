# Go Todo API

A simple REST API for managing todo items built with Go Fiber.

## Features

- RESTful API endpoints for todo management
- CORS enabled for cross-origin requests
- Request logging middleware
- Database integration

## Quick Start

```bash
go mod tidy
go run main.go
```

The API will start on `http://localhost:5000`

## API Endpoints

All endpoints are prefixed with `/api`

Example endpoints:
- `GET /api/todos` - Get all todos
- `POST /api/todos` - Create a new todo
- `PUT /api/todos/:id` - Update a todo
- `DELETE /api/todos/:id` - Delete a todo

## Dependencies

- [Fiber](https://github.com/gofiber/fiber) - Web framework
- [GORM](https://github.com/jinzhu/gorm) - ORM for database operations
- Custom database package for data persistence
- Custom todo package for business logic with repository pattern

## Structure

```
├── main.go          # Application entry point
├── database/        # Database connection and models
└── todo/           # Todo handlers and routes
```
