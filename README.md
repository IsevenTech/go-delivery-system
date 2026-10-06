# Go Delivery System

A simple delivery management REST API built with Go.

## 🚀 Features

- Create delivery orders
- List all orders
- Automatic order IDs
- Order status management
- JSON REST API
- In-memory data storage

## 🛠️ Technologies

- Go
- REST API
- HTTP
- JSON

## 📁 Project Structure

```text
go-delivery-system/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handler/
│   │   └── order.go
│   └── model/
│       └── order.go
├── go.mod
└── README.md 
git clone https://github.com/IsevenTech/go-delivery-system.git
cd go-delivery-system
go run ./cmd/api
http://localhost:8080
GET /
POST /orders
{
  "customer": "Joao",
  "address": "Dublin, Ireland"
}
GET /orders
{
  "id": 1,
  "customer": "Joao",
  "address": "Dublin, Ireland",
  "status": "pending"
}


