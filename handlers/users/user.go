package users

import (
	"Mahajodi_GOLANG_Dashboard/store"
	"Mahajodi_GOLANG_Dashboard/utils"
	"net/http"

	"github.com/sirupsen/logrus"
)

//GetMales handles the male users data
func GetMales(w http.ResponseWriter, r *http.Request) {

	maleUsers, err := store.DBState.GetMales()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}

	utils.JsonResponse(w, maleUsers, 200)
}

//GetFemales handles the female users data
func GetFemales(w http.ResponseWriter, r *http.Request) {

	femaleUsers, err := store.DBState.GetFemales()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}

	utils.JsonResponse(w, femaleUsers, 200)
}

//
func GetFreeUsers(w http.ResponseWriter, r *http.Request) {
	freeUsers, err := store.DBState.GetFreeUsers()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}

	utils.JsonResponse(w, freeUsers, 200)
}

func GetBasic3MonthsUsers(w http.ResponseWriter, r *http.Request) {
	basic, err := store.DBState.GetBasic3MonthsUsers()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}
	utils.JsonResponse(w, basic, 200)
}

func GetBasic6MonthsUsers(w http.ResponseWriter, r *http.Request) {
	basic, err := store.DBState.GetBasic6MonthsUsers()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}
	utils.JsonResponse(w, basic, 200)
}

func GetPremium6MonthsUsers(w http.ResponseWriter, r *http.Request) {
	premium, err := store.DBState.GetPremium6MonthsUsers()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}
	utils.JsonResponse(w, premium, 200)
}
