package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"publisher/models"
	"publisher/usecase"
	"strconv"
)

type Handler struct {
	ProductUsecase usecase.ProductUsecase
}

func NewHandler(productUsecase usecase.ProductUsecase) Handler {
	return Handler{productUsecase}
}

func (h *Handler) Store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "POST" {
		product := models.Product{}
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		dec := json.NewDecoder(r.Body)
		err := dec.Decode(&product)
		if err != nil {
			panic(err)
		}

		err = h.ProductUsecase.Store(r.Context(), product)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Write([]byte(`{"message": "success"}`))
	} else {
		w.Write([]byte(`{"message": "error"}`))
	}
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "PUT" {
		fmt.Println("update")
		rawID := r.FormValue("id")
		id, err := strconv.Atoi(rawID)
		if err != nil {
			fmt.Println("1")
			fmt.Println(err)
			w.Write([]byte(err.Error()))
			return
		}

		product := models.Product{}
		r.Body = http.MaxBytesReader(w, r.Body, 1048576)
		dec := json.NewDecoder(r.Body)
		err = dec.Decode(&product)
		if err != nil {
			fmt.Println("2")
			fmt.Println(err)
			panic(err)
		}

		product.ID = int64(id)

		err = h.ProductUsecase.Update(r.Context(), product)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "success"}`))
	} else {
		w.Write([]byte(`{"message": "error"}`))
	}
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "DELETE" {
		rawID := r.FormValue("id")
		id, err := strconv.Atoi(rawID)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		err = h.ProductUsecase.Delete(r.Context(), int64(id))
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "success"}`))
	} else {
		w.Write([]byte(`{"message": "error"}`))
	}
}
