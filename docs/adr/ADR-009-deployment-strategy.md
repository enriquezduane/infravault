# ADR-009: Deployment Strategy (Manual Git + SSH)

## Context
InfraVault Phase 1 requires a deployment strategy that can:
- Deploy the containerized application to a production VPS
- Update the application with new code changes
- Maintain system reliability during deployments
- Provide rollback capabilities in case of issues
- Be simple enough for MVP while establishing patterns for automation
- Support the single-server Docker Compose architecture
- Enable development team to deploy updates efficiently

The deployment strategy must be reliable, repeatable, and provide a foundation for the automated CI/CD pipeline planned for Phase 2.

## Decision
We will implement a **manual deployment strategy** using **Git version control** and **SSH-based deployment scripts** for Phase 1.

### Deployment Architecture:
- **Version Control**: Git repository with production branch
- **Deployment Method**: SSH-based shell scripts executed on production server
- **Application Updates**: Git pull + Docker Compose restart cycle
- **Environment Management**: Environment-specific configuration files
- **Rollback Strategy**: Git-based rollback with container image tagging