# Match Maker

**Match Maker** is a full-stack web application designed to manage and interact with football match events. It supports role-based functionality for guests, EFA staff, and master administrators. This system allows users to browse, reserve, and manage matches, stadiums, and teams through a clean and responsive interface.

---

## 🧩 Project Structure

```
.
├── backend         # Go-based backend with routing, models, auth, and DB integration
├── frontend        # Next.js frontend application with TailwindCSS
└── README.md       # Project documentation
```

---

## 🌐 Live Features

- **Role-based access** (Guest, EFA, Master)
- **Match reservation system**
- **Match, stadium, team, and user management**
- **Socket-based real-time updates**
- **JWT authentication and session management**
- **Swagger documentation for API testing**

---

## 🛠 Tech Stack

### Backend
- **Language**: Go
- **Framework**: net/http
- **Database**: Embedded (`db.db`)
- **Docs**: Swagger (JSON/YAML)
- **Testing**: HTTP files for API test cases

### Frontend
- **Framework**: Next.js (App Router)
- **Language**: TypeScript
- **Styling**: TailwindCSS
- **Auth Management**: React Context API

---

## 📂 Backend Directory Overview

```
backend/
├── api-test/         # HTTP files for manual API testing
├── db/               # Database connection logic
├── docs/             # Swagger files for API documentation
├── middlewares/      # Auth middleware for different roles
├── models/           # Go structs for DB and API
├── routes/           # All HTTP route handlers
├── utils/            # Utility functions (JWT, password hashing)
├── main.go           # Application entry point
```

---

## 💻 Frontend Directory Overview

```
frontend/match-maker/
├── src/
│   ├── app/          # Page and layout routing for each role
│   ├── components/   # Shared UI components like NavBar, Modals
│   ├── context/      # Global state providers (Auth, Error)
│   └── utils/        # API utility for backend communication
```

---

## 🧪 Running the Project Locally

### Prerequisites
- Go >= 1.18
- Node.js >= 18

### Backend

```bash
cd backend
go run main.go
```

### Frontend

```bash
cd frontend/match-maker
npm install
npm run dev
```

---

## 📘 API Documentation

Once the backend is running, access the Swagger UI:

```
http://localhost:8080/swagger/index.html
```

API test cases are located under `backend/api-test/`.

---

## 🧑‍💻 Contributing

1. Fork this repository.
2. Create a new branch: `git checkout -b feature/your-feature-name`
3. Make your changes and commit them: `git commit -m 'Add some feature'`
4. Push to the branch: `git push origin feature/your-feature-name`
5. Create a pull request.

---
