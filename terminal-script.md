Remove docs:
    rm -r docs

Create new swagger docs: 
    swag init -g cmd/api/main.go

Access swager:
    http://localhost:8080/swagger/index.html
    
Docker:
    docker compose up -d
    docker compose --build
    docker compose up --build -d
    docker compose down