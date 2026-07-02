package service

import (
	"fmt"
	"gesthome/internal/apperror"
	"gesthome/internal/application/dto"
	"gesthome/internal/domain/entity"
	"gesthome/internal/domain/enum"
	"gesthome/internal/domain/repository"
)

type TransacaoService struct {
	repo repository.TransacaoRepository
}

func NewTransacaoService(repo repository.TransacaoRepository) *TransacaoService {
	return &TransacaoService{repo: repo}
}

func (s *TransacaoService) ObterTodasAsTransacoes() ([]dto.TransacaoResponseDTO, error) {
	transacoes, err := s.repo.ObterTodasAsTransacoes()
	if err != nil {
		return nil, err
	}

	var response []dto.TransacaoResponseDTO
	for _, t := range transacoes {
		response = append(response, dto.TransacaoResponseDTO{
			ID:                  t.ID,
			Descricao:           t.Descricao,
			Valor:               t.Valor,
			TipoTransacao:       t.TipoTransacao,
			CategoriaId:         t.CategoriaId,
			CategoriaDescricao:  t.Categoria.Descricao,
			CategoriaFinalidade: t.Categoria.Finalidade,
			PessoaId:            t.PessoaId,
			PessoaNome:          t.Pessoa.Nome,
		})
	}

	return response, nil
}

func (s *TransacaoService) ObterTransacaoPorId(id int) (*dto.TransacaoResponseDTO, error) {
	t, err := s.repo.ObterTransacaoPorId(id)
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, nil
	}

	return &dto.TransacaoResponseDTO{
		ID:                  t.ID,
		Descricao:           t.Descricao,
		Valor:               t.Valor,
		TipoTransacao:       t.TipoTransacao,
		CategoriaId:         t.CategoriaId,
		CategoriaDescricao:  t.Categoria.Descricao,
		CategoriaFinalidade: t.Categoria.Finalidade,
		PessoaId:            t.PessoaId,
		PessoaNome:          t.Pessoa.Nome,
	}, nil
}

func (s *TransacaoService) AdicionarTransacao(input dto.CreateTransacaoDTO) (*dto.TransacaoResponseDTO, error) {
	cat, pes, err := s.validaTransacao(input)
	if err != nil {
		return nil, err
	}

	novaTransacao := &entity.Transacao{
		Descricao:           input.Descricao,
		Valor:               input.Valor,
		TipoTransacao:       input.TipoTransacao,
		CategoriaFinalidade: cat.Finalidade,
		CategoriaId:         input.CategoriaId,
		PessoaId:            input.PessoaId,
	}

	if err := s.repo.AdicionarTransacao(novaTransacao); err != nil {
		return nil, err
	}

	return &dto.TransacaoResponseDTO{
		ID:                  novaTransacao.ID,
		Descricao:           novaTransacao.Descricao,
		Valor:               novaTransacao.Valor,
		TipoTransacao:       novaTransacao.TipoTransacao,
		CategoriaDescricao:  cat.Descricao,
		CategoriaFinalidade: cat.Finalidade,
		CategoriaId:         cat.ID,
		PessoaNome:          pes.Nome,
		PessoaId:            pes.ID,
	}, nil
}

func (s *TransacaoService) validaTransacao(input dto.CreateTransacaoDTO) (*entity.Categoria, *entity.Pessoa, error) {
	cat, err := s.repo.ObterCategoriaIdNaTransacao(input.CategoriaId)
	if err != nil {
		return nil, nil, err
	}
	if cat == nil {
		return nil, nil, apperror.NewNotFound(fmt.Sprintf("CategoriaId %d não encontrada.", input.CategoriaId))
	}

	pes, err := s.repo.ObterPessoaIdNaTransacao(input.PessoaId)
	if err != nil {
		return nil, nil, err
	}
	if pes == nil {
		return nil, nil, apperror.NewNotFound(fmt.Sprintf("PessoaId %d não encontrada.", input.PessoaId))
	}

	if len(input.Descricao) > 400 {
		return nil, nil, apperror.NewValidation("A descrição da transação não deve passar mais de 400 caracteres")
	}

	if input.Valor <= 0 {
		return nil, nil, apperror.NewValidation("O valor da transação deve ser maior que zero e positivo")
	}

	if pes.Idade < 18 && input.TipoTransacao == enum.TipoTransacaoReceita {
		return nil, nil, apperror.NewValidation("Menor de idade só pode possuir despesas.")
	}

	if int(cat.Finalidade) != int(input.TipoTransacao) && cat.Finalidade != enum.CategoriaFinalidadeAmbas {
		return nil, nil, apperror.NewValidation(fmt.Sprintf(
			"Transação inválida. A transação é do tipo %v e não pode ser usada em uma categoria do tipo %v.",
			input.TipoTransacao, cat.Finalidade,
		))
	}

	return cat, pes, nil
}
