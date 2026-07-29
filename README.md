# Full Stack Calculator

A full-stack calculator application built as part of a technical assessment.

The project consists of:

- **Frontend:** React + TypeScript + Vite
- **Backend:** Go REST API
- **Communication:** HTTP REST returning JSON

---

# Features

## Supported Operations

- Addition
- Subtraction
- Multiplication
- Division

Optional (planned):

- Exponentiation
- Square Root
- Percentage

---

# Tech Stack

## Frontend

- React
- TypeScript
- Vite
- Axios
- React Hooks

## Backend

- Go
- net/http
- JSON REST API

## Testing

Frontend

- Vitest
- React Testing Library

Backend

- Go testing package

---

# Project Structure

```
calculator-fullstack/
│
├── frontend/
├── backend/
├── docs/
└── prompts/
```

---

# Running the Project

## Backend

```bash
cd backend

go mod tidy

go run cmd/main.go
```

The API will start on:

```
http://localhost:8080
```

---

## Frontend

```bash
cd frontend

npm install

npm run dev
```

Application:

```
http://localhost:5173
```

---

# API

## POST /calculate

### Request

```json
{
  "operation": "add",
  "a": 5,
  "b": 10
}
```

### Response

```json
{
  "result": 15
}
```

---

### Operations

| Operation | Value |
|------------|---------|
| Addition | add |
| Subtraction | subtract |
| Multiplication | multiply |
| Division | divide |

---

# Error Example

Division by zero

```json
{
    "error":"division by zero"
}
```

---

# Testing

Backend

```bash
go test ./...
```

Frontend

```bash
npm test
```

Coverage

Backend

```bash
go test -cover ./...
```

Frontend

```bash
npm run coverage
```

---

# Design Decisions

## Backend

The backend follows a layered architecture:

- HTTP Handlers
- Business Logic (Services)
- Models
- Validation

This separation improves maintainability and testability.

## Frontend

The frontend separates:

- UI Components
- API layer
- Types
- Utility functions

Business logic is kept outside presentation components whenever possible.

---

# Assumptions

- Calculator operates on floating-point numbers.
- Validation is performed in both frontend and backend.
- The backend is the single source of truth for calculations.
- Errors are returned using appropriate HTTP status codes and JSON messages.

---

# Future Improvements

- Calculation history
- Scientific calculator functions
- Docker Compose deployment
- CI/CD pipeline
- Authentication
- Dark mode

---

# AI Usage

AI tools were used to assist with:

- Initial architecture planning
- Boilerplate generation
- Documentation drafting
- Code review suggestions

All generated code was reviewed, modified, and tested manually.

---

# Author

Technical Assessment Submission