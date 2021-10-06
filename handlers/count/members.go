package count

import (
	"Mahajodi_GOLANG_Dashboard/store"
	"Mahajodi_GOLANG_Dashboard/utils"
	"net/http"

	"github.com/sirupsen/logrus"
)

//GetTotalMembers handles number of total users
func GetTotalMembers(w http.ResponseWriter, r *http.Request) {
	totalMembers, err := store.DBState.GetTotalMembers()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}
	utils.JsonResponse(w, totalMembers, 200)
}

//GetTotalMales handles number of total male users
func GetTotalMales(w http.ResponseWriter, r *http.Request) {
	totalMales, err := store.DBState.GetTotalMales()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}
	utils.JsonResponse(w, totalMales, 200)
}

//GetTotalFemales handles number of total female users
func GetTotalFemales(w http.ResponseWriter, r *http.Request) {
	totalFemales, err := store.DBState.GetTotalFemales()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}
	utils.JsonResponse(w, totalFemales, 200)
}
