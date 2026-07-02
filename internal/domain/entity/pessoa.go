package entity

// Pessoa representa uma pessoa da família/casa.
// Relacionamento 1:N com Transacao (cascade delete).
type Pessoa struct {
	ID         int         `gorm:"primaryKey;autoIncrement" json:"id"`
	Nome       string      `gorm:"type:nvarchar(200);not null" json:"nome"`
	Idade      int         `gorm:"type:int;not null" json:"idade"`
	Transacoes []Transacao `gorm:"foreignKey:PessoaId;constraint:OnDelete:CASCADE" json:"-"`
}

// TableName define o nome da tabela no banco de dados.
func (Pessoa) TableName() string {
	return "Pessoas"
}
