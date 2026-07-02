package entity

import "gesthome/internal/domain/enum"

// Categoria representa uma categoria de transação financeira.
// Possui uma finalidade que define se aceita Receita, Despesa ou Ambas.
// Relacionamento 1:N com Transacao.
type Categoria struct {
	ID         int                        `gorm:"primaryKey;autoIncrement" json:"id"`
	Descricao  string                     `gorm:"type:nvarchar(400);not null" json:"descricao"`
	Finalidade enum.CategoriaFinalidade   `gorm:"type:int;not null" json:"finalidade"`
	Transacoes []Transacao                `gorm:"foreignKey:CategoriaId" json:"-"`
}

// TableName define o nome da tabela no banco de dados.
func (Categoria) TableName() string {
	return "Categorias"
}
