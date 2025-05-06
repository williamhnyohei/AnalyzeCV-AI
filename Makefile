# Makefile para o AnalyzeCV-AI
# Comandos para facilitar o desenvolvimento do backend

.PHONY: run test tidy up down

# Rodar a API localmente
run:
	cd backend && go run cmd/server/main.go

# Rodar os testes do backend
test:
	cd backend && go test ./...

# Atualizar as dependências Go
tidy:
	cd backend && go mod tidy

# Subir infraestrutura com docker-compose
up:
	docker-compose up -d

# Derrubar infraestrutura do docker-compose
down:
	docker-compose down
