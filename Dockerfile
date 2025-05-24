# Stage 1: Build do binário usando imagem golang oficial 1.22.5 (Linux)
FROM golang:1.22.5-alpine AS builder

# Dependências para compilação estática (gcc, musl-dev)
RUN apk add --no-cache build-base ca-certificates

WORKDIR /app

# Copia arquivos de modulação para cache de dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o código fonte do projeto
COPY . .

# Build estático para Linux (disable CGO)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o fidelity ./cmd/fidelity

# Stage 2: Imagem final minimalista para rodar o binário
FROM alpine:3.18

WORKDIR /app

# Copia o binário compilado da stage anterior
COPY --from=builder /app/fidelity .

# Copia o arquivo de configuração local_env.yaml para dentro da imagem
COPY config/local_env.yaml ./config/local_env.yaml

# Expõe a porta padrão da aplicação
EXPOSE 8080

# Comando padrão para rodar o binário
ENTRYPOINT ["./fidelity"]
