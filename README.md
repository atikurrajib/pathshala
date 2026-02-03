<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Pathshala Documentation</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/github-markdown-css/5.2.0/github-markdown.min.css">
    <style>
        body {
            box-sizing: border-box;
            min-width: 200px;
            max-width: 980px;
            margin: 0 auto;
            padding: 45px;
            background-color: #f6f8fa;
        }
        .markdown-body {
            background-color: #ffffff;
            padding: 45px;
            border: 1px solid #d0d7de;
            border-radius: 6px;
            box-shadow: 0 4px 12px rgba(0,0,0,0.05);
        }
        @media (max-width: 767px) {
            body { padding: 15px; }
            .markdown-body { padding: 20px; }
        }
        /* Fix for center alignment in HTML tags */
        .markdown-body h1[align="center"] {
            border-bottom: none;
            margin-bottom: 0;
        }
    </style>
</head>
<body>

<article class="markdown-body">
    
    <h1 align="center">🏫 Welcome to Pathshala!</h1>
    <p align="center">
        <strong>Pathshala</strong> is a robust, full-stack Student Management System (SMS) designed to digitize and streamline academic administration. Built with a focus on high performance and security, it provides a centralized platform for managing student records, academic progress, and institutional data.
    </p>

    <hr>

    <h2>💻 Project Overview</h2>
    <p>Through <strong>Pathshala</strong>, we have developed a scalable solution for educational institutions to manage their day-to-day operations. The system ensures data integrity and provides real-time insights for administrators and teachers.</p>
    
    <h3>Key Features:</h3>
    <ul>
        <li><strong>Student Lifecycle Management:</strong> Comprehensive records from admission to graduation.</li>
        <li><strong>Academic Tracking:</strong> Manage classes, sections, subjects, and student attendance.</li>
        <li><strong>Result & Grading System:</strong> Automated grade calculation and report card generation.</li>
        <li><strong>Role-Based Access Control (RBAC):</strong> Distinct portals and permissions for Admins, Teachers, and Students.</li>
        <li><strong>Financial Management:</strong> Tracking student fees, payments, and financial history.</li>
    </ul>

    <hr>

    <h2>✅ Technology Stack</h2>
    
    <h3>Frontend (Server-Side Templating)</h3>
    <ul>
        <li><strong>Templating Engine:</strong> Go <code>html/template</code> (Standard Library)</li>
        <li><strong>Architecture:</strong> Modular Layouts (Header, Sidebar, Footer partials).</li>
        <li><strong>Styling:</strong> Custom CSS / Tailwind CSS.</li>
        <li><strong>Client-Side:</strong> Vanilla JS &amp; <strong>Chart.js</strong> for administrative data visualization.</li>
    </ul>

    <h3>Backend (Server Side)</h3>
    <ul>
        <li><strong>Language:</strong> Golang (Go)</li>
        <li><strong>Architecture:</strong> <strong>Repository Pattern</strong> for scalability and testing.</li>
        <li><strong>Database:</strong> PostgreSQL (Handled via internal/driver and dbrepo).</li>
        <li><strong>Security:</strong> Session-based Auth, Bcrypt password hashing, and CSRF protection.</li>
    </ul>

    <h3>Infrastructure &amp; DevOps</h3>
    <ul>
        <li><strong>Containerization:</strong> Docker &amp; Docker Compose.</li>
        <li><strong>Database Versioning:</strong> SQL Migrations (stored in <code>/migrations</code>).</li>
        <li><strong>Automation:</strong> <strong>Makefile</strong> for streamlined development workflows.</li>
    </ul>

    <hr>

    <h2>📂 Project Structure</h2>
    <p>A clean and organized folder structure following the <strong>Repository Pattern</strong> for scalability and maintainability.</p>

<pre><code>pathshala/
├── cmd/web/            # Main entry point (Server, Routes, Middlewares)
├── internal/
│   ├── config/         # Application configuration
│   ├── driver/         # Database connection drivers (Postgres)
│   ├── forms/          # Form validation &amp; error handling logic
│   ├── handlers/       # Request handlers &amp; Business logic
│   ├── helpers/        # Reusable helper functions
│   ├── models/         # Database structs and schemas
│   ├── repository/     # Data persistence layer (DBRepo implementation)
├── ui/                 # Frontend assets
│   ├── html/           # Go HTML templates (Layouts, Pages, Partials)
│   └── static/         # Assets (CSS, JS, Images, Charts)
├── migrations/         # SQL migration files for DB schema versioning
├── Dockerfile          # Docker image configuration
├── docker-compose.yml  # Multi-container setup (Go + Postgres)
└── .github/workflows/  # CI/CD pipeline configuration</code></pre>

    <hr>

    <h2>🏗️ System Architecture</h2>
    <p>Pathshala is built as a <strong>Structured Monolith</strong> using <strong>Server-Side Rendering (SSR)</strong>. The architecture follows the <strong>Repository Pattern</strong>, which decouples the business logic from the data storage layer, making the system highly maintainable and easy to test.</p>

    

    <h3>🔍 Technical Highlights:</h3>
    <ol>
        <li><strong>Repository Pattern:</strong> Decouples business logic from SQL queries, ensuring the codebase is scalable and database-agnostic.</li>
        <li><strong>Server-Side Templating:</strong> Utilizes Go’s <code>html/template</code> for secure, high-performance rendering of the student dashboard and admin panels.</li>
        <li><strong>Database Design:</strong> Advanced relational schema designed for data normalization, managed through version-controlled SQL migrations.</li>
        <li><strong>Security First:</strong> Implementation of session-based authentication, Bcrypt password hashing, and custom middleware for granular Role-Based Access Control (RBAC).</li>
        <li><strong>Concurrency:</strong> Utilizes Go’s Goroutines for efficient background processing and non-blocking report generation.</li>
    </ol>

    <hr>

    <h2>🚀 Future Roadmap</h2>
    <ul>
        <li>☐ <strong>Parent Portal:</strong> Dedicated mobile-responsive view for parents to track child progress.</li>
        <li>☐ <strong>Automated Notifications:</strong> Email/SMS integration for attendance and fee alerts.</li>
        <li>☐ <strong>Exam Management:</strong> Modules for online MCQ exams and scheduling.</li>
    </ul>

    <hr>

    <h2>🛠️ Getting Started</h2>
    
    <h3>Prerequisites</h3>
    <ul>
        <li>Go 1.2x+</li>
        <li>PostgreSQL 14+</li>
        <li>Docker (Optional)</li>
    </ul>

    <h3>Installation</h3>
    <ol>
        <li><strong>Clone the repository:</strong>
            <pre><code>git clone https://github.com/atikurrajib/pathshala.git
cd pathshala</code></pre>
        </li>
        <li><strong>Setup Environment:</strong>
            <pre><code>cp .env.example .env # Configure your database credentials</code></pre>
        </li>
        <li><strong>Run migrations:</strong> Apply the SQL files in <code>/migrations</code> to your database.</li>
        <li><strong>Run the application:</strong>
            <pre><code>go run cmd/web/*.go</code></pre>
        </li>
    </ol>

</article>

</body>
</html>
