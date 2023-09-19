# API Go com Gin Framework

Esta API REST foi desenvolvida utilizando a linguagem [Go](https://go.dev/), o framework [Gin](https://gin-gonic.com/), o ORM [Gorm](https://gorm.io/) e o banco de dados [PostgreSQL](https://www.postgresql.org).

É uma API simples feita com foco no aprendizado da linguagem e das tecnologias que auxiliaram no desenvolvimento da API.

A API possui um único recurso, ``aluno``. É possível criar, listar, editar e excluir alunos. As buscas específicas podem ser feitas por ID ou por CPF.

As rotas para isso são:

- ``GET /aluno``: lista todos os alunos
- ``GET /aluno/{{id}}``: recupera um aluno por ID
- ``GET /aluno/cpf/{{cpf}}``: recupera um aluno por CPF
- ``POST /aluno``: cria um novo aluno
- ``PUT /aluno/{{id}}``: atualiza um aluno utilizando seu ID
- ``DELETE /aluno/{{id}}``: exclui um aluno por ID

O modelo para criação de um aluno é:

``Nome: string``
``Cpf: string``
``Rg: string``

As colunas ``ID``, ``CreatedAt`` e ``UpdatedAt`` são criadas e gerenciadas automaticamente pelo Gorm.

A API também conta com testes unitários feitos com o pacote [Testify](https://github.com/stretchr/testify) para garantir o correto funcionamento da API.

O arquivo ``docker-compose.example.yml`` é utilizado como base do arquivo utilizado para criação do banco de dados e da plataforma de gerenciamento PgAdmin do postgreSQL. Para utilizar, basta alterar o nome do arquivo para ``docker-compose.yml`` e alterar os valores das variáveis dentro do arquivo. Tendo o Docker instalado e rodando em sua máquina, utilize o comando ``docker compose up`` para criar e executar as imagens.As variáveis em questão são:

- ``POSTGRES_USER``: nome de usuário que terá acesso ao banco de dados
- ``POSTGRES_PASSWORD``: senha do usuário
- ``POSTGRES_DB``: nome do banco de dados que será utilizado
- ``PGADMIN_DEFAULT_EMAIL``: email que será usado para acessar o PgAdmin (não precisa ser um email válido)
- ``PGADMIN_DEFAULT_PASSWORD``: senha do usuário do PgAdmin

## Comandos

- ``go run main.go``: compila e executa a API
- ``go test``: executa os testes automatizados