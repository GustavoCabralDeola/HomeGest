package repository

import (
	"errors"
	"gesthome/internal/domain/entity"

	"gorm.io/gorm"
)

type PessoaRepositoryGorm struct {
	db *gorm.DB
}

func NewPessoaRepository(db *gorm.DB) *PessoaRepositoryGorm {
	return &PessoaRepositoryGorm{db: db}
}

func (r *PessoaRepositoryGorm) ObterRegistrosPorId(id int) (*entity.Pessoa, error) {
	var pessoa entity.Pessoa
	err := r.db.First(&pessoa, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &pessoa, nil
}

func (r *PessoaRepositoryGorm) ObterTodosOsRegistros() ([]entity.Pessoa, error) {
	var pessoas []entity.Pessoa
	if err := r.db.Find(&pessoas).Error; err != nil {
		return nil, err
	}
	return pessoas, nil
}

func (r *PessoaRepositoryGorm) ObterTotaisPorPessoa() ([]entity.TotaisPorPessoa, error) {
	var totais []entity.TotaisPorPessoa
	query := `
		SELECT 
			p.Nome as nome_pessoa,
			ISNULL(SUM(CASE WHEN t.TipoTransacao = 1 THEN t.Valor ELSE 0 END), 0) as total_receitas,
			ISNULL(SUM(CASE WHEN t.TipoTransacao = 2 THEN t.Valor ELSE 0 END), 0) as total_despesas
		FROM Pessoas p
		LEFT JOIN Transacoes t ON p.Id = t.PessoaId
		GROUP BY p.Nome
	`
	if err := r.db.Raw(query).Scan(&totais).Error; err != nil {
		return nil, err
	}
	return totais, nil
}

func (r *PessoaRepositoryGorm) AdicionarPessoa(pessoa *entity.Pessoa) error {
	return r.db.Create(pessoa).Error
}

func (r *PessoaRepositoryGorm) AtualizarRegistro(pessoa *entity.Pessoa) error {
	return r.db.Save(pessoa).Error
}

func (r *PessoaRepositoryGorm) DeletarRegistro(id int) error {
	return r.db.Delete(&entity.Pessoa{}, id).Error
}
