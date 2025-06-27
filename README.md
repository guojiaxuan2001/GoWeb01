# Kelly's Web APP 01: Full-Stack Blog System

A full-stack blogging platform built with **Go (Gin)** and **Vue 3 (TypeScript)**. Features include user authentication (JWT), article posting, likes, exchange rates, Redis caching, and MySQL persistence. Fully containerized with Docker and ready for orchestration.

---

## 🌐 Docker Images

- **Frontend Image**: [guojiaxuan2001/web01-frontend](https://hub.docker.com/repository/docker/guojiaxuan2001/web01-frontend/general)
- **Backend Image**: [guojiaxuan2001/web01-backend](https://hub.docker.com/repository/docker/guojiaxuan2001/web01-backend/general)

---

## 🧱 Tech Stack

| Layer      | Technology                |
|------------|---------------------------|
| Frontend   | Vue 3, TypeScript, Vite   |
| Backend    | Go, Gin, Gorm, JWT        |
| Database   | MySQL                     |
| Cache      | Redis                     |
| Container  | Docker, Kubernetes (WIP)  |

---

## 🖥️ Local Development

### Frontend

```bash
cd frontend
npm install
npm run dev
```

### Backend

```bash
cd backend
go run .
```

Make sure MySQL and Redis are running locally or via Docker.

⸻

Run with Docker

Start MySQL and Redis (Docker Compose recommended)

docker-compose up -d redis mysql

Frontend
```bash
docker pull guojiaxuan2001/web01-frontend
docker run -d -p 80:80 guojiaxuan2001/web01-frontend
```

Backend
```bash
docker pull guojiaxuan2001/web01-backend
docker run -d -p 8080:8080 \
  --env-file .env \
  guojiaxuan2001/web01-backend
```
