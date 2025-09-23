# Imagem base oficial do Go
FROM golang:1.22 AS builder

WORKDIR /app

# Copia os arquivos go.mod e go.sum para cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código da aplicação
COPY . .

# Compila a API (binário estático)
RUN go build -o api ./main.go

# ---- Stage final ----
FROM debian:bullseye-slim

WORKDIR /app

# Copia binário da API
COPY --from=builder /app/api .

# Porta exposta
EXPOSE 8080

# Comando de entrada
CMD ["./api"]
