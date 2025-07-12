# 📁 Hướng dẫn Import CSV cho Equipment

## 🎯 Tổng quan

Tính năng import CSV cho phép bạn tạo nhiều thiết bị (equipment) cùng lúc bằng cách upload file CSV. Hệ thống sẽ xử lý file và trả về kết quả chi tiết về việc import.

## 📋 API Endpoints

### 1. Download CSV Template
```
GET /api/equipments/template
```
Tải về file CSV template với headers và dữ liệu mẫu.

### 2. Import CSV
```
POST /api/equipments/import
Content-Type: multipart/form-data
```
Upload file CSV để import equipment.

## 📊 Cấu trúc CSV

File CSV phải có các cột sau (theo thứ tự):

| Cột | Tên | Bắt buộc | Mô tả | Ví dụ |
|-----|-----|----------|-------|-------|
| 1 | name_en | ✅ | Tên thiết bị (tiếng Anh) | "Production Web Server 01" |
| 2 | name_zh | ❌ | Tên thiết bị (tiếng Trung) | "生产网络服务器01" |
| 3 | name_vn | ❌ | Tên thiết bị (tiếng Việt) | "Máy chủ Web Sản xuất 01" |
| 4 | code | ✅ | Mã thiết bị | "SERVER-PROD-WEB-001" |
| 5 | serial_number | ❌ | Số serial | "SN1234567890ABC" |
| 6 | model | ❌ | Model thiết bị | "Dell PowerEdge R740" |
| 7 | manufacturer | ❌ | Nhà sản xuất | "Dell" |
| 8 | location | ❌ | Vị trí | "Server Room A, Rack 2" |
| 9 | purchase_date | ❌ | Ngày mua (ISO format) | "2023-01-15T00:00:00Z" |
| 10 | warranty_end_date | ❌ | Ngày hết bảo hành | "2026-01-15T00:00:00Z" |
| 11 | installation_date | ❌ | Ngày lắp đặt | "2023-01-20T00:00:00Z" |
| 12 | status | ❌ | Trạng thái | "active" |
| 13 | ip_address | ❌ | Địa chỉ IP | "192.168.1.10" |
| 14 | mac_address | ❌ | Địa chỉ MAC | "00:1A:2B:3C:4D:5E" |
| 15 | operating_system | ❌ | Hệ điều hành | "Ubuntu Server 22.04" |
| 16 | description | ❌ | Mô tả | "Primary web server" |
| 17 | notes | ❌ | Ghi chú | "Redundant power supplies" |
| 18 | last_maintenance_date | ❌ | Ngày bảo trì cuối | "2024-12-01T00:00:00Z" |
| 19 | next_maintenance_date | ❌ | Ngày bảo trì tiếp theo | "2025-03-01T00:00:00Z" |
| 20 | department_id | ❌ | ID phòng ban | "1" |
| 21 | equipment_type_id | ❌ | ID loại thiết bị | "1" |
| 22 | responsible_user_id | ❌ | ID người phụ trách | "1" |
| 23 | assigned_user_id | ❌ | ID người được giao | "1" |

## 📅 Định dạng ngày tháng

Tất cả các trường ngày tháng phải sử dụng định dạng ISO 8601:
```
YYYY-MM-DDTHH:mm:ssZ
```

Ví dụ: `2023-01-15T00:00:00Z`

## 🔧 Cách sử dụng

### Bước 1: Tải template
```bash
curl -X GET "http://localhost:8080/api/equipments/template" \
  -H "Content-Type: application/json" \
  --output equipment_template.csv
```

### Bước 2: Chuẩn bị file CSV
1. Mở file `equipment_template.csv` đã tải
2. Xóa dòng dữ liệu mẫu (dòng 2)
3. Thêm dữ liệu thiết bị của bạn
4. Lưu file với encoding UTF-8

### Bước 3: Upload file
```bash
curl -X POST "http://localhost:8080/api/equipments/import" \
  -H "Content-Type: multipart/form-data" \
  -F "file=@your_equipment_data.csv"
```

## 📤 Response Format

### Success Response
```json
{
  "code": 200,
  "message": "Success",
  "data": {
    "total": 5,
    "success": 4,
    "failed": 1,
    "errors": [
      "Row 3: NameEn is required"
    ],
    "created": [
      {
        "id": 1,
        "name_en": "Production Web Server 01",
        "code": "SERVER-PROD-WEB-001",
        // ... other fields
      }
    ]
  }
}
```

### Error Response
```json
{
  "code": 400,
  "message": "File must be a CSV file"
}
```

## ⚠️ Lưu ý quan trọng

1. **Encoding**: File CSV phải được lưu với encoding UTF-8
2. **Headers**: Dòng đầu tiên phải chứa tên các cột chính xác
3. **Bắt buộc**: Các trường `name_en` và `code` là bắt buộc
4. **ID References**: Các trường ID (department_id, equipment_type_id, etc.) phải tồn tại trong database
5. **Date Format**: Ngày tháng phải đúng định dạng ISO 8601
6. **File Size**: Khuyến nghị file không quá 10MB

## 🐛 Xử lý lỗi

### Lỗi thường gặp:

1. **"NameEn is required"**: Thiếu tên thiết bị tiếng Anh
2. **"Code is required"**: Thiếu mã thiết bị
3. **"Failed to parse data"**: Dữ liệu không đúng định dạng
4. **"Failed to create equipment"**: Lỗi database (có thể do duplicate code)

### Cách khắc phục:

1. Kiểm tra lại dữ liệu trong file CSV
2. Đảm bảo các trường bắt buộc không để trống
3. Kiểm tra định dạng ngày tháng
4. Đảm bảo các ID reference tồn tại trong database

## 📝 Ví dụ file CSV hoàn chỉnh

```csv
name_en,name_zh,name_vn,code,serial_number,model,manufacturer,location,purchase_date,warranty_end_date,installation_date,status,ip_address,mac_address,operating_system,description,notes,last_maintenance_date,next_maintenance_date,department_id,equipment_type_id,responsible_user_id,assigned_user_id
Production Web Server 01,生产网络服务器01,Máy chủ Web Sản xuất 01,SERVER-PROD-WEB-001,SN1234567890ABC,Dell PowerEdge R740,Dell,"Server Room A, Rack 2, Unit 1-2",2023-01-15T00:00:00Z,2026-01-15T00:00:00Z,2023-01-20T00:00:00Z,active,192.168.1.10,00:1A:2B:3C:4D:5E,Ubuntu Server 22.04 LTS,Primary web server for production environment,Configured with redundant power supplies,2024-12-01T00:00:00Z,2025-03-01T00:00:00Z,1,1,1,1
Database Server 02,数据库服务器02,Máy chủ Cơ sở dữ liệu 02,SERVER-PROD-DB-002,SN9876543210XYZ,HPE ProLiant DL380,HPE,"Server Room B, Rack 1, Unit 3-4",2023-03-10T00:00:00Z,2026-03-10T00:00:00Z,2023-03-15T00:00:00Z,active,192.168.1.11,00:2B:3C:4D:5E:6F,CentOS 8,Main database server for production data,Requires monthly backups,2025-01-15T00:00:00Z,2025-04-15T00:00:00Z,2,1,2,2
``` 