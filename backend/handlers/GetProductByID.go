package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Invalid product ID", 400)
		return
	}

	for _, product := range database.Products {
		if product.Id == pId {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "Product not found", 404)
}
