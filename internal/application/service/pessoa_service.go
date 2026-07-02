package service

import (
	"fmt"
	"gesthome/internal/apperror"
	"gesthome/internal/application/dto"
	"gesthome/internal/domain/entity"
	"gesthome/internal/domain/repository"
	"strings"
)

type PessoaService struct {
	repo repository.PessoaRepository
}

func NewPessoaService(repo repository.PessoaRepository) *PessoaService {
	return &PessoaService{repo: repo}
}

func (s *PessoaService) ObterTodasPessoas() ([]dto.PessoaResponseDTO, error) {
	pessoas, err := s.repo.ObterTodosOsRegistros()
	if err != nil {
		return nil, err
	}

	var response []dto.PessoaResponseDTO
	for _, p := range pessoas {
		response = append(response, dto.PessoaResponseDTO{
			ID:    p.ID,
			Nome:  p.Nome,
			Idade: p.Idade,
		})
	}

	return response, nil
}

func (s *PessoaService) ObterPessoaPorId(id int) (*dto.PessoaResponseDTO, error) {
	pessoa, err := s.repo.ObterRegistrosPorId(id)
	if err != nil {
		return nil, err
	}
	if pessoa == nil {
		return nil, apperror.NewNotFound(fmt.Sprintf("Pessoa com ID %d não encontrada.", id))
	}

	return &dto.PessoaResponseDTO{
		ID:    pessoa.ID,
		Nome:  pessoa.Nome,
		Idade: pessoa.Idade,
	}, nil
}

func (s *PessoaService) ObterResumoPessoas() (*dto.ResumoPessoasDTO, error) {
	dados, err := s.repo.ObterTotaisPorPessoa()
	if err != nil {
		return nil, err
	}

	var pessoasDto []dto.TotaisPorPessoaDTO
	var totalGeralReceitas, totalGeralDespesas float64

	for _, p := range dados {
		pessoasDto = append(pessoasDto, dto.TotaisPorPessoaDTO{
			NomePessoa:    p.NomePessoa,
			TotalReceitas: p.TotalReceitas,
			TotalDespesas: p.TotalDespesas,
			Saldo:         p.TotalReceitas - p.TotalDespesas,
		})
		totalGeralReceitas += p.TotalReceitas
		totalGeralDespesas += p.TotalDespesas
	}

	return &dto.ResumoPessoasDTO{
		Pessoas:            pessoasDto,
		TotalGeralReceitas: totalGeralReceitas,
		TotalGeralDespesas: totalGeralDespesas,
		SaldoLiquido:       totalGeralReceitas - totalGeralDespesas,
	}, nil
}

func (s *PessoaService) AdicionarPessoa(input dto.CreatePessoaDTO) (*dto.PessoaResponseDTO, error) {
	if strings.TrimSpace(input.Nome) == "" {
		return nil, apperror.NewValidation("O nome da pessoa é obrigatório.")
	}
	if len(input.Nome) > 200 {
		return nil, apperror.NewValidation("Nome deve ter no máximo 200 caracteres.")
	}

	novaPessoa := &entity.Pessoa{
		Nome:  input.Nome,
		Idade: input.Idade,
	}

	if err := s.repo.AdicionarPessoa(novaPessoa); err != nil {
		return nil, err
	}

	return &dto.PessoaResponseDTO{
		ID:    novaPessoa.ID,
		Nome:  novaPessoa.Nome,
		Idade: novaPessoa.Idade,
	}, nil
}

func (s *PessoaService) AtualizarDadosPessoa(id int, input dto.UpdatePessoaDTO) error {
	pessoaExistente, err := s.repo.ObterRegistrosPorId(id)
	if err != nil {
		return err
	}
	if pessoaExistente == nil {
		return apperror.NewNotFound(fmt.Sprintf("Pessoa com ID %d não encontrada.", id))
	}

	if strings.TrimSpace(input.Nome) == "" {
		return apperror.NewValidation("O nome da pessoa é obrigatório.")
	}
	if len(input.Nome) > 200 {
		return apperror.NewValidation("Nome deve ter no máximo 200 caracteres.")
	}

	pessoaExistente.Nome = input.Nome
	pessoaExistente.Idade = input.Idade

	return s.repo.AtualizarRegistro(pessoaExistente)
}

func (s *PessoaService) ExcluirPessoa(id int) error {
	if id <= 0 {
		return apperror.NewValidation("Id inválido")
	}

	pessoaExistente, err := s.repo.ObterRegistrosPorId(id)
	if err != nil {
		return err
	}
	if pessoaExistente == nil {
		return apperror.NewNotFound(fmt.Sprintf("Pessoa com ID %d não encontrada.", id))
	}

	return s.repo.DeletarRegistro(id)
}
