# 🎓 Welcome to Pathshala (Full-Stack LMS)

✨ **Pathshala** is a comprehensive, full-stack Learning Management System. It combines a high-performance **Golang** backend with a responsive **React** frontend to provide a seamless educational experience for both teachers and students. ✨

---

## 💻 Project Overview

Pathshala is designed to bridge the gap between complex academic management and user-friendly interfaces. It handles everything from student registration and course material hosting to secure payment processing and progress tracking.

### Core Modules:
* **Student Dashboard:** Personalized view of enrolled courses, progress bars, and upcoming assignments.
* **Instructor Studio:** Tools for educators to upload content, manage student lists, and track course performance.
* **Admin Panel:** Global control over user roles, course approvals, and system analytics.
* **Secure API:** A robust RESTful API layer ensuring data integrity and fast response times.

---

## ✅ Technology Stack

### Frontend (Client-Side)
* **Framework:** React.js / Vite
* **State Management:** Redux Toolkit or React Query
* **Styling:** Tailwind CSS (for a modern, mobile-first UI)
* **API Client:** Axios for seamless communication with the Go backend

### Backend (Server-Side)
* **Language:** Golang
* **Framework:** Gin Gonic
* **Database:** PostgreSQL (with SQLC for type-safe queries)
* **Security:** PASETO/JWT Authentication & Bcrypt password hashing

### Infrastructure & DevOps
* **Containerization:** Docker & Docker Compose
* **Orchestration:** Kubernetes (AWS EKS)
* **CI/CD:** GitHub Actions for automated testing and deployment
* **Documentation:** Swagger UI for API testing

---

## 🏗️ System Architecture

The application follows a decoupled architecture where the React frontend communicates with the Go backend via a JSON REST API. This allows for independent scaling of the frontend and backend services.



### 🔍 Technical Highlights:
1.  **Database Efficiency:** Using PostgreSQL with optimized indexing for fast course searching and student record retrieval.
2.  **State Consistency:** Implementing robust error handling on the frontend to manage API states (Loading, Success, Error) gracefully.
3.  **Scalable Storage:** Integration with AWS S3 for hosting heavy course assets like videos and PDFs.
4.  **Middleware Security:** Custom Go middleware for CORS handling, logging, and role-based route protection.

---

## 🚀 Future Roadmap
* [ ] Integration of a real-time chat system for student-teacher interaction.
* [ ] Live video streaming for virtual classrooms.
* [ ] AI-driven course recommendations based on student performance.

---

## 🛠️ Getting Started

### Prerequisites
* Go 1.2x
* Node.js & npm/yarn
* Docker & PostgreSQL

### Installation
1. **Clone the repo:**
   ```bash
   git clone [https://github.com/your-username/pathshala.git](https://github.com/your-username/pathshala.git)
