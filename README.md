### InfraVault

#### Details

To build a secure, self-hosted, multi-user image vault by architecting a resilient, scalable, and highly automated infrastructure. The project will demonstrate mastery over the entire lifecycle of a modern distributed application, from initial deployment and CI/CD to production-grade monitoring and security.

This project will be guided by the following architectural principles:

*   **Automation First:** All infrastructure provisioning, configuration, and application deployment will be automated and defined as code. Manual intervention will be eliminated wherever possible.
*   **Design for Failure:** The system will be architected to be resilient. The failure of any single component should not result in a total system outage.
*   **Security by Design:** Security is not an afterthought. Secure practices for data, access, and secrets management will be integrated at every stage.
*   **Measurability is Key:** The system's health, performance, and status will be centrally monitored and visualized, enabling proactive maintenance and data-driven decisions.

#### Features

*   **Authentication:** Users can register and log in securely via their Google account (OAuth 2.0).
*   **Image Upload:** Authenticated users can upload image files, providing a mandatory title and optional description.
*   **Personal Gallery:** Users can view, manage, and delete their own uploaded images.
*   **Public Gallery:** All users can view a gallery of all images uploaded by all users.

#### Architectural Roadmap

This project is divided into four distinct phases, each building upon the last.

---

##### **Phase 1: The Monolithic Foundation (MVP)**
*Goal: Establish a working, containerized application on a single Linux server.*

*   **Architecture:** A single Virtual Private Server (VPS) running multiple Docker containers managed by Docker Compose. Nginx acts as a reverse proxy, directing traffic to the backend API and serving the frontend application.
*   **Components & Technology:**
    *   **Containerization:** Docker & Docker Compose
    *   **Reverse Proxy:** Nginx
    *   **Backend API:** Go
    *   **Database:** PostgreSQL
    *   **File Storage:** MinIO
    *   **Frontend:** React
    *   **Host OS:** Linux (e.g., Ubuntu 22.04 LTS)
    *   **Security:** HTTPS via Let's Encrypt (managed by Certbot)
    *   **Domain:** Duck DNS (Free)
*   **Deployment Method:** Manual Git pull followed by shell scripts executed over SSH.
*   **Monitoring Method:** Standard Linux tools (`journalctl`, `docker logs`, `htop`, `df`, `du`).

---

##### **Phase 2: Automation & Repeatability (CI/CD & IaC)**
*Goal: Fully automate the provisioning of infrastructure and the deployment of the application.*

*   **Architecture:** The application remains on a single VPS, but the server itself is now provisioned and configured automatically. Application updates are deployed via a CI/CD pipeline.
*   **Evolution of Components & Technology:**
    *   **Infrastructure Provisioning:** **Terraform** will be used to create the VPS and necessary cloud resources (e.g., firewall rules, DNS records) on a cloud provider like AWS, Google Cloud, or DigitalOcean.
    *   **Configuration Management:** **Ansible** will be used to configure the provisioned VPS, installing Docker, Docker Compose, and other required packages.
    *   **CI/CD Pipeline:** **GitHub Actions** will manage the entire build-to-deploy workflow.
        *   **Trigger:** On push to `main` branch.
        *   **Jobs:**
            1.  **Test:** Lint and run unit tests for the Go API and React frontend.
            2.  **Build & Push:** Build production-ready Docker images for the API and frontend. Push images to a container registry (**Docker Hub** or **GitHub Container Registry**).
            3.  **Deploy:** Securely SSH into the target server and execute a script to pull the new images and restart the Docker Compose stack.
    *   **Deployment Method:** Fully automated via GitHub Actions.

---

##### **Phase 3: High Availability & Scalability (Distributed Systems)**
*Goal: Evolve the single-server architecture into a resilient, multi-node cluster that can tolerate failures.*

