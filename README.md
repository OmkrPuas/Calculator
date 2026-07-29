# Full Stack Calculator

A full-stack calculator application with dual UI modes and a Go backend calculator API.

## Overview

This repository contains a React frontend and a Go backend that exposes a `/calculate` REST API. The frontend supports two calculator styles: a traditional calculator and a modern calculator, both with value entry, operator selection, and result display.

## Features

- Arithmetic operations: add, subtract, multiply, divide
- Extended operations: exponentiation, square root, percentage
- Dual frontend UI modes: Traditional and Modern
- Responsive layout with mobile-only controls
- API validation and error handling
- Docker deployment support

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

For a detailed profile:

```bash
go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out
```

Frontend

```bash
cd frontend/calculator
npm run coverage
```

---

# Docker

Build and run the single combined container:

```bash
docker build -t calculator-app .
docker run -p 8080:8080 -e STATIC_DIR=/app/frontend/dist calculator-app
```

The app will be available at:

```bash
http://localhost:8080
```

Or use Docker Compose:

```bash
docker compose up --build
```

This starts the combined backend and frontend container on port 8080.

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