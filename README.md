# Log-Processor Microservice

**Purpose:**
Consumes logs from RabbitMQ, validates, and saves them to the database. Triggers rule engine for alerting.

---

## 🏗️ Architecture
- Consumes logs from RabbitMQ
- Validates and stores logs in PostgreSQL
- Runs detection rules and creates alerts

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
- Rule engine integration
- Alert creation

---

## 🚀 Setup & Configuration

### 1. Clone & Configure
```bash
git clone <repo-url>
cd log-processor
cp .env.example .env
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
go run cmd/main.go
```

### Docker Compose
```bash
docker-compose up --build
```

---

## 📡 Endpoints

### GET /health
Health check endpoint.

**Example curl:**
```bash
curl http://localhost:8081/health
```

---

## 🧪 Testing
```bash
go test ./...
```

---

## 🤝 Contributing
- Fork, branch, PR
- Write tests for new features
- Follow Go best practices

---

## 📄 License
MIT (or specify) 