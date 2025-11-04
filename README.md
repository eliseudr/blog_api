# **Backend Coding Challenge**

You are tasked with designing and implementing a RESTful API for managing a simple blogging platform. The core functionality of this platform includes managing blog posts and their associated comments.

**Requirements:**

- **Data Models:**
  - Create two data models: BlogPost and Comment. A BlogPost has a title and content, and each BlogPost can have multiple Comment objects associated with it.

- **API Endpoints:**
  - Implement the following API endpoints:
    - **GET** /api/posts: This endpoint should return a list of all blog posts, including their titles and the number of comments associated with each post.
    - **GET** /api/posts/{id}: Retrieve a specific blog post by its ID, including its title, content, and a list of associated comments.
    - **POST** /api/posts: Create a new blog post.
    - **POST** /api/posts/{id}/comments: Add a new comment to a specific blog post.

---

## 🚀 Getting Started

### Prerequisites

- **Go** 1.25.3 or higher ([Download](https://go.dev/dl/))
- **MySQL** 5.7+ or MySQL 8.0+ ([Download](https://dev.mysql.com/downloads/))
- **Git** (optional, for cloning the repository)

### Installation

1. **Clone the repository** (or download the source code):

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Set up environment variables**:
   
   Create a `.env` file in the root directory with the following variables:
   ```env
   DB_HOST=localhost
   DB_PORT=3306
   DB_USER=root
   DB_PASSWORD=your_password
   DB_NAME=blog_challenge
   SERVER_MODE=DEV
   ```
   
   **Environment Variables:**
   - `DB_HOST`: MySQL server host (default: `localhost`)
   - `DB_PORT`: MySQL server port (default: `3306`)
   - `DB_USER`: MySQL username
   - `DB_PASSWORD`: MySQL password
   - `DB_NAME`: Database name (will be created automatically if it doesn't exist)
   - `SERVER_MODE`: Set to `DEV` to enable SQL query logging, or `PRODUCTION` to disable it

### Running the Application

1. **Start MySQL server**:
   Make sure your MySQL server is running and accessible with the credentials specified in your `.env` file.

2. **Run the application**:
   ```bash
   go run main.go
   ```
   
   Or build and run:
   ```bash
   go build .
   ./blog_api.exe    # Windows
   ./blog_api        # Linux/Mac
   ```

3. **Verify the server is running**:
   The server will start on `http://localhost:8080` by default. You should see output like:
   ```
   Blog API is running...
   Connected to MySQL server
   Database created/verified: blog_challenge
   Connected to database: blog_challenge
   Database migration completed
   2025/11/04 12:00:00 Server is running on port :8080
   ```

### Testing the API

#### Using Insomnia Collection

1. **Import the collection**:
   - Open Insomnia
   - Import **From File** and choose `Insomnia_2025-11-04.yaml`
   - The collection will be imported with all endpoints pre-configured

2. **Test the endpoints**:
   - **GET /api/posts**: Retrieve all blog posts with comment counts
   - **GET /api/posts/{id}**: Retrieve a specific post with its comments
   - **POST /api/posts**: Create a new blog post
   - **POST /api/posts/{id}/comments**: Add a comment to a post

### Project Structure

```
Blog API/
├── controller/      # HTTP handlers (business logic)
├── database/        # Database connection and initialization
├── middleware/      # HTTP middleware (logging, error handling)
├── models/          # Domain entities (BlogPost, Comment)
├── repository/      # Data access layer
├── response/        # Standardized API response helpers
├── router/          # Route configuration
├── server/          # HTTP server setup
└── main.go          # Application entry point
```

### Architecture

This project follows a **Layered Architecture** pattern with:
- **Repository Pattern** for data access abstraction
- **MVC** principles adapted for REST APIs
- **Clean Architecture** principles for separation of concerns

### Troubleshooting

- **Database connection errors**: Verify MySQL is running and credentials in `.env` are correct
- **Port already in use**: Change the port in `main.go` or stop the process using port 8080
- **Migration errors**: Ensure the MySQL user has CREATE DATABASE permissions


