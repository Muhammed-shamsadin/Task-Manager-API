# Task Manager API (Go)

A simple RESTful API for managing tasks, built with Go, Gin, and MongoDB.

## Features
- Create, read, update, and delete tasks
- MongoDB integration
- RESTful endpoints using Gin
- Environment-based configuration

## Project Structure
```
.
├── controllers/      # HTTP handlers
├── data/             # Database logic and services
├── documentation/    # Project documentation
├── models/           # Data models'
├── router/           # Gin router setup
├── main.go           # Application entry point
├── go.mod, go.sum    # Go modules
├── .env              # Environment variables (not committed)
```

## Getting Started

### Prerequisites
- Go 1.18+
- MongoDB instance (local or Atlas)

### Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/<your-username>/<repo-name>.git
   cd <repo-name>
   ```
2. Create a `.env` file in the root directory and add your MongoDB URI:
   ```env
   MONGODB_URI=your_mongodb_connection_string
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```
4. Run the application:
   ```sh
   go run main.go
   ```

### API Endpoints
| Method | Endpoint         | Description         |
|--------|------------------|--------------------|
| POST   | /api/tasks       | Create a new task  |
| GET    | /api/tasks       | Get all tasks      |
| GET    | /api/tasks/:id   | Get task by ID     |
| PUT    | /api/tasks/:id   | Update a task      |
| DELETE | /api/tasks/:id   | Delete a task      |


