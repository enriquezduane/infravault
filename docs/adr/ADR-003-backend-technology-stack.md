# ADR-003: Backend Technology Stack (Go & PostgreSQL)

## Context
InfraVault Phase 1 requires a backend API that can handle:
- User authentication and session management
- Image metadata CRUD operations
- File upload/download coordination with MinIO
- RESTful API endpoints for frontend consumption
- Database operations for users and image metadata
- Integration with Google OAuth 2.0
- High performance and reliability

The backend must be containerizable, maintainable, and provide a solid foundation for future phases while keeping complexity appropriate for an MVP.

## Decision
We will implement the backend using **Go (Golang)** with **PostgreSQL** as the primary database.

### Technology Choices:

#### Programming Language: Go 1.21+
- **Web Framework**: Go Standard Library `net/http` package
- **Authentication**: JWT tokens with Google OAuth 2.0
- **Configuration**: os.Getenv
- **Logging**:`log/slog` package
- **Testing**: Go's built-in testing package
- **Database Driver**: `jackc/pgx/v5` package

#### Database: PostgreSQL 15
- **ACID Compliance**: Full transaction support
- **JSON Support**: Native JSON column types for flexible data
- **Extensions**: Support for UUID generation, full-text search
- **Connection Pooling**: Built-in connection pooling
- **Backup**: pg_dump for data backup strategies

### API Endpoints:

#### Authentication:
- `GET /api/auth/google` - Initiate Google OAuth flow
- `GET /api/auth/callback` - Handle OAuth callback
- `POST /api/auth/logout` - User logout
- `GET /api/auth/me` - Get current user info

#### Images:
- `POST /api/images` - Upload new image
- `GET /api/images` - Get user's images (paginated)
- `GET /api/images/public` - Get all public images (paginated)
- `GET /api/images/:id` - Get specific image metadata
- `PUT /api/images/:id` - Update image metadata
- `DELETE /api/images/:id` - Delete image
- `GET /api/images/:id/download` - Download image file

#### Users:
- `GET /api/users/profile` - Get user profile
- `PUT /api/users/profile` - Update user profile