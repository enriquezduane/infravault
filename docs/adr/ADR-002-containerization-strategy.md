# ADR-002: Containerization Strategy (Docker & Docker Compose)

## Context
InfraVault Phase 1 requires a containerized deployment strategy that provides service isolation, consistent environments, and simplified deployment while maintaining the monolithic architecture decision from ADR-001.

Requirements:
- Consistent development and production environments
- Service isolation for frontend, backend, database, and file storage
- Simplified orchestration for single-server deployment
- Easy local development setup
- Preparation for future container orchestration (Phase 3)

## Decision
We will use **Docker** for containerization and **Docker Compose** for orchestration.

### Container Strategy:
1. **Multi-container Application**: Each major service in its own container
2. **Docker Compose Orchestration**: Single `docker-compose.yml` file defines all services
3. **Custom Dockerfiles**: Application-specific containers for frontend and backend
4. **Official Base Images**: Use official images for third-party services (PostgreSQL, MinIO, Nginx)

#### Database Service (PostgreSQL):
- Use official `postgres` latest image
- Environment variables for configuration
- Named volume for data persistence

#### Storage Service (MinIO):
- Use official `minio` latest image
- S3-compatible API for file operations
- Named volume for data persistence

#### Reverse Proxy Service (Nginx):
- Use official `nginx:alpine` image
- Custom configuration for routing
- SSL certificate volume mounts