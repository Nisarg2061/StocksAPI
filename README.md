# Stocks Management System
This project is a simple stocks management system using Go and PostgreSQL.

## Project Structure
```
├── .github               # GitHub-related configuration files
├── middleware            # Middlewares for handling requests
│   ├── handlers.go       # Request handlers
│   └── middleware.go
├── models                # Database models
│   └── models.go
├── router                # API routing
│   └── router.go
├── .env                  # Environment variables
├── Makefile              # Makefile for running commands
├── go.mod                # Go modules file
├── go.sum                # Dependencies checksum
└── main.go               # Application entry point
```

## Prerequisites

- Go (version 1.16 or higher)
- PostgreSQL
- Environment variables specified in the `.env` file

## Installation and Setup
1. Clone the repository:
```
git clone https://github.com/Nisarg2061/StocksAPI.git
cd StocksAPI
```

2. Set up the database:
Make sure PostgreSQL is running locally. You can use the script.sql file to set up the database:
```
psql -h localhost -U postgres -d stocksdb -a -f script.sql
```
This will create the required tables and populate initial data into the stocksdb database.

3. Set up environment variables
Create a .env file in the root directory and set the required variables:
```
POSTGRES_URL=your_postgres_connection_url
```
Make sure to replace `your_postgres_connection_url` with your actual PostgreSQL connection string (e.g., `postgres://username:password@localhost:5432/stocksdb`).

4. Run the application
You can use the Makefile to build and run the project. To start the server, run:
```
make run
```
The application should now be running on the configured port.

## Usage

- The API routes are defined in router/router.go.
- Handlers for various routes are in handlers.go.
- Middleware logic is found in middleware/middleware.go.
- Database models are in models/models.go.
