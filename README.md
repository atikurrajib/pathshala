<h1 align = "center"> Pathshala </h1>

## 📖 Project Overview

Through **Pathshala**, we have developed a scalable web solution to to automate courses, student attendances, performances and some administrative workflows. The system ensures data integrity and provides real-time insights for administrators and teachers.

### 🌟 Key Features:
* **Student Lifecycle Management:** Comprehensive records from admission to graduation.
* **Academic Tracking:** Manage classes, sections, subjects, and student attendance.
* **Result & Grading System:** Automated grade calculation and report card generation.
* **Role-Based Access Control (RBAC):** Distinct portals and permissions for Admins, Teachers, and Students.
* **Financial Management:** Tracking student fees, payments, and financial history.

---

## ⚙️ System Architecture

Pathshala is built as a Structured Monolith using Server-Side Rendering (SSR).

Data flows strictly from the Router → Handler → Repository → Database, ensuring a clean separation of concerns.

### Technical Highlights:

1. **Repository Pattern:** Decouples business logic from SQL queries, ensuring the codebase is scalable and database-agnostic.
2. **Server-Side Templating:** Utilizes Go’s `html/template` for secure, high-performance rendering of the student dashboard and admin panels.
3. **Database Design:** Advanced relational schema designed for data normalization, managed through version-controlled SQL migrations.
4. **Security First:** Implementation of session-based authentication, Bcrypt password hashing, and custom middleware for granular Role-Based Access Control (RBAC).
5. **Concurrency:** Utilizes Go’s Goroutines for efficient background processing and non-blocking report generation.

---

## 🔮 Future Scope

* [ ] **Parent Portal:** Dedicated mobile-responsive view for parents to track child progress.
* [ ] **Automated Notifications:** Email/SMS integration for attendance and fee alerts.
* [ ] **Exam Management:** Modules for online MCQ exams and scheduling.

---

## 🚀 Getting Started

### Prerequisites

* [Go 1.21+](https://go.dev/dl/) (For manual setup)
* [PostgreSQL 14+](https://www.postgresql.org/) (For manual setup)
* [Docker](https://www.docker.com/) (Recommended for easy setup)

### Installation

You can run Pathshala using **Docker** (Recommended) or manually on your local machine.

#### Option 1: Using Docker 🐳
*No need to install Go or PostgreSQL manually on your machine.*

1. **Clone the repository:**
```bash
git clone https://github.com/atikurrajib/pathshala.git
cd pathshala
```
2. **Run the application:**
```bash
docker-compose up --build
```
3. **Access the app:**
   Open [http://localhost:8080](http://localhost:8080) in your browser.

---

#### Option 2: Manual Setup 🛠️
*Requires Go and PostgreSQL installed on your machine.*
   
1.  **Clone the repository:**
    ```bash
    git clone https://github.com/atikurrajib/pathshala.git
    cd pathshala
    ```
2.  **Setup Environment:**
    ```bash
    cp .env.example .env
    # Update .env with your local DB credentials (DB_HOST=localhost)
    ```
3.  **Run migrations:**
    Apply the SQL files in `/migrations` to your database to create the tables.

4.  **Run the application:**
    ```bash
    go run cmd/server/main.go
    ```
    *Visit [http://localhost:8080](http://localhost:8080) in your browser.*

---

<div align="center">
  Thanks for checking out <strong>Pathshala</strong>! Happy Coding! 🚀
</div>
