package dto

import "gesthome/internal/domain/enum"

type CreateCategoriaDTO struct {
	Descricao           string                   `json:"descricao"`
	CategoriaFinalidade enum.CategoriaFinalidade `json:"categoriaFinalidade"`
}

type CategoriaResponseDTO struct {
	ID                  int                      `json:"id"`
	Descricao           string                   `json:"descricao"`
	CategoriaFinalidade enum.CategoriaFinalidade `json:"categoriaFinalidade"`
}

type TotaisPorCategoriaDTO struct {
	Descricao     string  `json:"descricao"`
	TotalReceitas float64 `json:"totalReceitas"`
	TotalDespesas float64 `json:"totalDespesas"`
	Saldo         float64 `json:"saldo"`
}

type ResumoCategoriasDTO struct {
	Categorias         []TotaisPorCategoriaDTO `json:"categorias"`
	TotalGeralReceitas float64                 `json:"totalGeralReceitas"`
	TotalGeralDespesas float64                 `json:"totalGeralDespesas"`
	SaldoLiquido       float64                 `json:"saldoLiquido"`
}
