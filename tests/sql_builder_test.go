package tests

import (
	"strings"
	"testing"

	"github.com/oliveira533/cubic_ORM.git/internal/utils"
	models_test "github.com/oliveira533/cubic_ORM.git/tests/models"
)

type MySQLDialect struct{}

func (MySQLDialect) Name() string { return "mysql" }

func (MySQLDialect) Placeholder(i int) string {
	return "?"
}

func (MySQLDialect) InsertSuffix() string {
	return "" // MySQL não usa RETURNING
}

type PostgreSQLDialect struct{}

func (PostgreSQLDialect) Name() string { return "postgres" }

func (PostgreSQLDialect) Placeholder(i int) string {
	return "$1"
}

func (PostgreSQLDialect) InsertSuffix() string {
	return "RETURNING id" // PostgreSQL usa RETURNING
}

func TestSQLBuilderStruct(t *testing.T) {
	t.Run("Deve gerar SQL de INSERT para MySQL corretamente", func(t *testing.T) {
		user := models_test.User{
			Name:  "João Silva",
			Email: "joao@email.com",
		}

		sql, args, err := utils.BuildInsertQuery(MySQLDialect{}, user)

		// Verifica se não houve erro
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		// Verifica se contém INSERT
		if !strings.Contains(strings.ToUpper(sql), "INSERT") {
			t.Errorf("SQL deve conter 'INSERT', obtido: %s", sql)
		}

		// Verifica se contém nome da tabela
		if !strings.Contains(sql, "user") {
			t.Errorf("SQL deve conter nome da tabela 'user', obtido: %s", sql)
		}

		// Verifica se contém colunas (excluindo ID que tem auto_increment)
		expectedColumns := []string{"name", "email"}
		for _, col := range expectedColumns {
			if !strings.Contains(sql, col) {
				t.Errorf("SQL deve conter coluna '%s', obtido: %s", col, sql)
			}
		}

		// Verifica se NÃO contém ID (que tem auto_increment)
		if strings.Contains(sql, "id") {
			t.Errorf("SQL não deve conter coluna 'id' (auto_increment), obtido: %s", sql)
		}

		// Verifica placeholders
		expectedPlaceholders := []string{"?", "?"}
		for _, placeholder := range expectedPlaceholders {
			if !strings.Contains(sql, placeholder) {
				t.Errorf("SQL deve conter placeholder '%s', obtido: %s", placeholder, sql)
			}
		}

		// Verifica argumentos
		if len(args) != 2 {
			t.Errorf("Número de argumentos esperado: 2, obtido: %d", len(args))
		}

		if args[0] != "João Silva" {
			t.Errorf("Primeiro argumento esperado: 'João Silva', obtido: %v", args[0])
		}

		if args[1] != "joao@email.com" {
			t.Errorf("Segundo argumento esperado: 'joao@email.com', obtido: %v", args[1])
		}

		t.Logf("SQL gerado: %s", sql)
		t.Logf("Argumentos: %v", args)
	})

	t.Run("Deve gerar SQL de INSERT para PostgreSQL corretamente", func(t *testing.T) {
		user := models_test.User{
			Name:  "Maria Santos",
			Email: "maria@email.com",
		}

		sql, args, err := utils.BuildInsertQuery(PostgreSQLDialect{}, user)

		// Verifica se não houve erro
		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		// Verifica se contém INSERT
		if !strings.Contains(strings.ToUpper(sql), "INSERT") {
			t.Errorf("SQL deve conter 'INSERT', obtido: %s", sql)
		}

		// Verifica se contém RETURNING
		if !strings.Contains(strings.ToUpper(sql), "RETURNING") {
			t.Errorf("SQL deve conter 'RETURNING', obtido: %s", sql)
		}

		// Verifica placeholders do PostgreSQL
		if !strings.Contains(sql, "$1") {
			t.Errorf("SQL deve conter placeholder '$1', obtido: %s", sql)
		}

		t.Logf("SQL gerado: %s", sql)
		t.Logf("Argumentos: %v", args)
	})

	t.Run("Deve funcionar com ponteiro de struct", func(t *testing.T) {
		user := &models_test.User{
			Name:  "Pedro Costa",
			Email: "pedro@email.com",
		}

		sql, args, err := utils.BuildInsertQuery(MySQLDialect{}, user)

		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		if !strings.Contains(strings.ToUpper(sql), "INSERT") {
			t.Errorf("SQL deve conter 'INSERT', obtido: %s", sql)
		}

		if len(args) != 2 {
			t.Errorf("Número de argumentos esperado: 2, obtido: %d", len(args))
		}

		t.Logf("SQL gerado com ponteiro: %s", sql)
	})

	t.Run("Deve ignorar campos com auto_increment", func(t *testing.T) {
		user := models_test.User{
			ID:    999, // Este valor deve ser ignorado
			Name:  "Ana Lima",
			Email: "ana@email.com",
		}

		sql, args, err := utils.BuildInsertQuery(MySQLDialect{}, user)

		if err != nil {
			t.Errorf("Erro inesperado: %v", err)
		}

		// Verifica se ID não está na query
		if strings.Contains(sql, "id") {
			t.Errorf("SQL não deve conter coluna 'id' (auto_increment), obtido: %s", sql)
		}

		// Verifica se ID não está nos argumentos
		for i, arg := range args {
			if arg == 999 {
				t.Errorf("Argumento %d não deve conter ID (auto_increment): %v", i, arg)
			}
		}

		// Deve ter apenas 2 argumentos (name e email)
		if len(args) != 2 {
			t.Errorf("Número de argumentos esperado: 2, obtido: %d", len(args))
		}

		t.Logf("SQL gerado (ignorando auto_increment): %s", sql)
		t.Logf("Argumentos: %v", args)
	})
}

// Teste de benchmark para performance
func BenchmarkBuildInsertQuery(b *testing.B) {
	user := models_test.User{
		Name:  "Test User",
		Email: "test@email.com",
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		utils.BuildInsertQuery(MySQLDialect{}, user)
	}
}
