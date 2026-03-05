## Setup Instructions

### 1. Clone repository

git clone https://github.com/yourusername/country-search-api.git

### 2. Install dependencies

go mod tidy

### 3. Run application

go run cmd/server/main.go

Server will start at:

http://localhost:8000

### 4. TEST API

curl http://localhost:8000/api/countries/search?name=India


## Running Tests

Run all tests

go test ./...

Run with race detection

go test -race ./...

Check coverage

go test ./... -cover