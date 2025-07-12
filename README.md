# 🏭 SCADA Backend System

Dự án backend hệ thống **SCADA** quản lý thiết bị công nghiệp, cảnh báo trạng thái, và giao tiếp thời gian thực bằng **WebSocket**.

---

## 📁 Cấu trúc thư mục

```
SCADA/
├── cmd/api/                 # Entry point chính của backend
├── config/                  # Load biến môi trường, khởi tạo DB
├── internal/
│   ├── model/               # Định nghĩa các struct (Entity DB)
│   ├── repository/          # Tầng truy vấn DB (CRUD)
│   ├── service/             # Tầng xử lý logic nghiệp vụ
│   ├── http/
│   │   ├── handler/         # Xử lý request HTTP
│   │   ├── middleware/      # Middleware (auth, ...)
│   │   ├── response/        # Định nghĩa response
│   │   └── routes/          # Routing
│   └── websocket/           # Quản lý WebSocket
├── migrations/              # Migration DB
├── utils/                   # Tiện ích chung
├── docs/                    # Swagger docs (auto-gen)
├── Dockerfile               # Docker build
├── docker-compose.yml       # Docker Compose
├── go.mod / go.sum          # Quản lý package
├── .env                     # Thông tin kết nối DB, JWT
├── README.md                # 📄 YOU ARE HERE
```

---

## 🚀 Hướng dẫn chạy nhanh

1. **Cài đặt Go, Docker, PostgreSQL** (hoặc dùng docker-compose)
2. **Tạo file `.env`** (xem ENVIRONMENT_SETUP.md)
3. **Cài dependencies:**
   ```bash
   go mod download
   ```
4. **Tạo Swagger docs:**
   ```bash
   swag init -g cmd/api/main.go
   ```
5. **Chạy bằng Docker Compose:**
   ```bash
   docker-compose up --build
   ```
   Hoặc chạy local:
   ```bash
   go run ./cmd/api
   ```
6. **Truy cập Swagger UI:**
   - [http://localhost:5000/swagger/index.html](http://localhost:5000/swagger/index.html)

---

## 🔐 Hướng dẫn sử dụng Swagger với Token

1. **Đăng ký tài khoản:**
   - `POST /auth/register` (không cần token)
2. **Đăng nhập:**
   - `POST /auth/login` (không cần token)
   - Copy giá trị `token` trả về
3. **Authorize trên Swagger:**
   - Click nút "Authorize" (🔒)
   - Nhập: `Bearer <token>`
   - Click "Authorize" → "Close"
4. **Test các API cần token:**
   - Tất cả API `/api/*` đều yêu cầu JWT Bearer Token

> **Lưu ý:** Nếu vẫn bị lỗi 401, hãy kiểm tra lại:
> - Đã nhập đúng `Bearer <token>` chưa?
> - Token còn hạn không?
> - Đã build lại project và tạo lại Swagger docs chưa?
> - Đã bật CORS middleware chưa?

---

## 🧩 Các tính năng đã có

- [x] Đăng ký/Đăng nhập, sinh JWT token
- [x] CRUD thiết bị (equipment), loại thiết bị, phòng ban, nhân sự, nhà máy
- [x] Import thiết bị từ file CSV (UTF-8)
- [x] Download template CSV
- [x] Swagger UI với JWT Bearer Token
- [x] Middleware bảo vệ tất cả API `/api/*`
- [x] WebSocket real-time
- [x] Docker/Docker Compose support
- [x] Migration tự động
- [x] Hướng dẫn chi tiết (CSV_IMPORT_GUIDE.md, SWAGGER_AUTH_GUIDE.md, ENVIRONMENT_SETUP.md)

---

## 🚧 Feature đề xuất cho tương lai

- [ ] RBAC (Role-Based Access Control) - phân quyền chi tiết
- [ ] Lịch sử thao tác (audit log)
- [ ] Notification real-time qua WebSocket
- [ ] API quản lý cảnh báo (alert)
- [ ] API thống kê, dashboard
- [ ] API upload/download file đính kèm
- [ ] Đa ngôn ngữ (i18n)
- [ ] Unit test & integration test
- [ ] CI/CD pipeline
- [ ] Health check endpoint
- [ ] API rate limiting
- [ ] Export dữ liệu ra Excel/PDF

---

## 📝 Tài liệu tham khảo

- [CSV_IMPORT_GUIDE.md](./CSV_IMPORT_GUIDE.md) - Hướng dẫn import CSV
- [SWAGGER_AUTH_GUIDE.md](./SWAGGER_AUTH_GUIDE.md) - Hướng dẫn sử dụng Swagger với token
- [ENVIRONMENT_SETUP.md](./ENVIRONMENT_SETUP.md) - Hướng dẫn cấu hình môi trường
- [IMPLEMENTATION_SUMMARY.md](./IMPLEMENTATION_SUMMARY.md) - Tóm tắt các thay đổi

---

## 🐛 Troubleshooting

- **Lỗi 401 khi test API trên Swagger:**
  - Đảm bảo đã nhập đúng `Bearer <token>` khi Authorize
  - Đảm bảo token còn hạn
  - Đảm bảo đã build lại project và tạo lại Swagger docs
  - Đảm bảo CORS middleware đã bật
- **Lỗi encoding khi import CSV:**
  - File phải là UTF-8, không phải ANSI/UTF-16
- **Lỗi foreign key khi import CSV:**
  - Các ID (department_id, equipment_type_id, ...) phải tồn tại trong DB

---

## 💡 Đóng góp & phát triển

- Fork, tạo PR hoặc mở issue để đóng góp ý tưởng, báo lỗi hoặc đề xuất tính năng mới!

