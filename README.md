# Configuração da Base de Dados

Este projeto utiliza PostgreSQL como base de dados. Aqui estão as instruções para configurar a base de dados.

## Pré-requisitos

- PostgreSQL instalado em sua máquina.

## Configuração

1. Abra o terminal e conecte-se ao PostgreSQL. Você pode fazer isso com o comando `psql -h localhost -U postgres`.

2. Caso queira usar a mesma configuração de database e tables usada no projeto:
    - Após ter se conectado com o comando acima, escreva o comando ` \i caminho/para/setup.sql`
    - Lembre-se de substituir o `caminho/para/setup.sql`pelo caminho real do arquivo setup.sql.
    - Se tudo der certo, isso irá criar uma database de nome ceos e uma tabela de nome users.
    - Para sair do cliente psql basta digitar \q na linha de comando.

3. Caso queira usar databases e tables próprias já existentes:
    - Crie um arquivo `.env` na raiz do projeto e escreva o seguinte:
      ```
      DB_CONNECTION_STRING="sua_string_de_conexão"
      ```

## Conexão com a Base de Dados

O arquivo `main.go` contém a string de conexão com a base de dados. A string de conexão atual é `host=localhost port=5432 user=postgres dbname=ceos sslmode=disable`.

## Execução do Projeto

Depois de configurar a base de dados, você pode executar o projeto com o comando `go run main.go`. O servidor irá iniciar na porta 3000.

Se tudo estiver configurado corretamente, você deverá ser capaz de acessar as rotas `/users`, `/users/:id` e `/users` para criar, obter e listar usuários, respectivamente.
