-- Arquivo para caso precise criar a database e tabelas conforme o projeto main.go.
-- Ou podem alterar a string de conexão em main.go para uma database sua.

-- Criar database
CREATE DATABASE IF NOT EXISTS ceos;

-- Conectar com a database
USE ceos;

-- Criar tabela de usuários
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255),
    email VARCHAR(255) NOT NULL,
);