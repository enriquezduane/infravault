# ADR-005: File Storage Solution (MinIO)

## Context
InfraVault Phase 1 requires a reliable file storage solution for user-uploaded images that provides:
- Secure storage and retrieval of image files
- S3-compatible API for standardized object storage operations
- Integration with the Go backend for upload/download operations
- Scalable storage that can handle various file sizes and formats
- Self-hosted solution aligned with project goals
- Data persistence and backup capabilities
- Access control and security for stored files

The storage solution must be containerizable, maintainable, and provide a solid foundation for future scaling while keeping operational complexity appropriate for a single-server MVP.

## Decision
We will use **MinIO** as the object storage solution, deployed as a Docker container within the Docker Compose stack.

### MinIO Configuration:
- **Deployment**: Single-node MinIO server in standalone mode
- **API Compatibility**: S3-compatible REST API
- **Data Persistence**: Docker volume for persistent storage
- **Access Control**: Bucket policies and IAM for security
- **Integration**: Go MinIO client SDK for backend integration
- **Management**: MinIO Console for administrative operations