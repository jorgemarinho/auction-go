
# Auction App

## Descrição

Este é um sistema de leilões desenvolvido em Go. O sistema permite criar leilões, fazer lances e consultar leilões ativos. A aplicação pode ser executada localmente usando Docker e Docker Compose.

## Como Executar

### Requisitos

- Docker
- Docker Compose

### Passos para Executar Localmente

1. **Clone o repositório:**

   ```sh
   git clone git@github.com:jorgemarinho/auction-go.git
   cd auction-go
   ```

2. **Construir as Imagens**: Na raiz do projeto, execute:

   ```sh
   docker-compose build
   ```

3. **Executar o Docker Compose**: Na raiz do projeto (onde está o arquivo `docker-compose.yml`), execute:

   ```sh
   docker-compose up
   ```

4. **Acesse a aplicação:**

   A aplicação estará disponível em `http://localhost:8080`.

//precisa criar usuario
//preicsa criar o leilao
//precisar criar o bid

### Exemplos de Requisição

- **Para criar leilão:**

   ```sh
   curl -X POST http://localhost:8080/auction \
   -H "Content-Type: application/json" \
   -d '{
        "product_name": "Nintendo 64",
        "category": "Video Game",
        "description": "video game da nintendo com 64 bits",
        "condition": 1
       }'
   ```

- **Para buscar leilão por ID:**

   ```sh
   curl -X GET "http://localhost:8080/auction/{auctionId}" \
   -H "Content-Type: application/json"
   ```

- **Para buscar leilões por parâmetros de consulta:**

   ```sh
   curl -X GET "http://localhost:8080/auction?category=Video%Game&condition=1&status=0" \
   -H "Content-Type: application/json"
   ```
   
   Subir o serviço docker, caso não esteja iniciado

service docker start

Primeira vez, fazer um build

docker-compose up --build 

Outras vezes

docker-compose up -d




docker container run -d -p 27017:27017 name auctionDB mongo