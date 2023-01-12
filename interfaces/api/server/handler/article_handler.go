package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/shake551/go-sql-limit-sample/application/usecase"
)

type ArticleHandler interface {
	GetLimit(w http.ResponseWriter, r *http.Request)
}

type articleHandler struct {
	u usecase.ArticleUseCase
}

func NewArticleHandler(u usecase.ArticleUseCase) ArticleHandler {
	return &articleHandler{u}
}

func (h *articleHandler) GetLimit(w http.ResponseWriter, r *http.Request) {
	v := r.URL.Query()
	if v == nil {
		return
	}

	var limit = int64(30)
	if v.Get("limit") != "" {
		l, err := strconv.ParseInt(v.Get("limit"), 10, 64)
		if err != nil {
			log.Printf("failed to get limit. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		limit = l
	}

	var offset = int64(0)
	if v.Get("offset") != "" {
		o, err := strconv.ParseInt(v.Get("offset"), 10, 64)
		if err != nil {
			log.Printf("failed to get offset. err: %v", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		offset = o
	}

	cocktails, err := h.u.GetLimit(r.Context(), limit, offset)
	if err != nil {
		log.Printf("failed to get cocktails. err: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(cocktails)
	if err != nil {
		log.Printf("failed to parse json. err: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(b)))
	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
