# 🚀 Task Management Backend

A minimal, production-ready backend for managing tasks, users, and comments — built with **Go (Gin)**, featuring JWT authentication, role-based access control (RBAC), and soft delete functionality.

---

## 🧠 Overview

This project demonstrates clean backend engineering practices:

- **Clean layered architecture** — handler → service → repository
- **Secure authentication** using JWT + bcrypt
- **Role-based access control** — Admin and User roles with distinct permissions
- **Task & comment management** with ownership enforcement
- **Soft delete** for users, preserving data integrity

---

## ⚙️ Tech Stack

| Layer | Technology |
|---|---|
| Language | Go (Golang) |
| Framework | Gin |
| Database | PostgreSQL |
| DB Layer | `sqlx` |
| Authentication | JWT |
| Password Hashing | bcrypt |
| API Docs | Swagger / OpenAPI |

---

## 📁 Project Structure

```
.
├── cmd/
│   └── server/
│       └── main.go          # Entry point
└── internal/
    ├── config/              # App configuration
    ├── db/                  # Database connection
    ├── models/              # Data models
    ├── repository/          # Data access layer
    ├── service/             # Business logic
    ├── handler/             # HTTP handlers
    ├── middleware/          # Auth & RBAC middleware
    └── utils/               # Shared utilities
```

---

## 🔐 Authentication Flow

1. User registers with email & password
2. Password is hashed with **bcrypt** before storage
3. On login, a **JWT token** is issued
4. All protected routes require the token in the `Authorization` header:

```
Authorization: Bearer <token>
```

---

## 🛂 Role-Based Access Control (RBAC)

### Admin
- Create, update, and delete tasks
- View all tasks and comments across all users
- Soft delete users

### User
- View only their own assigned tasks
- Update the status of their own tasks
- Add comments to their own tasks

---

## 📋 API Endpoints

### 🔐 Auth
| Method | Endpoint | Description |
|---|---|---|
| POST | `/auth/register` | Register a new user |
| POST | `/auth/login` | Login and receive JWT |

### 👤 Users
| Method | Endpoint | Access |
|---|---|---|
| GET | `/me` | Authenticated user |
| GET | `/admin/users` | Admin only |
| PATCH | `/admin/users/:id/soft-delete` | Admin only |

### 📌 Tasks
| Method | Endpoint | Access |
|---|---|---|
| GET | `/tasks` | User → own tasks; Admin → all |
| PATCH | `/tasks/:id/status` | User → own tasks |
| POST | `/admin/tasks` | Admin only |
| DELETE | `/admin/tasks/:id` | Admin only |

### 💬 Comments
| Method | Endpoint | Description |
|---|---|---|
| POST | `/tasks/:id/comments` | Add a comment to a task |
| GET | `/tasks/:id/comments` | Get all comments for a task |

---

## 📖 API Documentation (Swagger)

This project includes interactive **Swagger UI** documentation powered by OpenAPI.

Once the server is running, visit:

```
http://localhost:8080/swagger/index.html
```

The Swagger UI covers all available endpoints grouped by resource:

| Group | Endpoints |
|---|---|
| **Tasks** | `GET /admin/tasks`, `POST /admin/tasks`, `PATCH /admin/tasks/{id}/status` |
| **Admin** | `DELETE /admin/users/{id}` |
| **Auth** | `POST /auth/login`, `POST /auth/register` |
| **Comments** | `POST /tasks/{id}/comments` |

### Authenticating in Swagger

1. Login via `POST /auth/login` to get your JWT token
2. Click the **Authorize 🔒** button at the top right
3. Enter your token as: `Bearer <your_token>`
4. All protected routes will now include your credentials automatically

---



Users are never physically removed from the database. Instead:

- `is_active` is set to `false`
- The user can no longer log in
- All associated tasks and comments remain intact
- Foreign key integrity is fully preserved

---

## 🧠 Design Decisions

**Why Go?**  
Chosen for its performance, simplicity, strong typing, and excellent concurrency primitives — ideal for building scalable APIs.

**Why layered architecture?**  
Separating the handler, service, and repository layers improves testability, maintainability, and makes the codebase easier to extend.

**Why JWT?**  
Stateless authentication that scales horizontally without requiring server-side session storage.

**Why soft delete?**  
Preserves historical data and audit trails while revoking access, without breaking relational integrity.

**Why middleware-based RBAC?**  
Centralizes access control logic, making it reusable, consistent, and easy to audit or modify.

---

## ⚡ Setup & Installation

### 1. Clone the Repository

```bash
git clone <your-repo-link>
cd credes-backend
```

### 2. Configure Environment Variables

Create a `.env` file in the project root:

```env
PORT=8080
DB_URL=postgres://postgres:password@localhost:5432/taskdb?sslmode=disable
JWT_SECRET=your_secret_here
```

### 3. Set Up the Database

Ensure PostgreSQL is running, then create the database:

```sql
CREATE DATABASE taskdb;
```

### 4. Run the Server

```bash
go run cmd/server/main.go
```

The server starts at:

```
http://localhost:8080
```

---

## 🧪 Testing

The API was manually tested using **Postman**, covering the following scenarios:

- Valid and invalid authentication flows
- Role-based access restrictions (Admin vs User)
- Task ownership validation
- Soft delete behavior and post-delete access denial

---

## 🚀 Future Improvements

- [ ] Pagination and filtering for list endpoints
- [ ] Unit and integration test coverage
- [ ] Rate limiting and response caching
- [ ] Refresh token support

---

## 👨‍💻 Author

**Aditya Sanskar Srivastav**

---

> This project prioritizes clean backend design, security, and correctness over premature optimization.
