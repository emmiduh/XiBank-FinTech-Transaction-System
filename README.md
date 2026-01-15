# XiBank FinTech Transaction System

![Build Status](https://img.shields.io/badge/build-passing-brightgreen) ![Docker](https://img.shields.io/badge/container-docker-blue) ![Kubernetes Ready](https://img.shields.io/badge/orchestration-k8s-326ce5)

A cloud-native, polyglot microservices architecture designed to simulate a high-reliability financial backend. This project demonstrates the integration of **Node.js**, **Go**, and **Python** services communicating to handle authentication, ledger transactions, and fraud detection.

**Role:** DevOps / Backend Architecture  
**Focus:** Orchestration, Containerization, and Service Interoperability.

---

## 🏗 Architecture

The system is composed of three decoupled microservices and persistent storage:

| Service | Language | Port | Responsibility |
| :--- | :--- | :--- | :--- |
| **Auth Service** | Node.js (Express) | `3000` | Handles User JWT issuance and validation. |
| **Ledger Service** | Go (Golang) | `4000` | High-performance transaction processing engine. |
| **Fraud Service** | Python (Flask) | `5000` | Data analysis service to flag suspicious transactions (> $5000). |
| **Database** | PostgreSQL | `5432` | Relational storage for transaction logs. |
| **Cache** | Redis | `6379` | Session caching and fast retrieval. |

### Directory Structure

```text
polyglot-fintech/
├── service-auth-node/     # Node.js Auth Microservice
├── service-ledger-go/     # Go Transaction Microservice
├── service-fraud-python/  # Python Fraud Detection Microservice
├── docker-compose.yml     # Local Orchestration
├──  kubernetes/            # Kubernetes Manifests
├──  terraform/             # Infrastructure as Code (IaC)
└── README.md
```
You can include your directory structure and the screenshot in README.md using Markdown like this:

# Polyglot Fintech Project

## Directory Structure


polyglot-fintech/
├── service-auth-node/ # Node.js Auth Microservice
├── service-ledger-go/ # Go Transaction Microservice
├── service-fraud-python/ # Python Fraud Detection Microservice
├── docker-compose.yml # Local Orchestration
├── kubernetes/ # Kubernetes Manifests
├── terraform/ # Infrastructure as Code (IaC)
└── README.md


## Screenshot
![Project Screenshot]("https://github.com/user-attachments/assets/c1c760db-0f39-4c43-8ea1-0813bc4d7628")

