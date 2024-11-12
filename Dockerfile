# Use uma imagem oficial como uma etapa de build
FROM golang:1.21-bullseye AS builder

# Definir o diretório de trabalho
WORKDIR /app

# Copie e baixe as dependências do módulo
COPY go.mod go.sum ./
RUN go mod download

# Copie o restante dos arquivos da aplicação
COPY . .

# Instalar o Swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Inicializar o Swagger
RUN swag init

# Compile a aplicação
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Utilize uma imagem menor para o contêiner final
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
# COPY --from=builder /app/swagger ./swagger
COPY .env .env

# Expor a porta que a aplicação vai usar (opcional)
EXPOSE 8080

# Comando para rodar a aplicação Go
CMD ["./main"]