# Go Auth API

A RESTful authentication API built with **Go (Golang)**, designed to provide a foundation for secure user authentication, JWT-based authorization, and database-backed user management.

## 🚀 Features

- **User Authentication** — API endpoints for user authentication and account management.
- **JWT Authentication** — JSON Web Token configuration for token-based authentication.
- **Database Integration** — Database connection and configuration management.
- **Database Migrations** — Database schema management using migration scripts.
- **Modular Architecture** — Organized codebase with separate configuration, handlers, repositories, and helpers.
- **Environment Configuration** — Environment variables for application and database settings.

## 🛠️ Tech Stack

- **Language:** Go (Golang)
- **Authentication:** JWT
- **Database:** Configure according to your database driver and connection settings
- **Database Migrations:** Migration scripts
- **Version Control:** Git and GitHub

## 📁 Project Structure

```text
Go-Auth-API/
├── cmd/
│   └── api/
│       └── main.go
├── config/
│   └── database and JWT configuration
├── internal/
│   ├── handlers/
│   ├── repositories/
│   ├── services/
│   └── helpers/
├── tmp/
├── .env
├── .gitignore
├── go.mod
├── go.sum
├── migrate.bat
└── README.md
```

*Note: The internal folder names above are illustrative. Adjust them to match your actual directory structure.*

## ⚙️ Prerequisites

Before running the project, ensure you have:

- [Go](https://go.dev/dl/) installed.
- A supported database server configured for your application.
- Git installed.

## 📥 Getting Started

### 1. Clone the repository

```bash
git clone https://github.com/harijp11/Go-Auth-API.git
cd Go-Auth-API
```

### 2. Install dependencies

```bash
go mod download
```

### 3. Configure environment variables

Create a `.env` file in the project root and configure the variables required by your application.

Example:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=your_database

JWT_SECRET=your_secure_secret
```

**Important:** Update these variables to match the names expected by your configuration code. The database example assumes PostgreSQL; change the settings if you use another database. Never commit real credentials or JWT secrets to GitHub.

### 4. Run database migrations

If your migration scripts use a supported migration tool, configure the database connection and run the migrations using your project's migration workflow.

On Windows, you can inspect and run the included batch script:

```powershell
.\migrate.bat
```

Ensure the script is configured correctly before executing it.

### 5. Start the API

```bash
go run ./cmd/api
```

The server will start according to the host and port configured in your application.

## 🔐 Authentication Workflow

The intended authentication workflow is:

1. A user submits their credentials.
2. The API validates the credentials.
3. The application generates a JWT after successful authentication.
4. The client sends the token with subsequent requests.
5. Protected endpoints validate the token before processing requests.

Implement and configure password hashing, token expiration, and authorization checks as appropriate for your application.

## 🧪 Testing

Run the project's Go tests with:

```bash
go test ./...
```

## 🔒 Security Best Practices

- Store secrets in environment variables.
- Never commit `.env` files containing real credentials.
- Hash passwords using a suitable password-hashing algorithm.
- Validate JWT signatures, expiration, and claims.
- Validate incoming requests and handle errors safely.
- Use HTTPS in production.

## 🛣️ Roadmap

- [ ] Complete user registration and login workflows.
- [ ] Add JWT authentication middleware.
- [ ] Implement refresh-token handling if required.
- [ ] Add comprehensive unit and integration tests.
- [ ] Add Docker support.
- [ ] Set up CI/CD automation.

## 👨‍💻 Author

**Harij Sagar J S**

- GitHub: [@harijp11](https://github.com/harijp11)

---

## Local Development

1. Clone the repository.
2. Configure the required environment variables.
3. Run database migrations.
4. Start the API using `go run ./cmd/api`.

*Built as a Go backend project to explore authentication, API development, and backend architecture.*
