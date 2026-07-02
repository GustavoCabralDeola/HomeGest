package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config armazena as configurações da aplicação carregadas de variáveis de ambiente.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

// Load carrega as configurações a partir do arquivo .env e variáveis de ambiente.
func Load() (*Config, error) {
	// Tenta carregar o .env, mas não falha se o arquivo não existir
	_ = godotenv.Load()

	cfg := &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "1433"),
		DBUser:     getEnv("DB_USER", "sa"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "GestHome"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}

	if cfg.DBPassword == "" {
		return nil, fmt.Errorf("DB_PASSWORD é obrigatório")
	}

	return cfg, nil
}

// DSN retorna a string de conexão formatada para SQL Server.
func (c *Config) DSN() string {
	return fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
	)
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
