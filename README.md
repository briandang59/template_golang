<<<<<<< HEAD
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

=======
# be_scada



## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Add your files

- [ ] [Create](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#create-a-file) or [upload](https://docs.gitlab.com/ee/user/project/repository/web_editor.html#upload-a-file) files
- [ ] [Add files using the command line](https://docs.gitlab.com/ee/gitlab-basics/add-file.html#add-a-file-using-the-command-line) or push an existing Git repository with the following command:

```
cd existing_repo
git remote add origin http://10.2.1.161/briandang/be_scada.git
git branch -M main
git push -uf origin main
```

## Integrate with your tools

- [ ] [Set up project integrations](http://gitlab.example.com/briandang/be_scada/-/settings/integrations)

## Collaborate with your team

- [ ] [Invite team members and collaborators](https://docs.gitlab.com/ee/user/project/members/)
- [ ] [Create a new merge request](https://docs.gitlab.com/ee/user/project/merge_requests/creating_merge_requests.html)
- [ ] [Automatically close issues from merge requests](https://docs.gitlab.com/ee/user/project/issues/managing_issues.html#closing-issues-automatically)
- [ ] [Enable merge request approvals](https://docs.gitlab.com/ee/user/project/merge_requests/approvals/)
- [ ] [Set auto-merge](https://docs.gitlab.com/ee/user/project/merge_requests/merge_when_pipeline_succeeds.html)

## Test and Deploy

Use the built-in continuous integration in GitLab.

- [ ] [Get started with GitLab CI/CD](https://docs.gitlab.com/ee/ci/quick_start/index.html)
- [ ] [Analyze your code for known vulnerabilities with Static Application Security Testing (SAST)](https://docs.gitlab.com/ee/user/application_security/sast/)
- [ ] [Deploy to Kubernetes, Amazon EC2, or Amazon ECS using Auto Deploy](https://docs.gitlab.com/ee/topics/autodevops/requirements.html)
- [ ] [Use pull-based deployments for improved Kubernetes management](https://docs.gitlab.com/ee/user/clusters/agent/)
- [ ] [Set up protected environments](https://docs.gitlab.com/ee/ci/environments/protected_environments.html)

***

# Editing this README

When you're ready to make this README your own, just edit this file and use the handy template below (or feel free to structure it however you want - this is just a starting point!). Thanks to [makeareadme.com](https://www.makeareadme.com/) for this template.

## Suggestions for a good README

Every project is different, so consider which of these sections apply to yours. The sections used in the template are suggestions for most open source projects. Also keep in mind that while a README can be too long and detailed, too long is better than too short. If you think your README is too long, consider utilizing another form of documentation rather than cutting out information.

## Name
Choose a self-explaining name for your project.

## Description
Let people know what your project can do specifically. Provide context and add a link to any reference visitors might be unfamiliar with. A list of Features or a Background subsection can also be added here. If there are alternatives to your project, this is a good place to list differentiating factors.

## Badges
On some READMEs, you may see small images that convey metadata, such as whether or not all the tests are passing for the project. You can use Shields to add some to your README. Many services also have instructions for adding a badge.

## Visuals
Depending on what you are making, it can be a good idea to include screenshots or even a video (you'll frequently see GIFs rather than actual videos). Tools like ttygif can help, but check out Asciinema for a more sophisticated method.

## Installation
Within a particular ecosystem, there may be a common way of installing things, such as using Yarn, NuGet, or Homebrew. However, consider the possibility that whoever is reading your README is a novice and would like more guidance. Listing specific steps helps remove ambiguity and gets people to using your project as quickly as possible. If it only runs in a specific context like a particular programming language version or operating system or has dependencies that have to be installed manually, also add a Requirements subsection.

## Usage
Use examples liberally, and show the expected output if you can. It's helpful to have inline the smallest example of usage that you can demonstrate, while providing links to more sophisticated examples if they are too long to reasonably include in the README.

## Support
Tell people where they can go to for help. It can be any combination of an issue tracker, a chat room, an email address, etc.

## Roadmap
If you have ideas for releases in the future, it is a good idea to list them in the README.

## Contributing
State if you are open to contributions and what your requirements are for accepting them.

For people who want to make changes to your project, it's helpful to have some documentation on how to get started. Perhaps there is a script that they should run or some environment variables that they need to set. Make these steps explicit. These instructions could also be useful to your future self.

You can also document commands to lint the code or run tests. These steps help to ensure high code quality and reduce the likelihood that the changes inadvertently break something. Having instructions for running tests is especially helpful if it requires external setup, such as starting a Selenium server for testing in a browser.

## Authors and acknowledgment
Show your appreciation to those who have contributed to the project.

## License
For open source projects, say how it is licensed.

## Project status
If you have run out of energy or time for your project, put a note at the top of the README saying that development has slowed down or stopped completely. Someone may choose to fork your project or volunteer to step in as a maintainer or owner, allowing your project to keep going. You can also make an explicit request for maintainers.
>>>>>>> 70d3f6c71a98a67a3a5c84f62f73076c0a051716
