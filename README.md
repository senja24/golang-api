# Golang API with Gin and GORM

This is a simple RESTful API built with Go using the Gin web framework and GORM for database interactions. The API allows for creating and retrieving posts.

## Features

- Create new post
- Get post by id
- Edit post by id
- Delete post by id
- Retrieve all posts

## Technologies Used

- [Golang](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/index.html)
- [MySQL](https://www.mysql.com/)

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 22.5 or higher
- [MySQL](https://www.mysql.com/downloads/)

### Installation

1. Clone the repository:

```bash
git clone https://github.com/senja24/golang-api.git
cd golang-api
```

2. Install dependencies
3. Set up database
4. Run the application:
```bash
go run main.go
```

### API Endpoints

1. Get all posts
```bash
GET /api/posts
```

2. Create a new post
```bash
POST /api/posts
```

3. Get post by id
```bash
GET /api/posts/:id
```

4. Edit post by id
```bash
PUT /api/posts/:id
```

4. Delete post by id
```bash
DELETE /api/posts/:id
```

### Project Structure

.
├── controllers
│   └── postsController.go
├── models
│   ├── post.go
│   └── setup.go
├── main.go
└── go.mod
