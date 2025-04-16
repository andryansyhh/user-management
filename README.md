# user-management

# User Management Backend (gRPC + Go)

Sistem backend untuk autentikasi dan manajemen user menggunakan **gRPC**, **Golang**, **PostgreSQL**, dan **Redis**, dibangun dengan prinsip **Clean Architecture**.

---

## 🚀 Setup Cepat

### 1. Clone & install dependency

git clone https://github.com/yourname/user-management.git
cd user-management
go mod tidy

### 2. Setup Env
GRPC_PORT=:50051

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=tablelink

REDIS_ADDR=localhost:6379
REDIS_PASS=

### 3. Run Postgres & redis via docker (optional)
docker run --name postgres -e POSTGRES_PASSWORD=postgres -p 5432:5432 -d postgres
docker run --name redis -p 6379:6379 -d redis

### 4. Database Migration
brew install golang-migrate
make migrate-up

