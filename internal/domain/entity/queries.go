package entity

// TotaisPorCategoria armazena os totais agregados de receitas e despesas por categoria.
// Usado para consultas de resumo/relatório.
type TotaisPorCategoria struct {
	Descricao     string  `json:"descricao"`
	TotalReceitas float64 `json:"totalReceitas"`
	TotalDespesas float64 `json:"totalDespesas"`
}

// TotaisPorPessoa armazena os totais agregados de receitas e despesas por pessoa.
// Usado para consultas de resumo/relatório.
type TotaisPorPessoa struct {
	NomePessoa    string  `json:"nomePessoa"`
	TotalReceitas float64 `json:"totalReceitas"`
	TotalDespesas float64 `json:"totalDespesas"`
}
