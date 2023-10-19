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
