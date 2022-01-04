package payment

import (
	"Mahajodi_GOLANG_Dashboard/store"
	"Mahajodi_GOLANG_Dashboard/utils"
	"net/http"

	"github.com/sirupsen/logrus"
)

//Get Payment Details  handles the payment  data
func GetPayment(w http.ResponseWriter, r *http.Request) {

	payment, err := store.DBState.GetPayment()
	if err != nil {
		logrus.Error(err)
		utils.ErrorResponse(w, "internal server error", 500)
	}

	utils.JsonResponse(w, payment, 200)
}
