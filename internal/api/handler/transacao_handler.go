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

type TransacaoHandler struct {
	service *service.TransacaoService
}

func NewTransacaoHandler(s *service.TransacaoService) *TransacaoHandler {
	return &TransacaoHandler{service: s}
}

func (h *TransacaoHandler) ObterTodasAsTransacoes(w http.ResponseWriter, r *http.Request) {
	transacoes, err := h.service.ObterTodasAsTransacoes()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, transacoes)
}

func (h *TransacaoHandler) ObterTransacaoPorId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	transacao, err := h.service.ObterTransacaoPorId(id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if transacao == nil {
		respondError(w, http.StatusNotFound, "Transação não encontrada")
		return
	}

	respondJSON(w, http.StatusOK, transacao)
}

func (h *TransacaoHandler) AdicionarTransacao(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateTransacaoDTO
	if err := readJSON(r, &input); err != nil {
		respondError(w, http.StatusBadRequest, "Payload inválido")
		return
	}

	transacao, err := h.service.AdicionarTransacao(input)
	if err != nil {
		if apperror.IsNotFound(err) {
			respondError(w, http.StatusNotFound, err.Error())
		} else if apperror.IsValidation(err) {
			respondError(w, http.StatusBadRequest, err.Error())
		} else {
			respondError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/api/transacao/obter-transacao/%d", transacao.ID))
	respondJSON(w, http.StatusCreated, transacao)
}
