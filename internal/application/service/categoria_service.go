package service

import (
	"gesthome/internal/apperror"
	"gesthome/internal/application/dto"
	"gesthome/internal/domain/entity"
	"gesthome/internal/domain/repository"
	"strings"
)

type CategoriaService struct {
	repo repository.CategoriaRepository
}

func NewCategoriaService(repo repository.CategoriaRepository) *CategoriaService {
	return &CategoriaService{repo: repo}
}

func (s *CategoriaService) ObterTodasCategorias() ([]dto.CategoriaResponseDTO, error) {
	categorias, err := s.repo.ObterTodasAsCategorias()
	if err != nil {
		return nil, err
	}

	var response []dto.CategoriaResponseDTO
	for _, c := range categorias {
		response = append(response, dto.CategoriaResponseDTO{
			ID:                  c.ID,
			Descricao:           c.Descricao,
			CategoriaFinalidade: c.Finalidade,
		})
	}

	return response, nil
}

func (s *CategoriaService) ObterCategoriaPorId(id int) (*dto.CategoriaResponseDTO, error) {
	categoria, err := s.repo.ObterCategoriaPorId(id)
	if err != nil {
		return nil, err
	}
	if categoria == nil {
		return nil, nil
	}

	return &dto.CategoriaResponseDTO{
		ID:                  categoria.ID,
		Descricao:           categoria.Descricao,
		CategoriaFinalidade: categoria.Finalidade,
	}, nil
}

func (s *CategoriaService) AdicionarCategoria(input dto.CreateCategoriaDTO) (*dto.CategoriaResponseDTO, error) {
	if strings.TrimSpace(input.Descricao) == "" {
		return nil, apperror.NewValidation("A descrição da categoria é obrigatória.")
	}
	if len(input.Descricao) > 400 {
		return nil, apperror.NewValidation("A descrição da categoria não pode exceder 400 caracteres.")
	}
	if !input.CategoriaFinalidade.IsValid() {
		return nil, apperror.NewValidation("Finalidade da categoria inválida.")
	}

	novaCategoria := &entity.Categoria{
		Descricao:  input.Descricao,
		Finalidade: input.CategoriaFinalidade,
	}

	if err := s.repo.AdicionarCategoria(novaCategoria); err != nil {
		return nil, err
	}

	return &dto.CategoriaResponseDTO{
		ID:                  novaCategoria.ID,
		Descricao:           novaCategoria.Descricao,
		CategoriaFinalidade: novaCategoria.Finalidade,
	}, nil
}

func (s *CategoriaService) ObterTotaisPorCategoria() ([]dto.TotaisPorCategoriaDTO, error) {
	totais, err := s.repo.ObterTotaisPorCategoria()
	if err != nil {
		return nil, err
	}

	var response []dto.TotaisPorCategoriaDTO
	for _, t := range totais {
		response = append(response, dto.TotaisPorCategoriaDTO{
			Descricao:     t.Descricao,
			TotalReceitas: t.TotalReceitas,
			TotalDespesas: t.TotalDespesas,
			Saldo:         t.TotalReceitas - t.TotalDespesas,
		})
	}

	return response, nil
}

func (s *CategoriaService) ObterResumoCategorias() (*dto.ResumoCategoriasDTO, error) {
	totaisDto, err := s.ObterTotaisPorCategoria()
	if err != nil {
		return nil, err
	}

	var totalGeralReceitas, totalGeralDespesas float64
	for _, t := range totaisDto {
		totalGeralReceitas += t.TotalReceitas
		totalGeralDespesas += t.TotalDespesas
	}

	return &dto.ResumoCategoriasDTO{
		Categorias:         totaisDto,
		TotalGeralReceitas: totalGeralReceitas,
		TotalGeralDespesas: totalGeralDespesas,
		SaldoLiquido:       totalGeralReceitas - totalGeralDespesas,
	}, nil
}
