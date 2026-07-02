package dto

import "gesthome/internal/domain/enum"

type CreateTransacaoDTO struct {
	Descricao     string             `json:"descricao"`
	Valor         float64            `json:"valor"`
	TipoTransacao enum.TipoTransacao `json:"tipoTransacao"`
	CategoriaId   int                `json:"categoriaId"`
	PessoaId      int                `json:"pessoaId"`
}

type TransacaoResponseDTO struct {
	ID                  int                      `json:"id"`
	Descricao           string                   `json:"descricao"`
	Valor               float64                  `json:"valor"`
	TipoTransacao       enum.TipoTransacao       `json:"tipoTransacao"`
	PessoaNome          string                   `json:"pessoaNome"`
	CategoriaDescricao  string                   `json:"categoriaDescricao"`
	CategoriaFinalidade enum.CategoriaFinalidade `json:"categoriaFinalidade"`
	PessoaId            int                      `json:"pessoaId"`
	CategoriaId         int                      `json:"categoriaId"`
}

type TotaisTransacaoDTO struct {
	TotalReceita float64 `json:"totalReceita"`
	TotalDespesa float64 `json:"totalDespesa"`
	Saldo        float64 `json:"saldo"`
}
