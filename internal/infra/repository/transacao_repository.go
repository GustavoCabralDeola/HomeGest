package repository

import (
	"errors"
	"gesthome/internal/domain/entity"

	"gorm.io/gorm"
)

type TransacaoRepositoryGorm struct {
	db *gorm.DB
}

func NewTransacaoRepository(db *gorm.DB) *TransacaoRepositoryGorm {
	return &TransacaoRepositoryGorm{db: db}
}

func (r *TransacaoRepositoryGorm) AdicionarTransacao(t *entity.Transacao) error {
	return r.db.Create(t).Error
}

func (r *TransacaoRepositoryGorm) ObterTodasAsTransacoes() ([]entity.Transacao, error) {
	var transacoes []entity.Transacao
	if err := r.db.Preload("Categoria").Preload("Pessoa").Find(&transacoes).Error; err != nil {
		return nil, err
	}
	return transacoes, nil
}

func (r *TransacaoRepositoryGorm) ObterTransacaoPorId(id int) (*entity.Transacao, error) {
	var t entity.Transacao
	err := r.db.Preload("Categoria").Preload("Pessoa").First(&t, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *TransacaoRepositoryGorm) ObterCategoriaIdNaTransacao(id int) (*entity.Categoria, error) {
	var cat entity.Categoria
	err := r.db.First(&cat, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

func (r *TransacaoRepositoryGorm) ObterPessoaIdNaTransacao(id int) (*entity.Pessoa, error) {
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
