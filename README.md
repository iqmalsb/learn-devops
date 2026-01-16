# Introduction
This repository serves as a personal learning journal and a collection of exercises as I delve into the world of Docker. My goal is to understand containerization concepts, build and manage Docker images, orchestrate multi-container applications, and ultimately leverage Docker for efficient development and deployment workflows.

# 🚀 DevOps & SRE Learning Roadmap: Enterprise Suite Project

This roadmap focuses on transitioning from a Web Developer mindset to an Infrastructure/SRE mindset by building a multi-service enterprise platform.

---

## 🏗️ Phase 1: The "Service Chassis" (Weeks 1–2)
**Goal:** Build a single, reusable Go backend that can "pretend" to be any service.

### Key Concepts
* **Go (Echo Framework):** Building lightweight REST APIs.
* **Identity Injection:** Using Environment Variables to change service behavior (e.g., `APP_NAME=Payroll`).
* **Health Patterns:** Implementing `/health` (liveness) and `/ready` (readiness) endpoints.
* **Data Persistence:** Using `database/sql` with SQLite for local state.

### Deliverables
- [/] A Go API that returns its name and a simulated "work" response.
- [ ] A simple Vue.js dashboard to trigger API calls and display JSON results.
- [ ] A functional SQLite database file (`enterprise.db`) initialized with basic tables.

---

## 📦 Phase 2: Containerization & State (Weeks 3–4)
**Goal:** Package the apps and manage "Stateful" data with Docker.

### Key Concepts
* **Multi-Stage Builds:** Creating production-ready images (Build in Go, Run in Alpine).
* **Docker Volumes:** Mounting the SQLite `.db` file so data survives container restarts.
* **Container Networking:** Getting the Vue container to talk to the Go container.
* **Orchestration Lite:** Using Docker Compose to spin up the "Departments" locally.

### Deliverables
- [ ] `Dockerfile` for the Go Backend (optimized size).
- [ ] `Dockerfile` for the Vue Frontend.
- [ ] `docker-compose.yml` that launches at least 3 services (HR, Payroll, Sales) sharing a volume.

---

## 🤖 Phase 3: The Automation Pipeline (Weeks 5–6)
**Goal:** Automate the build and registry process.

### Key Concepts
* **GitHub Actions:** Writing workflows to automate testing and building.
* **Container Registries:** Pushing images to GitHub Packages (GHCR) or Docker Hub.
* **Semantic Versioning:** Tagging images with versions (v1.0.1) instead of just `latest`.

### Deliverables
- [ ] A `.github/workflows/main.yml` file.
- [ ] Successful automated push of your "Chassis" image to a registry.

---

## ☸️ Phase 4: Kubernetes Orchestration (Weeks 7–9)
**Goal:** The core of the DevOps journey. Moving from "Containers" to "Orchestration."

### Key Concepts
* **Namespaces:** Creating virtual isolation for HR, Payroll, Finance, etc.
* **Deployments & Services:** Managing the lifecycle and internal DNS naming of your pods.
* **ConfigMaps & Secrets:** Injecting configuration without rebuilding images.
* **Persistent Volume Claims (PVC):** Mapping the SQLite database to a cloud disk.

### Deliverables
- [ ] Kubernetes manifests (YAML) for your Enterprise Suite.
- [ ] A local cluster (K3s/Minikube) running 8 namespaces with your app in each.

---

## 📊 Phase 5: Cloud, IaC & Observability (Weeks 10–12)
**Goal:** Professional SRE operations in a cloud environment.

### Key Concepts
* **Infrastructure as Code (IaC):** Using Terraform to provision a GKE or EKS cluster.
* **Observability Stack:** Installing Prometheus (metrics) and Grafana (dashboards).
* **Traffic Simulation:** Using `k6` scripts to mimic hundreds of users hitting your APIs.
* **SRE Discipline:** Setting Alerting rules for when a service becomes too slow.

### Deliverables
- [ ] Terraform files to create your cloud cluster.
- [ ] A Grafana Dashboard showing "Requests Per Second" across all 8 services.
- [ ] A "Chaos Report" documenting what happens when you manually kill a service pod.

---

## 🛠️ Recommended Tech Stack (2026 Standards)
* **Language:** Go 1.2x (Echo Framework)
* **Frontend:** Vue 3 (Vite)
* **Database:** SQLite
* **Infrastructure:** Terraform
* **CI/CD:** GitHub Actions + ArgoCD (GitOps)
* **Metrics:** Prometheus / Grafana / OpenTelemetry