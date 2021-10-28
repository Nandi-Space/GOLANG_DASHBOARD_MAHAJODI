package auth

import (
	"Mahajodi_GOLANG_Dashboard/store"
	"Mahajodi_GOLANG_Dashboard/utils"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

//respone after login
type loginResponse struct{
	AccessToken string `json:"access_token"`
	Admin adminResponse `json:"admin"`
}
//adminResponse
type adminResponse struct{
	ID       int64  `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
}

//VerifyEmail verifies if any admin is present with the provided email
func VerifyEmail(w http.ResponseWriter, r *http.Request) {
	//supposed to get request body in this format
	var req struct {
		Email string `json:"email"`
	}
	// Populating req struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, "invalid json data", 400)
		return
	}
	//  checking if admin is present or not
	isAdmin, err := store.DBState.IsPresent(req.Email)
	//handling error
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(w, "internal server error", 500)
		return
	}
	//if admin is not present
	if !isAdmin {
		utils.ErrorResponse(w, "no admin with provided email", 404)
		return
	}

	//Generating otp
	otp := utils.GenerateOTP(6)
	//saving otp in DB
	saveErr := store.DBState.SaveOTP(req.Email, otp)
	if saveErr != nil {
		logrus.Debug("error while saving otp: ", saveErr)
		utils.JsonResponse(w, "error while saving otp", 500)
		return
	}
	// Fetching admin data
	admin, err := store.DBState.GetAdmin(req.Email)
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(w, "admin not found", 404)
		return
	}
	// sending otp over email
	otpErr := store.DBState.SendOTP(admin.Email, admin.UserName, otp)
	if otpErr != nil {
		logrus.Debug("error while sending otp: ", otpErr)
		utils.JsonResponse(w, "error while sending otp", 500)
		return
	}
	//sending response
	utils.JsonResponse(w, "Email verified successfully", 200)
}

//VerifyOtp is the login handler
func VerifyOtp(w http.ResponseWriter, r *http.Request) {

	//supposed to get request body in this format
	var req struct {
		Email string `json:"Email"`
		OTP   string `json:"otp"`
	}

	// Populating req struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		utils.ErrorResponse(w, "invalid json data", 400)
		return
	}
	
	// Fetching admin data
	admin, err := store.DBState.GetAdmin(req.Email)
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(w, "admin not found", 404)
		return
	}
	//comparing otp from request and DB
	result := strings.Compare(req.OTP, admin.OTP)
	if result != 0 {
		logrus.Debug("otp didnt match")
		utils.ErrorResponse(w, "incorrect otp", 404)
		return
	}
	//token
	token, err := utils.SignData(admin.ID)
	if err != nil {
		logrus.Debug(err)
		utils.ErrorResponse(w, "could not generate token", 500)
		return
	}

	//Once verified remove the otp from DB
	removeErr := store.DBState.DeleteOTP(admin.Email)
	if removeErr != nil {
		logrus.Debug("error while deleting otp: ", removeErr)
		utils.ErrorResponse(w, "internal server error", 500)
		return
	}
	//preparing response
	adminRsp := adminResponse{
		ID: admin.ID,
		UserName: admin.UserName,
		Email: admin.Email,
	}
	rsp:= loginResponse{
		AccessToken: token,
		Admin: adminRsp,
	}
	//sending response
	utils.JsonResponse(w, rsp, 200)
}
