# ADR-006: Reverse Proxy and Load Balancing (Nginx)

## Context
InfraVault Phase 1 requires a reverse proxy solution that can:
- Route external HTTP/HTTPS traffic to appropriate backend services
- Serve static frontend assets efficiently
- Terminate SSL/TLS connections for HTTPS
- Handle CORS headers for API requests
- Provide basic load balancing capabilities
- Serve as the single entry point for all external traffic
- Support future scaling and service discovery

The solution must be containerizable, performant, and provide a solid foundation for the monolithic Docker Compose architecture while preparing for future distributed deployments.

## Decision
We will use **Nginx** as the reverse proxy and load balancer, deployed as a Docker container within the Docker Compose stack.

### Nginx Configuration:
- **Deployment**: Single Nginx container as the edge proxy
- **SSL Termination**: HTTPS termination with Let's Encrypt certificates
- **Static Serving**: Direct serving of React frontend assets
- **API Proxying**: Reverse proxy for Go backend API requests
- **Load Balancing**: Upstream configuration for backend services
- **Security**: Security headers and rate limiting