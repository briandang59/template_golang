# 🔐 Hướng dẫn sử dụng Swagger với Authentication

## 🎯 Tổng quan

Swagger UI đã được cấu hình với JWT Bearer Token authentication. Tất cả các API trong nhóm `/api` đều yêu cầu xác thực.

## 📋 Cách sử dụng

### 1. Truy cập Swagger UI
```
http://localhost:5000/swagger/index.html
```

### 2. Đăng ký tài khoản (nếu chưa có)
- Tìm API `POST /auth/register`
- Click "Try it out"
- Nhập thông tin:
```json
{
  "username": "admin",
  "password": "password123"
}
```
- Click "Execute"

### 3. Đăng nhập để lấy token
- Tìm API `POST /auth/login`
- Click "Try it out"
- Nhập thông tin đăng nhập:
```json
{
  "username": "admin",
  "password": "password123"
}
```
- Click "Execute"
- Copy token từ response (không bao gồm dấu ngoặc kép)

### 4. Cấu hình Authentication trong Swagger
- Click nút **"Authorize"** (🔒) ở góc trên bên phải
- Trong trường "Value", nhập: `Bearer YOUR_TOKEN_HERE`
- Ví dụ: `Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...`
- Click "Authorize"
- Click "Close"

### 5. Sử dụng các API được bảo vệ
Bây giờ bạn có thể sử dụng tất cả các API trong nhóm Equipment:
- `GET /api/equipments` - Lấy danh sách thiết bị
- `POST /api/equipments` - Tạo thiết bị mới
- `POST /api/equipments/import` - Import CSV
- `GET /api/equipments/template` - Download template CSV
- `PATCH /api/equipments/{id}` - Cập nhật thiết bị
- `DELETE /api/equipments/{id}` - Xóa thiết bị

## 🔧 API Endpoints

### Authentication (Không cần token)
- `POST /auth/register` - Đăng ký tài khoản
- `POST /auth/login` - Đăng nhập

### Equipment (Cần token)
- `GET /api/equipments` - Lấy danh sách thiết bị
- `GET /api/equipments/template` - Download CSV template
- `POST /api/equipments` - Tạo thiết bị mới
- `POST /api/equipments/import` - Import CSV
- `PATCH /api/equipments/{id}` - Cập nhật thiết bị
- `DELETE /api/equipments/{id}` - Xóa thiết bị

### Các API khác (Cần token)
- Factory APIs
- Department APIs
- Equipment Type APIs
- Personnel APIs

## 📝 Ví dụ sử dụng

### 1. Import CSV với Authentication
```bash
# 1. Đăng nhập để lấy token
curl -X POST "http://localhost:5000/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password123"}'

# 2. Sử dụng token để import CSV
curl -X POST "http://localhost:5000/api/equipments/import" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE" \
  -F "file=@equipment_data.csv"
```

### 2. Lấy danh sách thiết bị
```bash
curl -X GET "http://localhost:5000/api/equipments?page=1&page_size=10" \
  -H "Authorization: Bearer YOUR_TOKEN_HERE"
```

## ⚠️ Lưu ý quan trọng

1. **Token Expiration**: JWT token có thời hạn, nếu hết hạn cần đăng nhập lại
2. **Authorization Header**: Luôn sử dụng format `Bearer TOKEN`
3. **CSRF Protection**: Swagger UI tự động xử lý CSRF protection
4. **Error Handling**: Nếu token không hợp lệ, API sẽ trả về 401 Unauthorized

## 🐛 Xử lý lỗi thường gặp

### 401 Unauthorized
- Token đã hết hạn → Đăng nhập lại
- Token không đúng format → Kiểm tra lại "Bearer TOKEN"
- Token không hợp lệ → Đăng nhập lại

### 403 Forbidden
- Tài khoản không có quyền truy cập API này

### 400 Bad Request
- Dữ liệu gửi lên không đúng format
- Thiếu các trường bắt buộc

## 🔒 Bảo mật

- Tất cả API trong `/api` đều được bảo vệ bởi JWT middleware
- Token được lưu trong memory của Swagger UI
- Không lưu token vào localStorage để tránh XSS
- Sử dụng HTTPS trong production

## 📱 Sử dụng với Postman/Insomnia

1. Đăng nhập để lấy token
2. Thêm header: `Authorization: Bearer YOUR_TOKEN`
3. Gọi các API bình thường

## 🚀 Production Deployment

Trong môi trường production:
- Sử dụng HTTPS
- Cấu hình CORS phù hợp
- Thiết lập JWT secret mạnh
- Cấu hình token expiration time
- Sử dụng refresh token nếu cần 