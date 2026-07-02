package dto

type CreatePessoaDTO struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

type UpdatePessoaDTO struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

type PessoaResponseDTO struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
}

type TotaisPorPessoaDTO struct {
	NomePessoa    string  `json:"nomePessoa"`
	TotalReceitas float64 `json:"totalReceitas"`
	TotalDespesas float64 `json:"totalDespesas"`
	Saldo         float64 `json:"saldo"`
}

type ResumoPessoasDTO struct {
	Pessoas            []TotaisPorPessoaDTO `json:"pessoas"`
	TotalGeralReceitas float64              `json:"totalGeralReceitas"`
	TotalGeralDespesas float64              `json:"totalGeralDespesas"`
	SaldoLiquido       float64              `json:"saldoLiquido"`
}
