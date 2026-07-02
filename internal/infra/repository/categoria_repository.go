package repository

import (
	"errors"
	"gesthome/internal/domain/entity"

	"gorm.io/gorm"
)

type CategoriaRepositoryGorm struct {
	db *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) *CategoriaRepositoryGorm {
	return &CategoriaRepositoryGorm{db: db}
}

func (r *CategoriaRepositoryGorm) ObterTodasAsCategorias() ([]entity.Categoria, error) {
	var categorias []entity.Categoria
	if err := r.db.Find(&categorias).Error; err != nil {
		return nil, err
	}
	return categorias, nil
}

func (r *CategoriaRepositoryGorm) ObterCategoriaPorId(id int) (*entity.Categoria, error) {
	var categoria entity.Categoria
	err := r.db.First(&categoria, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &categoria, nil
}

func (r *CategoriaRepositoryGorm) AdicionarCategoria(categoria *entity.Categoria) error {
	return r.db.Create(categoria).Error
}

func (r *CategoriaRepositoryGorm) ObterTotaisPorCategoria() ([]entity.TotaisPorCategoria, error) {
	var totais []entity.TotaisPorCategoria
	query := `
		SELECT 
			c.Descricao as descricao,
			ISNULL(SUM(CASE WHEN t.TipoTransacao = 1 THEN t.Valor ELSE 0 END), 0) as total_receitas,
			ISNULL(SUM(CASE WHEN t.TipoTransacao = 2 THEN t.Valor ELSE 0 END), 0) as total_despesas
		FROM Categorias c
		LEFT JOIN Transacoes t ON c.Id = t.CategoriaId
		GROUP BY c.Descricao
	`
	if err := r.db.Raw(query).Scan(&totais).Error; err != nil {
		return nil, err
	}
	return totais, nil
}
