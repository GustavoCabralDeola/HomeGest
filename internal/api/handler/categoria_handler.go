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

type CategoriaHandler struct {
	service *service.CategoriaService
}

func NewCategoriaHandler(s *service.CategoriaService) *CategoriaHandler {
	return &CategoriaHandler{service: s}
}

func (h *CategoriaHandler) ObterTodasCategorias(w http.ResponseWriter, r *http.Request) {
	categorias, err := h.service.ObterTodasCategorias()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, categorias)
}

func (h *CategoriaHandler) ObterCategoriaPorId(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		respondError(w, http.StatusBadRequest, "ID inválido")
		return
	}

	categoria, err := h.service.ObterCategoriaPorId(id)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if categoria == nil {
		respondError(w, http.StatusNotFound, fmt.Sprintf("Categoria com ID %d não encontrada.", id))
		return
	}

	respondJSON(w, http.StatusOK, categoria)
}

func (h *CategoriaHandler) ObterResumoCategorias(w http.ResponseWriter, r *http.Request) {
	resumo, err := h.service.ObterResumoCategorias()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, resumo)
}

func (h *CategoriaHandler) AdicionarCategoria(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateCategoriaDTO
	if err := readJSON(r, &input); err != nil {
		respondError(w, http.StatusBadRequest, "Payload inválido")
		return
	}

	categoria, err := h.service.AdicionarCategoria(input)
	if err != nil {
		if apperror.IsValidation(err) {
			respondError(w, http.StatusBadRequest, err.Error())
		} else if apperror.IsNotFound(err) {
			respondError(w, http.StatusNotFound, err.Error())
		} else {
			respondError(w, http.StatusBadRequest, err.Error())
		}
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/api/categoria/%d", categoria.ID))
	respondJSON(w, http.StatusCreated, categoria)
}
