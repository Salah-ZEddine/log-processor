# Log-Processor Microservice

  

**Purpose:**

Consumes logs from RabbitMQ, validates, and saves them to the database. Triggers rule engine for alerting (not yet implemented adds more complexity to the system requires tight configuration).

  

---

  

## 🏗️ Architecture

- Consumes logs from RabbitMQ

- Validates and stores logs in PostgreSQL


[![](https://mermaid.ink/img/pako:eNp9lF1vmzAUhv-K5d60EiThI4GgqlIJ4SqV2rXqpI1pMnAC1gzObJMmjfLfZyDJ2IdqbjB-n_MeHx98wBnPAQd4zfhbVhKh0EuU1EiP-68J_kTSlKqHp9tUjO-uGS_k958NNHCT4G_INO9QqEUrXpiPgmcgJRfoGcSWZqAFfZiw0y207pUwmhMFSAPoupG0LpDUlpCjSifBbjqXW9mkdwVVZZOOMl6NJWGkNN8hz2kNY1pnNIdamTmRZcqJyE0tqnh9O265i-uic4206zPZAlIcPXKpCgHPTysUEUVSIvsce71UewboHq0pY8FVHC7nS8-QSvAfoKexN3Gd09R8o7kqA3uzG5LhmYxjf2n9JsPFxI4-JBcncunEdhxdSNe-n8bTD8noTFrxNF5eSHsezuLZf0hs4ELQHAdKNGDgCkRF2ik-tFETrEqodE0C_ZrDmjRMJdgYLL0SQUnKQLaaQ59JgjeCVkTsF5xx0cNX1trzU_cEDzQvsFND3bob_-pCLnIQQ6XTjYGS6Wb40FKBUPSvvGDSPgluNcfTzi59_8emskZsT6XQjULlIHCpKrYiKbCuDF0pz0uNhAey-9xW_Lw2tJKgw1K1X8EWWB-ccd62YVIf9eFsSP2F8-p8PoI3RYmDNWFSz5pN--tElBSCVJevAuquUk2tcGA7XQwcHPAOB87cHTme503dtnlt38B7HFjeyHMc17amE9-aW_7saOD3znMy8j1rZrmW7bkz39EKA0NOFRcP_QXR3RPHX83cUe8?type=png)](https://mermaid.live/edit#pako:eNp9lF1vmzAUhv-K5d60EiThI4GgqlIJ4SqV2rXqpI1pMnAC1gzObJMmjfLfZyDJ2IdqbjB-n_MeHx98wBnPAQd4zfhbVhKh0EuU1EiP-68J_kTSlKqHp9tUjO-uGS_k958NNHCT4G_INO9QqEUrXpiPgmcgJRfoGcSWZqAFfZiw0y207pUwmhMFSAPoupG0LpDUlpCjSifBbjqXW9mkdwVVZZOOMl6NJWGkNN8hz2kNY1pnNIdamTmRZcqJyE0tqnh9O265i-uic4206zPZAlIcPXKpCgHPTysUEUVSIvsce71UewboHq0pY8FVHC7nS8-QSvAfoKexN3Gd09R8o7kqA3uzG5LhmYxjf2n9JsPFxI4-JBcncunEdhxdSNe-n8bTD8noTFrxNF5eSHsezuLZf0hs4ELQHAdKNGDgCkRF2ik-tFETrEqodE0C_ZrDmjRMJdgYLL0SQUnKQLaaQ59JgjeCVkTsF5xx0cNX1trzU_cEDzQvsFND3bob_-pCLnIQQ6XTjYGS6Wb40FKBUPSvvGDSPgluNcfTzi59_8emskZsT6XQjULlIHCpKrYiKbCuDF0pz0uNhAey-9xW_Lw2tJKgw1K1X8EWWB-ccd62YVIf9eFsSP2F8-p8PoI3RYmDNWFSz5pN--tElBSCVJevAuquUk2tcGA7XQwcHPAOB87cHTme503dtnlt38B7HFjeyHMc17amE9-aW_7saOD3znMy8j1rZrmW7bkz39EKA0NOFRcP_QXR3RPHX83cUe8)
  

---

  

## ⚙️ Tech Stack

- Go 1.24+

- RabbitMQ (consumer)

- PostgreSQL (GORM)

- Docker & Docker Compose

  

---

  

## 🧱 Features

- Message queue consumer

- Log validation

- Database persistence


  

---

  

## 🚀 Setup & Configuration

  

### 1. Clone & Configure

```bash

git  clone https://github.com/Salah-ZEddine/log-processor.git

cd  log-processor

cp  .env.example  .env

```

  

### 2. Environment Variables
| Variable         | Description                | Example         |
|-----------------|----------------------------|-----------------|
| RABBITMQ_URL    | RabbitMQ connection string | amqp://...      |
| DB_HOST         | Database host              | localhost       |
| DB_PORT         | Database port              | 5432            |
| DB_USER         | Database user              | postgres        |
| DB_PASSWORD     | Database password          | password        |
| DB_NAME         | Database name              | log_db          |
| ...             | ...                        | ...             |

---

## 🐳 Running

  

### Local

```bash

go  run  cmd/main.go

```

  

### Docker Compose

```bash

docker-compose  up  --build

```

  

---

  


  

## 🧪 Testing

```bash

go  test  ./...

```
**no tests yet**
  

---