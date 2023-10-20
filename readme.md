# Imersão FullCycle - Desafio 01

## Escopo

Criação de um server gRPC em Golang que realiza as operações abaixo, baseada no protofile "product.proto":
 - Criar produtos
 - Cosultar produtos

## Requisitos:

 - Utilizar banco SQLite
 - Utilizar ORM GORM
 - Usar AutoMigrate
 - Rodar serviço na porta 50051, acessível via localhost
 - Execução via "go run main.go"

 ## Considerações de solução
 - Como não houve um requisito específico de persistência, foi Utilizada persistencia apenas na memória, apenas para validação do serviço.
 - Arquivo main.go encontra-se na pasta "cmd", para iniciar rodar o comando go run ./cmd/main.go

 # Imersão FullCycle - Desafio 02

 ## Escopo

 Criar uma aplicação Nest.js para atuar como gRPC client do server criado no desafio 01 , com 2 endpoints REST para integrar com o gRPC:
 - GET /products
 - POST /products

 As duas aplicações devem rodar em containers docker, iniciados automaticamente a partir e um mesmo docker-compose.yaml

## Requisitos
- Subir a aplicação completa com o comando "docker compose up"
- As das aplicações devem ser disponibilizadas no mesmo docker-compose.yaml
- Server gRPC deve subir automaticamente na porta 50051
- Aplicação Nest.js deve subir automativamente na porta 3000
- o endpoint POST /products deve fazer chamada para o CreateProduct do gRPC server e retornar os dados do produto criado
- o endpoint GET /products deve fazer chamada para o FindProducts do gRPC server e retornar a lista de produtos do gRPC server

### Execução
- no diretório principal, executar "docker compose up"
> OBS.: A aplicação Nest.js sobe mais rápido que a gRPC devido ao download das dependências do go. Acopanhar os logs até que as duas aplicações retornem a mensagem de "started"

### TODO
- GET /product não trata o erro do gRPC "no products found" (mensagem é mostrada no console do container Nest.js), quando ainda não há nenhum produto criado
- POST /product não trata os erros de validação de campos do gRPC (mensagem é mostrada no console do container Nest.js)
