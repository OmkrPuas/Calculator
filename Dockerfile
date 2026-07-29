# Build Go backend
FROM golang:1.20-alpine AS backend-builder
WORKDIR /app/backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download
COPY backend/. ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /calculator-backend ./cmd/main.go

# Build React frontend
FROM node:20-alpine AS frontend-builder
WORKDIR /app/frontend
COPY frontend/calculator/package.json frontend/calculator/package-lock.json ./
RUN npm install
COPY frontend/calculator/. ./
RUN npm run build

# Final image with backend serving frontend
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=backend-builder /calculator-backend ./calculator-backend
COPY --from=frontend-builder /app/frontend/dist ./frontend/dist
ENV STATIC_DIR=/app/frontend/dist
EXPOSE 8080
CMD ["./calculator-backend"]
