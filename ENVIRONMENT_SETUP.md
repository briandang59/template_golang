# 🔧 Hướng dẫn cấu hình môi trường

## 📋 Tạo file .env

Tạo file `.env` trong thư mục gốc của dự án với nội dung sau:

```env
# Database Configuration
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=123456
DB_NAME=scada_db
DB_PORT=5432

# Server Configuration
PORT=5000

# JWT Configuration
JWT_SECRET=your-super-secret-jwt-key-here-change-in-production
```

## 🔐 Cấu hình JWT Secret

**QUAN TRỌNG**: Thay đổi `JWT_SECRET` thành một chuỗi ngẫu nhiên mạnh trong production:

```env
JWT_SECRET=your-very-long-and-random-secret-key-here
```

### Tạo JWT Secret mạnh:
```bash
# Sử dụng OpenSSL
openssl rand -base64 32

# Hoặc sử dụng Python
python -c "import secrets; print(secrets.token_urlsafe(32))"
```

## 🗄️ Cấu hình Database

### PostgreSQL
1. Cài đặt PostgreSQL
2. Tạo database:
```sql
CREATE DATABASE scada_db;
CREATE USER scada_user WITH PASSWORD 'your_password';
GRANT ALL PRIVILEGES ON DATABASE scada_db TO scada_user;
```

### Docker (Khuyến nghị)
```bash
# Chạy PostgreSQL với Docker
docker run --name scada-postgres \
  -e POSTGRES_DB=scada_db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=123456 \
  -p 5432:5432 \
  -d postgres:16
```

## 🚀 Khởi chạy ứng dụng

### 1. Cài đặt dependencies
```bash
go mod download
```

### 2. Tạo Swagger docs
```bash
swag init -g cmd/api/main.go
```

### 3. Build và chạy
```bash
go build ./cmd/api
./api
```

### 4. Hoặc chạy trực tiếp
```bash
go run ./cmd/api
```

## 🔍 Kiểm tra

### 1. API Health Check
```bash
curl http://localhost:5000/api/health
```

### 2. Swagger UI
```
http://localhost:5000/swagger/index.html
```

### 3. WebSocket
```
ws://localhost:5000/ws
```

## 🐳 Docker Compose

Sử dụng file `docker-compose.yml` có sẵn:

```bash
docker-compose up -d
```

## ⚠️ Lưu ý bảo mật

### Development
- Sử dụng JWT secret đơn giản cho development
- Database có thể không cần SSL

### Production
- **BẮT BUỘC** thay đổi JWT_SECRET
- Sử dụng HTTPS
- Cấu hình database SSL
- Thiết lập firewall
- Sử dụng environment variables thay vì hardcode

## 🔧 Troubleshooting

### Lỗi kết nối database
- Kiểm tra PostgreSQL đã chạy chưa
- Kiểm tra thông tin kết nối trong .env
- Kiểm tra firewall

### Lỗi JWT
- Đảm bảo JWT_SECRET đã được set
- Kiểm tra format của JWT token

### Lỗi Swagger
- Chạy lại `swag init -g cmd/api/main.go`
- Kiểm tra import docs trong main.go 