# 🎬 Movie Management System

A lightweight backend service for managing movies, showtimes, and reservations. Built with **Go**, **Fiber v3**, and **PostgreSQL**.

> ⚠️ **Note:** Fiber v3 is currently used, but it is **not recommended for production**.

---

## 📦 Tech Stack

* **Go** – Primary backend language
* **Fiber v3** – Web framework (lightweight and fast)
* **PostgreSQL** – Relational database
* **JWT** – Authentication
* **Regex-based Routing** – For cleaner and more secure path parameters
* **Multipart File Uploads** – For handling movie poster uploads

---

## ✅ Features

* 🔐 **JWT-based Authentication**

  * User registration and login
  * Admin-only protected routes

* 🎥 **Movie Management**

  * CRUD operations for movies
  * Poster uploads (`/poster/:filename`)

* 🕒 **Showtime Management**

  * Create, update, delete, and list showtimes

* 🎟️ **Reservation System**

  * Book, cancel, and list reservations
  * Get reservations for the authenticated user

* 📁 **File Uploads**

  * Image uploads for movie posters with filename validation

* 🧪 **Regex-based Parameters**

  * Secure route matching (e.g., only allow valid image extensions, numeric IDs)

* ✅ **Request Payload Validation**

  * Centralized and consistent request data validation

* 🚨 **Centralized Error Handling**

  * Unified place to manage and customize error responses

---

## 🚀 API Structure

### Public Routes

| Method | Path                | Description               |
| ------ | ------------------- | ------------------------- |
| GET    | `/`                 | Root health check         |
| POST   | `/auth/login`       | User login                |
| POST   | `/auth/register`    | User registration         |
| GET    | `/poster/:filename` | Get uploaded poster image |

### Protected Routes (`/api`)

* **Movies**

  * `GET /movies` – Get all movies
  * `GET /movies/:id` – Get movie by ID

* **Showtimes**

  * `GET /showtimes` – Get all showtimes
  * `GET /showtimes/:id` – Get showtime by ID

* **Reservations**

  * `GET /reservations` – List all reservations (admin)
  * `GET /user/reservations` – Get current user's reservations
  * `POST /reservations` – Create a new reservation
  * `DELETE /reservations/:id` – Cancel a reservation

### Admin-Only Routes

* **Movies**

  * `POST /movies` – Create new movie
  * `PUT /movies/:id` – Update movie
  * `DELETE /movies/:id` – Delete movie

* **Showtimes**

  * `POST /showtimes` – Create new showtime
  * `PUT /showtimes/:id` – Update showtime
  * `DELETE /showtimes/:id` – Delete showtime

---

## 📂 Project Structure (Simplified)

```
movie-management-system/
├── handler/           # HTTP handlers for routes
├── middleware/        # JWT auth, Admin checks, etc.
├── repository/        # DB interaction layer
├── uploads/poster/    # Uploaded poster images
├── server/            # Server setup and route registration
└── main.go            # App entry point
```

---

## 🛠️ Getting Started

```bash
# clone repo
git clone https://github.com/Abhishek2010dev/movie-management-system.git
cd movie-management-system

# setup .env with DB and JWT_SECRET

# run the server
go run main.go
```

---

## 📌 Notes

* Ensure PostgreSQL is running and accessible
* Use tools like **Postman** for testing the API
* Use multipart form uploads for uploading poster images

---


## 🧠 Inspiration

This project is based on the [Movie Reservation System project idea from roadmap.sh](https://roadmap.sh/projects/movie-reservation-system), a part of their project-based learning path to solidify backend development skills.

---

## 📃 License

MIT
