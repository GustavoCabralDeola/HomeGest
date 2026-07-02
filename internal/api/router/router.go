package router

import (
	"gesthome/internal/api/handler"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// NewRouter configura e retorna o router com todas as rotas.
func NewRouter(
	categoriaHandler *handler.CategoriaHandler,
	pessoaHandler *handler.PessoaHandler,
	transacaoHandler *handler.TransacaoHandler,
) chi.Router {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.AllowContentType("application/json"))

	// Rotas de Categoria
	r.Route("/api/categoria", func(r chi.Router) {
		r.Get("/obter-categorias", categoriaHandler.ObterTodasCategorias)
		r.Get("/resumocategoria", categoriaHandler.ObterResumoCategorias)
		r.Get("/{id}", categoriaHandler.ObterCategoriaPorId)
		r.Post("/", categoriaHandler.AdicionarCategoria)
	})

	// Rotas de Pessoa
	r.Route("/api/pessoa", func(r chi.Router) {
		r.Get("/obter-pessoas", pessoaHandler.ObterTodasPessoas)
		r.Get("/resumopessoas", pessoaHandler.ObterResumoPessoas)
		r.Get("/{id}", pessoaHandler.ObterPessoaPorId)
		r.Post("/", pessoaHandler.AdicionarPessoa)
		r.Put("/{id}", pessoaHandler.AtualizarPessoa)
		r.Delete("/{id}", pessoaHandler.ExcluirPessoa)
	})

	// Rotas de Transacao
	r.Route("/api/transacao", func(r chi.Router) {
		r.Get("/obter-transacoes", transacaoHandler.ObterTodasAsTransacoes)
		r.Get("/obter-transacao/{id}", transacaoHandler.ObterTransacaoPorId)
		r.Post("/adicionar-transacao", transacaoHandler.AdicionarTransacao)
	})

	return r
}
