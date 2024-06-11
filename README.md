***API STORE***

Applicação para cadastro de usuário. categorias e produtos.

A aplicação é feita com go e postgres, sqlc como ORM, chi como manipulador http.

O postgres esta rodando com o docker, com o comando ```docker compose up -d --build```

Após o banco ficar on, precisa rodar as migrations, criei um script com makefile para isso, basta, os comandos os disponiveis são: 
```create_migration``` para criar uma nova migration
```migrate_up``` para rodar as migrations
```migrate_down``` para descartar as migrations

Então, para iniciar o projeto, rode ```migrate_up```, apos isso, ```go run cmd/main.go```.

Autenticação:

A maioria dos endpoints da API exige um token de portador válido para autorização. Você pode obter um token de portador fazendo login (consulte a rota Login). O token de portador deve ser incluído no cabeçalho Authorization usando o formato:
Authorization: Bearer {seu_token_aqui}

Rotas:


**Produtos**

**Criar Produto (POST /product)**

Cria um novo produto. Sem necessidade de autorização.

Corpo: JSON contendo detalhes do produto:

title: Título do produto (obrigatório)
description: Descrição do produto (opcional)
categories: Array de IDs de categorias do produto (obrigatório)
price: Preço do produto em centavos (obrigatório)
Exemplo:
```json
{
  "title": "Camisa Nike",
  "description": "Camisa Nike Dri-FIT para corrida",
  "categories": ["3f768233-7049-4779-bd22-65b0ad09d980"],
  "price": 149900
}
```

**Atualizar Produto (PATCH /product/:id)**

Atualiza um produto existente. Requer autorização.

Variável do Caminho:

id: ID do produto a ser atualizado
Corpo: JSON contendo detalhes do produto para atualizar (atualizações parciais permitidas):

title: Título do produto (opcional)
description: Descrição do produto (opcional)
categories: Array de IDs de categorias do produto (opcional)
price: Preço do produto em centavos (opcional)
Exemplo:

```json
{
  "title": "Camisa Nike Atualizada",
  "price": 159900
}
```
**Excluir Produto por ID (DELETE /product/:id)**

Exclui um produto. Requer autorização.

Variável do Caminho:

id: ID do produto a ser excluído

**Obter Todos os Produtos (GET /products)**

Recupera uma lista de todos os produtos. Requer autorização.




***Categorias***

**Criar Categoria (POST /category)**

Cria uma nova categoria. Requer autorização.

Corpo: JSON contendo detalhes da categoria:

title: Título da categoria (obrigatório)
Exemplo:
```json
{
  "title": "Roupas"
}
```

**Encontrar Todas as Categorias (GET /categories)**

Recupera uma lista de todas as categorias. Requer autorização.




***Usuários***
**Criar Usuário (POST /user)**

Cria um novo usuário. Sem necessidade de autorização.

Corpo: JSON contendo detalhes do usuário:

name: Nome do usuário (obrigatório)
email: E-mail do usuário (obrigatório)
password: Senha do usuário (obrigatório)
cep: CEP do usuário (opcional)
Exemplo:

```json
{
  "name": "João Silva",
  "email": "joaosilva@email.com",
  "password": "senha1234",
  "cep": "13500163"
}
```

**Atualizar Usuário (PATCH /user/me)**

Atualiza o usuário logado no momento. Requer autorização.

Corpo: JSON contendo detalhes do usuário para atualizar (atualizações parciais permitidas):

name: Nome do usuário (opcional)
password: Senha do usuário (opcional)
Exemplo:

```json
{
  "name": "João Silva Atualizado"
}
```

**Obter Usuário Logado (GET /user/me)**

Recupera detalhes do usuário logado no momento. Requer autorização.

**Atualizar Senha (PATCH /user/password)**

Atualiza a senha do usuário logado no momento. Requer autorização.

Corpo: JSON contendo nova senha e senha antiga para
Exemplo:
```json
{
    "password": "123445@1",
    "old_password": "12345%#1"
}
```
