# Go ORM

Um ORM simples e extensível escrito em **Go**, com suporte a múltiplos bancos de dados (Postgres, MySQL, SQLite).  
Inspirado em bibliotecas como GORM, mas focado em **clareza, modularidade e aprendizado**.

---

## 📂 Estrutura do Projeto

```
cubic/
├── go.mod
├── go.sum
├── cmd/
│   └── orm-cli/               # CLI opcional para migrações e geração de código
├── internal/                  # Código interno que não deve ser exportado
│   ├── db/                    # Drivers e inicialização de conexão
│   │   ├── connection.go      # Gerencia pool de conexões
│   │   ├── postgres.go        # Driver Postgres
│   │   ├── mysql.go           # Driver MySQL
│   │   └── sqlite.go          # Driver SQLite
│   ├── dialects/              # Específicos de cada banco
│   │   ├── postgres.go
│   │   ├── mysql.go
│   │   └── sqlite.go
│   ├── migrator/              # Lógica de migrações
│   │   ├── migrator.go
│   │   └── file_loader.go
│   └── utils/                 # Utilitários internos
│       ├── mapping.go         # Utilidades com reflection (para mapear structs)
│       ├── sql_builder.go     # Construtor genérico de queries
│       └── logger.go          # Logger interno
├── pkg/                       # API pública do ORM
│   ├── orm.go                 # Ponto de entrada (ex: orm.Open, orm.Model)
│   ├── model.go               # Definição de entidades e metadata
│   ├── query.go               # Query Builder genérico
│   ├── session.go             # Sessão/transaction manager
│   ├── migration.go           # API pública para migrações
│   └── errors.go              # Tratamento de erros padrão
├── examples/                  # Exemplos de uso
│   ├── basic.go
│   └── relationships.go
└── tests/                     # Testes unitários e integração
    ├── orm_test.go
    ├── query_test.go
    ├── migration_test.go
    └── integration_test.go
```