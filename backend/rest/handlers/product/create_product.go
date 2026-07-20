package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imageUrl"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	// r.Body => description, imageUrl, price, title => Product er structure e convert kora => append

	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "please give me valid JSON", 201)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title:       req.Title,
		Description: req.Description,
		Price:       req.Price,
		ImgURL:      req.ImgURL,
	})

	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusCreated, createdProduct)
}
