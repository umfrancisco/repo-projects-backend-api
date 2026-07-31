# Repo Projects API

API backend desenvolvida em Go utilizando o **Encore Cloud**, com integração a banco de dados PostgreSQL para armazenamento e consulta de projetos.

## Descrição

Esta API fornece endpoints para gerenciar projetos, permitindo:

* Criar projetos
* Listar projetos armazenados
* Persistir dados em banco PostgreSQL
* Inicializar dados automaticamente (seed)

Cada projeto contém:

* Nome
* Descrição
* Linguagem
* Link do repositório

---

## Tecnologias

* Go (Golang)
* Encore Cloud
* PostgreSQL
* SQL (migrations)

---

## Setup do Projeto

### 1. Instalar Encore

```bash
curl -L https://encore.dev/install.sh | bash
```

---

### 2. Rodar localmente

```bash
encore run
```

A API estará disponível em:

```
http://localhost:4000
```

---

## Banco de Dados

O projeto utiliza PostgreSQL gerenciado automaticamente pelo Encore.

### Migration

Tabela criada:

```sql
CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    language TEXT NOT NULL,
    link TEXT NOT NULL
);
```

---

## Endpoints

### 🔹 Listar projetos

```
GET /projects
```

#### Resposta:

```json
{
  "projects": [
    {
      "name": "project-name",
      "description": "short-desc",
      "language": "java",
      "link": "link-to-github-repo"
    }
  ]
}
```

---

## Deploy

Para fazer deploy:

```bash
encore deploy
```

O Encore irá provisionar automaticamente:

* API
* Banco PostgreSQL
* Infraestrutura

---

## Autor

Francisco Guitler

---

## 📄 Licença

Este projeto está sob a licença MIT.
