# Job Application Management API

This is a RESTful API for managing job applications. It includes user authentication, job creation, retrieval, updating, and deletion functionalities.

### Features

- User Authentication:

    - Register and login with JWT-based authentication.

- Job Management:

    - Create, retrieve, update, and delete job applications.

- Health Check:

    - API endpoint for checking server health.

- JWT Middleware:

    - Protect sensitive routes with JWT validation.

### Requirements

- Go (Golang) 1.20+

- PostgreSQL (or another supported database with GORM)

- Fiber web framework

- UUID generation library

- Validator for input validation

## Setup Instructions

### Prerequisites

Before you begin, ensure you have met the following requirements:

- GO(Golang): You should have Go installed. [Download GO](https://go.dev/doc/install)
- Postgresql [Download Postgresql](https://www.postgresql.org/download/)

### Installation

Clone the repo and install dependencies

```shell
git clone git@github.com:udodinho/job_app_go.git
cd job_app_go
```

```shell
$ go mod tidy
```

### Setup environment

```shell
DB_HOST=localhost
DB_PORT=
DB_USER=
DB_PASS=
DB_NAME=
DB_SSLMODE=disable

JWT_SECRET_KEY=
JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT=
JWT_REFRESH_KEY=
JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT=
```

### Start server

```shell
$ go run main.go
```

The API will be available at http://localhost:8080

## API Endpoints

### Health Check

- GET /api/v1/healthcheck

    - Response: { "msg": "OK" }

### Authentication

- POST /api/v1/auth/register

    - Register a new user.

    - Body:

    ```shell
    {
    "email": "string",
    "password": "string"
    }
    ```

- POST /api/v1/auth/login

    - Log in an existing user.

    - Body:

    ```shell
    {
    "email": "string",
    "password": "string"
    }
    ```

## Job Management

### Authenticated Routes

JWT token is required for all routes below:

- POST /api/v1/job

    - Create a new job.

    - Body:

    ```shell
    {
    "company": "string",
    "position": "string",
    "status": "string"
    }
    ```

- GET /api/v1/job

    - Retrieve all jobs created by the authenticated user.

- GET /api/v1/job/:id

    - Retrieve a single job by ID.

- PATCH /api/v1/job/:id

    - Update a job by ID.

    - Body:

    ```shell
    {
    "company": "string",
    "position": "string",
    "status": "string"
    }
    ```
- DELETE /api/v1/job/:id

    - Delete a job by ID.

### Error Handling

The API returns JSON responses for errors. Example:

 ```shell
    {
    "error": true,
    "msg": "Error message"
    }
    ```
