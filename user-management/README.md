# user-management

# User Management gRPC Backend

Sistem backend sederhana untuk manajemen user & autentikasi menggunakan **Golang**, **gRPC**, **PostgreSQL**, dan **Redis**, dengan pendekatan **Clean Architecture**.


## Setup

### 1. Clone & install dependency

```bash
git clone https://github.com/yourname/user-management.git
cd user-management
go mod tidy

GRPC_PORT=:50051

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=tablelink

REDIS_ADDR=localhost:6379
REDIS_PASS=
