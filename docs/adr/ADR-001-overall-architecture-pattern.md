# ADR-001: Overall Architecture Pattern (Phase 1)

## Context
InfraVault Phase 1 requires establishing a working, containerized multi-user image vault application. This is the foundational phase (MVP) that needs to demonstrate core functionality while keeping complexity manageable for initial development and deployment.

The application needs to support:
- User authentication via Google OAuth 2.0
- Image upload with metadata (title, description)
- Personal image galleries for users
- Public gallery viewing all uploaded images
- Secure file storage and retrieval

## Decision
We will implement a **monolithic architecture** deployed via **Docker Compose** on a single Virtual Private Server (VPS).

### Architecture Components:
1. **Single VPS Deployment**: All services run on one Linux server
2. **Docker Compose Orchestration**: Multi-container application managed by Docker Compose
3. **Service Separation**: Each major component runs in its own container
4. **Reverse Proxy**: Nginx container handles external traffic routing

### Container Structure:
- **Frontend Container**: React application served by Nginx
- **Backend Container**: Go API server
- **Database Container**: PostgreSQL for user and image metadata
- **File Storage Container**: MinIO for image file storage
- **Reverse Proxy Container**: Nginx for traffic routing and SSL termination