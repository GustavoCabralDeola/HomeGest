package main

import (
	"fmt"
	"log"
	"net/http"

	"gesthome/internal/api/handler"
	"gesthome/internal/api/router"
	"gesthome/internal/application/service"
	"gesthome/internal/config"
	"gesthome/internal/infra/database"
	infraRepo "gesthome/internal/infra/repository"

	"github.com/rs/cors"
)

func main() {
	// Carrega configuração
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	// Conecta ao banco de dados
	db, err := database.NewConnection(cfg.DSN())
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	// AutoMigrate
	if err := database.AutoMigrate(db); err != nil {
		log.Fatalf("Erro na migração: %v", err)
	}

	// Repositories
	categoriaRepo := infraRepo.NewCategoriaRepository(db)
	pessoaRepo := infraRepo.NewPessoaRepository(db)
	transacaoRepo := infraRepo.NewTransacaoRepository(db)

	// Services
	categoriaSvc := service.NewCategoriaService(categoriaRepo)
	pessoaSvc := service.NewPessoaService(pessoaRepo)
	transacaoSvc := service.NewTransacaoService(transacaoRepo)

	// Handlers
	categoriaHandler := handler.NewCategoriaHandler(categoriaSvc)
	pessoaHandler := handler.NewPessoaHandler(pessoaSvc)
	transacaoHandler := handler.NewTransacaoHandler(transacaoSvc)

	// Router
	r := router.NewRouter(categoriaHandler, pessoaHandler, transacaoHandler)

	// CORS (mesma config do C# - permite frontend Vite)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
			"http://localhost:5174",
			"http://localhost:5175",
			"https://localhost:5173",
			"https://localhost:5174",
			"https://localhost:5175",
		},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
	})

	corsHandler := c.Handler(r)

	// Inicia servidor
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Servidor GestHome iniciando em %s", addr)
	if err := http.ListenAndServe(addr, corsHandler); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
