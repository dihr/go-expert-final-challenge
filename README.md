# go-expert-final-challenge

```
Olá devs!

Agora é a hora de botar a mão na massa. Pra este desafio, você precisará criar a listagem das orders.

Esta listagem precisa ser feita com:

- Endpoint REST (GET /order)

- Service ListOrders com GRPC

- Query ListOrders GraphQL

Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.
```

# Requirements

- GO 1.80 or higher
- Docker and Docker compose
- Golang migrations

# Running Project

1. Clone the repository:

   ```shell
   git clone git@github.com:dihr/go-expert-final-challenge.git
   ```
2. Sync dependencies with:
   ```shell
   go mod tidy
   ```   
3. Start docker compose services:

   ```shell
   docker-compose up -d
   ```
4. Run up migrations:

   ```shell
   migrate -path=sql/migrations -database="mysql://root:root@tcp(localhost:3306)/orders" -verbose up
   ```
5. Enter cmd/ordersystem folder:

   ```shell
   cd cmd/ordersystem/
   ```
6. Run main and wire file:

   ```shell
   go run main.go wire_gen.go
   ```   

# Testing the ListOrder endpoints

- list orders through GRPC:
  ![img_3.png](img_3.png)

- list orders through graphQL:
  ![img_1.png](img_1.png)

- list orders through REST:
  ![img_2.png](img_2.png)