<div align="center">
# 🏫 Welcomne to Pathshala

✨ **Pathshala** is a robust, full-stack Student Management System (SMS) designed to digitize and streamline academic administration. Built with a focus on high performance and security, it provides a centralized platform for managing student records, academic progress, and institutional data. ✨
</div>
---

## 💻 Project Overview

Through **Pathshala**, we have developed a scalable solution for educational institutions to manage their day-to-day operations. The system ensures data integrity and provides real-time insights for administrators and teachers.

### Key Features:
* **Student Lifecycle Management:** Comprehensive records from admission to graduation.
* **Academic Tracking:** Manage classes, sections, subjects, and student attendance.
* **Result & Grading System:** Automated grade calculation and report card generation.
* **Role-Based Access Control (RBAC):** Distinct portals and permissions for Admins, Teachers, and Students.
* **Financial Management:** Tracking student fees, payments, and financial history.

---

## ✅ Technology Stack

### Frontend (User Interface)
* **Framework:** React.js (Vite)
* **Styling:** Tailwind CSS (for a clean, professional dashboard)
* **State Management:** React Query / Context API
* **Communication:** Axios for REST API integration

### Backend (Server Side)
* **Language:** Golang (Go)
* **Web Framework:** Gin Gonic
* **Database:** PostgreSQL (with SQLC for type-safe database operations)
* **Security:** PASETO & JWT for secure session management
* **Documentation:** Swagger UI for API discovery

### Infrastructure & DevOps
* **Containerization:** Docker & Docker Compose
* **Cloud:** AWS (EKS for orchestration, RDS for managed DB)
* **CI/CD:** GitHub Actions for automated building and testing

---

## 🏗️ System Architecture

Pathshala utilizes a modern **Decoupled Architecture**. The React frontend serves as a Single Page Application (SPA), communicating with the Golang backend via a high-speed RESTful API.

### 🔍 Technical Highlights:
1.  **Database Design:** Advanced schema designed with **DBML**, focusing on data normalization and referential integrity.
2.  **Concurrency:** Leveraging Go’s goroutines for efficient report generation and bulk data processing.
3.  **Security First:** Implementation of Bcrypt for password hashing and PASETO tokens to ensure top-tier authentication security.
4.  **Scalable Deployment:** Ready for production with Kubernetes (K8s) deployment scripts and environment-specific configurations.

---

## 🚀 Future Roadmap
* [ ] **Parent Portal:** Dedicated mobile-responsive view for parents to track child progress.
* [ ] **Automated Notifications:** Email/SMS integration for attendance and fee alerts.
* [ ] **Exam Management:** Modules for online MCQ exams and scheduling.

---

## 🛠️ Getting Started

### Prerequisites
* Go 1.2x+
* Node.js (v18+)
* PostgreSQL
* Docker (Optional)

### Installation
1. **Clone the repository:**
   ```bash
git clone [https://github.com/your-username/pathshala.git](https://github.com/your-username/pathshala.git)

2. **Backend Setup:**
   ```bash
cd pathshala-backend
cp app.env.example app.env # Configure your DB credentials
go run main.go

3. **Frontend Setup:**
   ```bash
cd pathshala-frontend
npm install
npm run dev
