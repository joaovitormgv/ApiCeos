Configuração da Base de Dados
Este projeto utiliza PostgreSQL como base de dados. Aqui estão as instruções para configurar a base de dados.

Pré-requisitos
PostgreSQL instalado em sua máquina.
Configuração
Abra o terminal e conecte-se ao PostgreSQL. Você pode fazer isso com o comando psql.

Caso queria usar a mesma configuração de database e tables usada no projeto:
Execute o script SQL localizado em sql/setup.sql para criar a base de dados e a tabela de usuários. Você pode fazer isso com o comando psql -f sql/setup.sql.

O script SQL irá criar uma base de dados chamada ceos e uma tabela chamada users com as colunas id, first_name, last_name e email.

Caso queira criar as databases e tables com seus próprios nomes e configurações:
Crie um arquivo .env na raiz do projeto e escreva o seguinte:
DB_CONNECTION_STRING="sua_string_de_conexão"

Conexão com a Base de Dados
O arquivo main.go contém a string de conexão com a base de dados. A string de conexão atual é host=localhost port=5432 user=postgres dbname=ceos sslmode=disable.

Execução do Projeto
Depois de configurar a base de dados, você pode executar o projeto com o comando go run main.go. O servidor irá iniciar na porta 3000.

Se tudo estiver configurado corretamente, você deverá ser capaz de acessar as rotas /users, /users/:id e /users para criar, obter e listar usuários, respectivamente.

