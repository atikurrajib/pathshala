# 🏫 Welcome to Pathshala!

**Pathshala** is a robust, full-stack Student Management System designed to digitize and streamline academic administration. Built with a focus on high performance and security, it provides a centralized platform for managing student records, academic progress, and institutional data.

## Project Overview

Through **Pathshala**, we have developed a scalable solution for educational institutions to manage their day-to-day operations. The system ensures data integrity and provides real-time insights for administrators and teachers.

### Key Features:
* **Student Lifecycle Management:** Comprehensive records from admission to graduation.
* **Academic Tracking:** Manage classes, sections, subjects, and student attendance.
* **Result & Grading System:** Automated grade calculation and report card generation.
* **Role-Based Access Control (RBAC):** Distinct portals and permissions for Admins, Teachers, and Students.
* **Financial Management:** Tracking student fees, payments, and financial history.

## Project Structure

A clean and organized folder structure following the **Repository Pattern** for scalability and maintainability.

```bash
pathshala/
├── cmd/
│   └── server/
│       └── main.go        # Entry point
├── internal/
│   ├── config/            # Database ও Env configurations
│   ├── handlers/          # HTTP handlers (Controller)
│   ├── models/            # Database Models/Structs
│   └── repository/        # Database operations (SQL query)
├── web/
│   ├── template/          # Go HTML templates
│   └── static/            # CSS, JS, Images
├── go.mod
└── .env                   # Database credentials

```
---
## 🏗️ System Architecture

Pathshala is built as a **Structured Monolith** using **Server-Side Rendering (SSR)**. The architecture follows the **Repository Pattern**, which decouples the business logic from the data storage layer, making the system highly maintainable and easy to test.

### 🔍 Technical Highlights:

1. **Repository Pattern:** Decouples business logic from SQL queries, ensuring the codebase is scalable and database-agnostic.
2. **Server-Side Templating:** Utilizes Go’s `html/template` for secure, high-performance rendering of the student dashboard and admin panels.
3. **Database Design:** Advanced relational schema designed for data normalization, managed through version-controlled SQL migrations.
4. **Security First:** Implementation of session-based authentication, Bcrypt password hashing, and custom middleware for granular Role-Based Access Control (RBAC).
5. **Concurrency:** Utilizes Go’s Goroutines for efficient background processing and non-blocking report generation.

## Future Roadmap

* [ ] **Parent Portal:** Dedicated mobile-responsive view for parents to track child progress.
* [ ] **Automated Notifications:** Email/SMS integration for attendance and fee alerts.
* [ ] **Exam Management:** Modules for online MCQ exams and scheduling.

## Getting Started

### Prerequisites

* Go 1.2x+
* PostgreSQL 14+
* Docker (Optional)

### Installation

1. **Clone the repository:**
```bash
git clone [https://github.com/atikurrajib/pathshala.git](https://github.com/atikurrajib/pathshala.git)
cd pathshala
```
2. **Setup Environment:**
```bash
cp .env.example .env # Configure your database credentials
```
3. **Run migrations:** Apply the SQL files in `/migrations` to your database.
4. **Run the application:**
```bash
go run cmd/web/*.go
```
---
## 📬 Contact Me
If you have any questions, suggestions, or just want to connect, feel free to reach out!

* **LinkedIn:** [https://www.linkedin.com/in/atikurajib](https://www.linkedin.com/in/atikurajib)
* **GitHub:** [@atikurrajib](https://github.com/atikurrajib)
---
<div align="center">
  Thanks for checking out <strong>Pathshala</strong>! Happy Coding! 🚀
</div>