*   **Architecture:** The application is migrated from Docker Compose to a lightweight Kubernetes cluster spanning 2-3 VPS nodes. Services are replicated for redundancy.
*   **Evolution of Components & Technology:**
    *   **Container Orchestration:** **K3s** will be used to create a multi-node Kubernetes cluster. Docker Compose is replaced by Kubernetes manifest files (YAML for Deployments, Services, Ingress, etc.).
    *   **High Availability - API:** The Go backend will be deployed as a Kubernetes `Deployment` with at least 2 replicas, ensuring the API remains available if one node fails.
    *   **High Availability - Database:** **PostgreSQL** will be configured in a primary-replica streaming replication setup using a Kubernetes `StatefulSet`.
    *   **High Availability - File Storage:** **MinIO** will be deployed in its distributed mode across multiple cluster nodes, providing data redundancy for uploaded images.
    *   **Load Balancing:** A Kubernetes **Ingress Controller** (e.g., Traefik, built into K3s) will replace the standalone Nginx proxy to manage external traffic and load balance requests across the API replicas.

---

##### **Phase 4: Production-Grade Observability & Security**
*Goal: Implement a comprehensive monitoring, logging, and security solution for the entire distributed system.*

*   **Architecture:** A centralized "observability stack" is deployed within the Kubernetes cluster to collect data from all components. Secrets management is hardened.
*   **Evolution of Components & Technology:**
    *   **Metrics Collection & Visualization:**
        *   **Prometheus:** Deployed to scrape metrics from all cluster nodes and application services.
        *   **Grafana:** Deployed to create dashboards for visualizing key health and performance metrics (CPU/memory usage, API latency, database connections, etc.).
    *   **Alerting:** **Alertmanager** (part of the Prometheus ecosystem) will be configured to send alerts (e.g., via email or Slack) based on predefined rules (e.g., service down, high error rate).
    *   **Centralized Logging:** The **Loki Stack** will be deployed.
        *   **Promtail:** Collects logs from all containers in the cluster.
        *   **Loki:** Aggregates and stores log streams efficiently.
        *   Logs will be searchable and viewable within **Grafana**.
    *   **Secrets Management:** **HashiCorp Vault** or Kubernetes Secrets will be used to securely store and inject sensitive data (database passwords, API keys, OAuth credentials) into application pods at runtime.
    *   **Container Security:** **Trivy** will be added to the CI/CD pipeline to automatically scan Docker images for known vulnerabilities before they are deployed.

#### **5.0 Final Technology Stack Summary**

| Category                  | Technology                                     |
| ------------------------- | ---------------------------------------------- |
| **Cloud & Provisioning**  | Any IaaS Provider (AWS, GCP, etc.), Terraform  |
| **Configuration Mgmt**    | Ansible                                        |
| **Orchestration**         | K3s (Kubernetes)                               |
| **CI/CD**                 | GitHub Actions, Docker Hub / GHCR              |
| **Containerization**      | Docker                                         |
| **Ingress & Proxy**       | K3s's built-in Traefik Ingress Controller      |
| **Backend API**           | Go                                             |
| **Frontend**              | React                                          |
| **Database**              | PostgreSQL (with streaming replication)        |
| **Object Storage**        | MinIO (in distributed mode)                    |
| **Monitoring & Metrics**  | Prometheus, Grafana, Alertmanager            |
| **Logging**               | Loki, Promtail                                 |
| **Security**              | Let's Encrypt, HashiCorp Vault, Trivy          |

#### **6.0 Portfolio Outcomes & Key Learning Objectives**

Upon completion, this project will demonstrate:

*   **End-to-End Application Lifecycle Management:** From code commit to a fully monitored, production-grade deployment.
*   **Infrastructure as Code (IaC):** Proficiency in provisioning and managing cloud infrastructure declaratively using Terraform and Ansible.
*   **Advanced CI/CD:** The ability to design and implement automated testing and deployment pipelines.
*   **Container Orchestration:** Deep, practical experience with Kubernetes for deploying and managing distributed applications.
*   **Distributed System Design:** An understanding of high availability, data redundancy, and load balancing principles.
*   **Comprehensive Observability:** Expertise in setting up and using modern monitoring and logging stacks (Prometheus, Grafana, Loki).
*   **DevSecOps Practices:** Competency in integrating security tooling for vulnerability scanning and secrets management.

