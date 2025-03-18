# Task-Project

A Basic Go backend API integration with JSONPlaceholder just to practice my Go skills (https://jsonplaceholder.typicode.com/).

## Features

- RESTful API endpoints that proxy requests to JSONPlaceholder
- Clean architecture with separation of concerns
- Error handling and logging
- Graceful shutdown

## Available Endpoints

- `GET /api/posts` - Get all posts
- `GET /api/posts/{id}` - Get a specific post
- `GET /api/posts/{id}/comments` - Get comments for a specific post
- `GET /api/comments` - Get all comments
- `GET /api/users` - Get all users
- `GET /api/users/{id}` - Get a specific user
- `GET /api/albums` - Get all albums
- `GET /api/photos` - Get all photos
- `GET /api/todos` - Get all todos

## Getting Started

### Prerequisites

- Go 1.16 or higher --> I am on go 1.21

### Installation

1. Clone the repository:
   ```
   git clone https://github.com/rnb03/task-project.git
   cd task-project
   ```

2. Install dependencies:
   ```
   go mod download
   ```

3. Run the server:
   ```
   go run main.go
   ```

4. The server will start on port 8080. You can access the API at `http://localhost:8080/api/`.

## Usage Examples

### Get all posts