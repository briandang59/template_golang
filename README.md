# 🏭 SCADA Backend System

Dự án backend hệ thống **SCADA** để quản lý thiết bị công nghiệp, cảnh báo trạng thái, và giao tiếp thời gian thực bằng **WebSocket**.

## ✅ Tech Stack

- 💻 Golang
- 🛢️ PostgreSQL (qua Docker)
- 🐋 Docker Compose
- 📡 Gin (Web Framework)
- 💬 WebSocket
- 🔄 GORM (ORM)
- 🔐 godotenv (quản lý biến môi trường)

---

## 📁 Cấu trúc thư mục

```bash
be_scada/
├── cmd/
│   └── api/               # Entry point chính của backend
│       └── main.go
├── config/                # Load biến môi trường, khởi tạo DB
│   └── config.go
├── internal/
│   ├── model/             # Định nghĩa các struct (Entity DB)
│   │   └── equipment.go
│   ├── repository/        # Tầng truy vấn DB (CRUD)
│   │   └── equipment_repo.go
│   ├── service/           # Tầng xử lý logic nghiệp vụ
│   │   └── equipment_service.go
│   ├── http/
│   │   ├── handler/       # Tầng xử lý request HTTP
│   │   │   ├── equipment_handler.go
│   │   │   └── dependencies.go
│   │   └── routes/        # Tổ chức routing theo module
│   │       └── equipment_routes.go
│   └── websocket/         # Hub + client quản lý kết nối WebSocket
│       └── websocket.go
├── go.mod / go.sum        # Quản lý package
├── .env                   # Thông tin kết nối DB (PORT, DSN, ...)
├── Dockerfile             # Docker image build cho Golang
├── docker-compose.yml     # Tạo network + Postgres + API
└── README.md              # 📄 YOU ARE HERE
```

---

## 📜 Giải thích các file chính

### `cmd/api/main.go` – điểm khởi chạy

```go
func main() {
	_ = godotenv.Load()
	config.Init()

	eqRepo := repository.NewEquipmentRepo()
	eqSvc := service.NewEquipmentService(eqRepo)
	eqHdl := handler.NewEquipmentHandler(eqSvc)

	deps := &handler.Dependencies{
		Equipment: eqHdl,
	}

	r := gin.Default()

	routes.RegisterRoutes(r, deps)

	hub := ws.NewHub()
	go hub.Run()
	r.GET("/ws", ws.ServeWs(hub))

	log.Fatal(r.Run(":" + os.Getenv("PORT")))
}
```

---

### `config/config.go` – khởi tạo DB GORM

```go
func Init() {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("🚫 Không kết nối được DB: %v", err)
    }

    DB = db
    err = db.AutoMigrate(&model.Equipment{})
}
```

---

### `model/equipment.go` – entity thiết bị

```go
type Equipment struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"        gorm:"size:100;not null"`
	Location    string         `json:"location"    gorm:"size:50"`
	Status      string         `json:"status"      gorm:"size:20"`
	Description string         `json:"description" gorm:"type:text"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"-"           gorm:"index"`
}

```

---

### `repository/equipment_repo.go`

```go
func (r *EquipmentRepo) FindAllPaginate(page, size int) ([]model.Equipment, int64, error) {
    var list []model.Equipment
    var total int64
    r.db.Model(&model.Equipment{}).Count(&total)
    r.db.Offset((page-1)*size).Limit(size).Find(&list)
    return list, total, nil
}
```

---

### `service/equipment_service.go`

```go
func (s *EquipmentService) GetAllPaginate(page, size int) ([]model.Equipment, int64, error) {
    return s.repo.FindAllPaginate(page, size)
}
```

---

### `handler/equipment_handler.go`

```go
func (h *EquipmentHandler) GetAll(c *gin.Context) {
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    size, _ := strconv.Atoi(c.DefaultQuery("size", "10"))

    data, total, err := h.svc.GetAllPaginate(page, size)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "code":    200,
        "message": "Success",
        "data":    data,
        "pagination": gin.H{
            "page":      page,
            "page_size": size,
            "total":     total,
        },
    })
}
```

---

### `websocket/websocket.go` – WebSocket Hub

```go
type Hub struct {
    clients    map[*Client]bool
    broadcast  chan []byte
    register   chan *Client
    unregister chan *Client
}

func (h *Hub) Run() {
    for {
        select {
        case client := <-h.register:
            h.clients[client] = true
        case message := <-h.broadcast:
            for client := range h.clients {
                client.send <- message
            }
        }
    }
}
```

---

## 🐋 Dockerfile

```dockerfile
FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o be_scada ./cmd/api

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/be_scada .
COPY .env .
CMD ["./be_scada"]
```

---

## 🛠️ docker-compose.yml

```yaml
services:
  db:
    image: postgres:16
    ports:
      - "5432:5432"
    environment:
      POSTGRES_PASSWORD: 123456
    volumes:
      - scada_pg-data:/var/lib/postgresql/data

  backend:
    build:
      context: .
    ports:
      - "8080:8080"
    depends_on:
      - db
    env_file:
      - .env

volumes:
  scada_pg-data:
```

---

## 🧪 Test API

```bash
GET     http://localhost:8080/api/equipment
POST    http://localhost:8080/api/equipment
PUT     http://localhost:8080/api/equipment/:id
DELETE  http://localhost:8080/api/equipment/:id
```

---

# Server config

```bash
PORT=8080
GIN_MODE=release
HOST=localhost
```


# PostgreSQL DB config

```bash
DB_HOST=db
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=123456
DB_NAME=scada
```



## 🔮 Gợi ý mở rộng

- Xác thực JWT
- Module Users, Alerts, Work Orders
- Export CSV/Excel
- Giao tiếp MQTT với thiết bị

