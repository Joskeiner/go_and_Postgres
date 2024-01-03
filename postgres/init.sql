
-- init.sql

-- Crear la base de datos si no existe
CREATE DATABASE Store;

-- Conectar a la base de datos Store
\c Store;

-- Crear la tabla 'stocks' si no existe

CREATE TABLE stocks (
      stockid SERIAL PRIMARY KEY,
      name VARCHAR(255) NOT NULL,
      price BIGINT NOT NULL,
      company VARCHAR(255) NOT NULL
    );

 \i /docker-entrypoint-initdb.d/init.sql;
