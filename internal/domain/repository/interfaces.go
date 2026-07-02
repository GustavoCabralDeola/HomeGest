package repository

import "gesthome/internal/domain/entity"

// CategoriaRepository define as operações de acesso a dados para Categoria.
type CategoriaRepository interface {
	ObterTodasAsCategorias() ([]entity.Categoria, error)
	ObterCategoriaPorId(id int) (*entity.Categoria, error)
	AdicionarCategoria(categoria *entity.Categoria) error
	ObterTotaisPorCategoria() ([]entity.TotaisPorCategoria, error)
}

// PessoaRepository define as operações de acesso a dados para Pessoa.
type PessoaRepository interface {
	ObterRegistrosPorId(id int) (*entity.Pessoa, error)
	ObterTodosOsRegistros() ([]entity.Pessoa, error)
	ObterTotaisPorPessoa() ([]entity.TotaisPorPessoa, error)
	AdicionarPessoa(pessoa *entity.Pessoa) error
	AtualizarRegistro(pessoa *entity.Pessoa) error
	DeletarRegistro(id int) error
}

// TransacaoRepository define as operações de acesso a dados para Transacao.
type TransacaoRepository interface {
	AdicionarTransacao(transacao *entity.Transacao) error
	ObterTodasAsTransacoes() ([]entity.Transacao, error)
	ObterTransacaoPorId(id int) (*entity.Transacao, error)
	ObterCategoriaIdNaTransacao(id int) (*entity.Categoria, error)
	ObterPessoaIdNaTransacao(id int) (*entity.Pessoa, error)
}
