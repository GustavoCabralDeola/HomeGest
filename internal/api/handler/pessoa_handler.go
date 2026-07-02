package handler

import (
	"fmt"
	"gesthome/internal/apperror"
	"gesthome/internal/application/dto"
	"gesthome/internal/application/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PessoaHandler struct {
	service *service.PessoaService
}

func NewPessoaHandler(s *service.PessoaService) *PessoaHandler {
	return &PessoaHandler{service: s}
}

func (h *PessoaHandler) ObterTodasPessoas(w http.ResponseWriter, r *http.Request) {
	pessoas, err := h.service.ObterTodasPessoas()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, pessoas)
}

func (h *PessoaHandler) ObterPessoaPorId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	pessoa, err := h.service.ObterPessoaPorId(id)
	if err != nil {
		if apperror.IsNotFound(err) {
			respondError(w, http.StatusNotFound, err.Error())
		} else {
			respondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondJSON(w, http.StatusOK, pessoa)
}

func (h *PessoaHandler) ObterResumoPessoas(w http.ResponseWriter, r *http.Request) {
	resumo, err := h.service.ObterResumoPessoas()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, resumo)
}

func (h *PessoaHandler) AdicionarPessoa(w http.ResponseWriter, r *http.Request) {
	var input dto.CreatePessoaDTO
	if err := readJSON(r, &input); err != nil {
		respondError(w, http.StatusBadRequest, "Payload inválido")
		return
	}

	pessoa, err := h.service.AdicionarPessoa(input)
	if err != nil {
		if apperror.IsValidation(err) {
			respondError(w, http.StatusBadRequest, err.Error())
		} else {
			respondError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/api/pessoa/%d", pessoa.ID))
	respondJSON(w, http.StatusCreated, pessoa)
}

func (h *PessoaHandler) AtualizarPessoa(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	var input dto.UpdatePessoaDTO
	if err := readJSON(r, &input); err != nil {
		respondError(w, http.StatusBadRequest, "Payload inválido")
		return
	}

	if err := h.service.AtualizarDadosPessoa(id, input); err != nil {
		if apperror.IsNotFound(err) {
			respondError(w, http.StatusNotFound, err.Error())
		} else if apperror.IsValidation(err) {
			respondError(w, http.StatusBadRequest, err.Error())
		} else {
			respondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *PessoaHandler) ExcluirPessoa(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.service.ExcluirPessoa(id); err != nil {
		if apperror.IsNotFound(err) {
			respondError(w, http.StatusNotFound, err.Error())
		} else {
			respondError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
