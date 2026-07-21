package user

import (
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid req body")
		return
	}

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid Request Data")
		return
	}
	// fmt.Println("1")
	usr, err := h.userRepo.Find(req.Email, req.Password)

	if err != nil {
		fmt.Println(err)
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if usr == nil {
		util.SendError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, http.StatusCreated, accessToken)
}
