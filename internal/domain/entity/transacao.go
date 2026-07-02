package entity

import "gesthome/internal/domain/enum"

// Transacao representa uma transação financeira (receita ou despesa).
// Possui chaves estrangeiras para Categoria e Pessoa.
type Transacao struct {
	ID                  int                        `gorm:"primaryKey;autoIncrement" json:"id"`
	Descricao           string                     `gorm:"type:nvarchar(400);not null" json:"descricao"`
	Valor               float64                    `gorm:"type:decimal(18,2);not null" json:"valor"`
	TipoTransacao       enum.TipoTransacao         `gorm:"type:int;not null" json:"tipoTransacao"`
	CategoriaFinalidade enum.CategoriaFinalidade   `gorm:"type:int" json:"categoriaFinalidade"`

	// Chaves estrangeiras explícitas
	CategoriaId int `gorm:"not null" json:"categoriaId"`
	PessoaId    int `gorm:"not null" json:"pessoaId"`

	// Propriedades de navegação
	Categoria Categoria `gorm:"foreignKey:CategoriaId;constraint:OnDelete:RESTRICT" json:"-"`
	Pessoa    Pessoa    `gorm:"foreignKey:PessoaId;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName define o nome da tabela no banco de dados.
func (Transacao) TableName() string {
	return "Transacoes"
}
