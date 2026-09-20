
# Posts crud

### Sobre:
 Este repositório se trata de um crud de users e posts, eu estou desenvolvendo isso
 por motivos de prática em criação de backends que envolvam Docker e Postgresql.

### Linguagem:
 Go Lang

### Tecnologias: 
 Docker,Postgres,Gin-Gonic

## Rotas:
### /Users
- #### GET /users retorna todos os usuários
- #### POST /users WITH JSON cria um usuário:
``` 
{
 "username":"SeuNome",
 "bio":"Conte sobre você"
}
```
### /Posts
- #### GET /posts retorna todos os posts
- #### GET /posts/:id retorna um post selecionado por id
- #### POST /posts WITH JSON cria um post:
```
{
  "username":"Nome",
  "body":"Olá, tudo bem?"
}
```
- #### DELETE /posts/:id deleta um post selecionado por id
## Blog
### Status 01: 
 Este projeto ainda está em desenvolvimento, até agora há o módulo post que tem repository
 com um método de Insert e um service com um método de AddPost, além da conexão do banco 
 de dados postgresql e o docker file.

### Status 02:
 Agora neste projeto há uma função SetupTables que serve para inicializar
 as tables, além de agora eu estar utilizando ENVs e também criei duas rotas
 para adicionar e listar os posts.

### Status 03:
 Status de hoje é que agora criei rotas de findOne,findMany e deleteOne,
 tudo com repository e service e coloquei no ar a API e o banco de dados postresql
### Status 04:
 O Status de hoje é que eu criei uma tabela de users, criei duas rotas de criar e listar usuários tudo com repository e service incluso.
