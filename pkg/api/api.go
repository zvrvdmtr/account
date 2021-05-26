package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zvrvdmtr/account/pkg/services"
	"strconv"
)

type Some struct {
	Name string
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Println("index route")
}

func GetUserAccountBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
		services.GetUserBalance(userId)
		w.WriteHeader(http.StatusOK)
		return
	}
}

func ChangeUserAccountBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := struct {
			UserId    int
			Amount    float64
			Direction string
		}{}

		byteBody, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(byteBody, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		services.ChangeBalance(dto.UserId, dto.Amount, dto.Direction)
		w.WriteHeader(http.StatusAccepted)
		return
	}
}

func TransactionUserToUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dto := struct {
			withdrawUserId int
			refillUserId   int
			amount         float64
		}{}

		byteBody, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(byteBody, &dto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		services.Transaction(dto.withdrawUserId, dto.refillUserId, float64(dto.amount))
		w.WriteHeader(http.StatusAccepted)
		return
	}
}
