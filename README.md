# Language List API

A RESTful API for managing a list of programming languages built with Go, Gin, GORM, and PostgreSQL.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete programming languages
- **PostgreSQL Database**: Robust data persistence with GORM ORM
- **Swagger Documentation**: Interactive API documentation
- **Docker Support**: Easy deployment with Docker Compose
- **Structured Architecture**: Clean separation of concerns with controllers, repositories, and models

## Tech Stack

- **Backend**: Go 1.24, Gin Framework
- **Database**: PostgreSQL 17
- **ORM**: GORM
- **Documentation**: Swagger/OpenAPI
- **Containerization**: Docker & Docker Compose

## Quick Start

### Prerequisites

- Docker and Docker Compose
- Git

### Installation

1. Clone the repository:
```bash
git clone git@github.com:Adrien11111/LanguageList.git
cd LanguageList
```

2. Copy environment variables:
```bash
cp .env.example .env
```

3. Start the application with Docker Compose:
```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`

### Swagger Documentation

Access the interactive API documentation at:
```
http://localhost:8080/swagger/index.html
```

**Note**: When making changes to API endpoints or Swagger annotations, regenerate the documentation by running:
```bash
cd src/
swag init -g main.go -o docs
```

## API Endpoints

### Languages

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/api/v1/languages` | Get all programming languages |
| GET | `/api/v1/languages/{id}` | Get a language by ID |
| POST | `/api/v1/languages` | Create a new language |
| DELETE | `/api/v1/languages/{id}` | Delete a language by ID |

### Example Requests

#### Create a Language
```bash
curl -X POST http://localhost:8080/api/v1/languages \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Go",
    "description": "A statically typed, compiled programming language designed for simplicity and efficiency."
  }'
```

#### Get All Languages
```bash
curl http://localhost:8080/api/v1/languages
```

#### Get Language by ID
```bash
curl http://localhost:8080/api/v1/languages/1
```

#### Delete a Language
```bash
curl -X DELETE http://localhost:8080/api/v1/languages/1
```

## Project Structure

```
.
├── src/
│   ├── controllers/       # HTTP request handlers
│   ├── models/            # Database models
│   ├── repositories/      # Data access layer
│   ├── routes/            # Route definitions
│   ├── middlewares/       # Custom middleware
│   ├── database/          # Database connection
│   ├── docs/              # Swagger documentation
│   └── main.go            # Application entry point
├── build-docker/          # Docker configuration
│   ├── go/                # Go Dockerfile
│   └── database/          # PostgreSQL Dockerfile
├── docker-compose.yaml    # Docker Compose configuration
├── go.mod                 # Go module dependencies
└── README.md
```

## Development

### Local Development without Docker

1. Install Go 1.24+
2. Install PostgreSQL
3. Set up environment variables in `.env`
4. Install dependencies:
```bash
go mod tidy
```

5. Generate Swagger docs:
```bash
cd src/
swag init -g main.go -o docs
```

6. Run the application:
```bash
go run src/main.go
```

### Regenerating API Documentation

After making changes to API endpoints, Swagger annotations, or route definitions, you need to regenerate the documentation:

```bash
cd src/
swag init -g main.go -o docs
```

This command must be executed from the `src/` directory to properly generate the documentation files.

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_HOST | Database host | database |
| DB_PORT | Database port | 5432 |
| DB_USER | Database user | ugoapi |
| DB_PASSWORD | Database password | pgoapi |
| DB_NAME | Database name | goapi |
| SECRET_KEY | JWT secret key | my-secret |
| GIN_MODE | Gin mode (debug/release) | debug |

## Database Schema

### Language Model

```go
type Language struct {
    ID          uint   `json:"id" gorm:"primaryKey"`
    Name        string `json:"name" gorm:"unique;not null"`
    Description string `json:"description" gorm:"not null"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## Contact

**Adrien Panis**
- Email: adrienpani@gmail.com

## Future Enhancements

- [x] Authentication and authorization
- [ ] Language categories and tags
- [ ] Search and filtering capabilities
- [ ] Rating and popularity system
