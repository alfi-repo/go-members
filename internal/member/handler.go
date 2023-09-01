package member

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/google/uuid"

	"github.com/go-chi/chi/v5"
)

var (
	errInvalidID = errors.New("invalid id")
)

func WriteErrorResponse(w http.ResponseWriter, httpCode int, message string) {
	w.WriteHeader(httpCode)
	_ = json.NewEncoder(w).Encode(message)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) Router() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", h.list)
	r.Post("/", h.new)
	r.Get("/{id}", h.findByID)
	r.Put("/{id}", h.edit)
	r.Delete("/{id}", h.delete)
	return r
}

func (h Handler) list(w http.ResponseWriter, _ *http.Request) {
	res, err := h.service.list()
	if err != nil {
		log.Err(err)
		WriteErrorResponse(w, http.StatusInternalServerError, "server error")
		return
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func (h Handler) new(w http.ResponseWriter, r *http.Request) {
	var req newRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "check your request body")
		return
	}

	res, err := h.service.new(req)
	if err != nil {
		switch {
		case errors.Is(err, ErrMemberValidation):
			WriteErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
			return
		default:
			log.Err(err)
			WriteErrorResponse(w, http.StatusInternalServerError, "server error")
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}

func (h Handler) findByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errInvalidID.Error())
		return
	}

	req := findRequest{ID: id}
	res, err := h.service.findByID(req)
	if err != nil {
		switch {
		case errors.Is(err, ErrMemberNotFound):
			WriteErrorResponse(w, http.StatusNotFound, ErrMemberNotFound.Error())
			return
		default:
			log.Err(err)
			WriteErrorResponse(w, http.StatusInternalServerError, "server error")
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func (h Handler) edit(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errInvalidID.Error())
		return
	}

	var req editRequest
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, "check your request body")
		return
	}
	req.ID = id

	res, err := h.service.edit(req)
	if err != nil {
		switch {
		case errors.Is(err, ErrMemberNotFound):
			WriteErrorResponse(w, http.StatusNotFound, ErrMemberNotFound.Error())
			return
		case errors.Is(err, ErrMemberAlreadyVerified):
			WriteErrorResponse(w, http.StatusBadRequest, ErrMemberAlreadyVerified.Error())
			return
		case errors.Is(err, ErrMemberValidation):
			WriteErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
			return
		default:
			log.Err(err)
			WriteErrorResponse(w, http.StatusInternalServerError, "server error")
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}

func (h Handler) delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest, errInvalidID.Error())
		return
	}

	req := findRequest{ID: id}
	err = h.service.delete(req)
	if err != nil {
		switch {
		case errors.Is(err, ErrMemberNotFound):
			WriteErrorResponse(w, http.StatusNotFound, ErrMemberNotFound.Error())
			return
		default:
			log.Err(err)
			WriteErrorResponse(w, http.StatusInternalServerError, "server error")
			return
		}
	}

	w.WriteHeader(http.StatusNoContent)
}
