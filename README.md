# go-manage-product

## Descrição

O projeto **go-manage-product** é um projeto de exemplo para a criação de um sistema de gerenciamento de produtos.

Este projeto foi escrito na linguagem Golang usando os princípios do DDD.
A aplicação está arquitetada de forma a ser mais escalável e reutilizável.

## Tecnologias envolvidas

- Docker
- Gin
- Gorm
- Postgres

## Como usar

### Requisitos
- Docker
- docker-compose


### 1-Instalação
```bash
  docker-compose up
```
### 2-Execução
A porta padrão do serviço é 8080.

Primeiramente será necessário rodar as migrações do banco de dados.

```bash
  GET http://localhost:8080/migration
```

Para usar os endpoints será necessário autenticar a aplicação no seguinte endereço:

```bash
  POST http://localhost:8080/login
```
Com o usuário e senha padrão:

```sh
  {
    "username": "admin",
    "password": "123456"
  }
```
O cabeçalho da requisição deve possuir o Authorization com o seguinte valor:

```sh
  Bearer aqui_vai_o_token
```

Os demais endereços para acessar os recursos da aplicação são:

- Cadastro de produto:
```bash
  POST   http://localhost:8080/api/v1/product
  {
    "code": "code1",
    "name": "produto x",
    "stock_total": 10,
    "stock_cute": 8,
    "price_from": 2.99,
    "price_to": 2.00
  }
```

- Alteração de produto
```bash
  PATCH  http://localhost:8080/api/v1/product/:id
  {
    "name": "produto nome",
    "stock_total": 22,
    "stock_cute": 3,
    "price_from": 5.99,
    "price_to": 1.00
  }
```

- Busca de produto por código
```bash
  GET    http://localhost:8080/api/v1/product/:id
```

- Listagem de produtos
```bash
  GET    http://localhost:8080/api/v1/products
```

- Remoção de produto
```bash
  DELETE http://localhost:8080/api/v1/product/:id
```
