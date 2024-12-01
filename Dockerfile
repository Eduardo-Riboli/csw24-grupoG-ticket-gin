FROM golang:1.21-bullseye

# Definir o diretório de trabalho
WORKDIR /app

# Copiar todo o projeto para o contêiner
COPY . .

# Copiar o .env para o container
COPY .env .env

# Instalar dependências
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest
# RUN swag init

# Expor a porta que a aplicação vai usar
EXPOSE 8080

# Comando para rodar a aplicação Go
CMD ["sh", "-c", "swag init && go run main.go"]