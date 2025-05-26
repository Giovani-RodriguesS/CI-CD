# API

## Descrição geral
Está aplicação Go foi criada para permitir interagir com banco de dados SQL. 
A API fornece um conjunto de operações RESTful para gerenciamento de produtos

## Dependências
- `github.com/gin-gonic/gin v1.10.0`
- `github.com/joho/godotenv v1.5.1`
- `gorm.io/driver/postgres v1.5.11`
- `gorm.io/gorm v1.25.12`

## Arquitetura
Desenvolvida na arquitetura Clean, tem como principal objetivo separar responsabilidades de forma clara, favorecendo a manutenibilidade, testabilidade e escalabilidade da aplicação. Em `internal/controller`, definimos os controllers que serão usados em `internal/router`. Em `internal/usecase`, definimos a lógica de negócio da aplicação. Por fim, em `internal/repository`, configuramos a interação com o banco de dados usando Gorm como ORM para facilitar o processo de inserção, exclução e alteração dos dados. 

## Controllers
### 1. Listar todos os produtos
* **Método:** `GET`
* **Endpoint:** `/products`
* **Descrição:** Recupera a lista completa de produtos disponíveis no sistema.
* **Resposta:** Lista de objetos de produto com status `200 OK`, ou erro `500 Internal Server Error`.

---

### 2. Buscar um produto por ID

* **Método:** `GET`
* **Endpoint:** `/products/:id`
* **Descrição:** Busca um único produto com base no seu ID.
* **Validação:** O ID é validado antes da consulta.
* **Resposta:**

  * Produto encontrado: status `200 OK`
  * Produto não encontrado: status `404 Not Found`
  * Erros de validação ou internos: status `400` ou `500`.

---

### 3. Criar um novo produto

* **Método:** `POST`
* **Endpoint:** `/products`
* **Descrição:** Cria um novo produto a partir de um JSON contendo os dados do produto.
* **Validação:** O corpo da requisição é validado.
* **Resposta:**

  * Sucesso: status `201 Created` com o ID do novo produto.
  * Erro de validação ou criação: status `400` ou `500`.

---

### 4. Atualizar um produto existente

* **Método:** `PUT`
* **Endpoint:** `/products/:id`
* **Descrição:** Atualiza os dados de um produto com base no ID fornecido.
* **Validação:** ID e corpo da requisição são validados.
* **Resposta:**

  * Sucesso: status `202 Accepted`
  * Erros: `400 Bad Request`, `500 Internal Server Error`

---

### 5. Deletar um produto

* **Método:** `DELETE`
* **Endpoint:** `/products/:id`
* **Descrição:** Remove um produto do sistema com base em seu ID.
* **Validação:** O ID é validado antes da exclusão.
* **Resposta:**

  * Sucesso: status `200 OK` com mensagem de confirmação
  * Erro: `400` ou `500`