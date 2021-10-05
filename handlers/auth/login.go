package auth

import (
	"encoding/json"
	"mahajodi/dashboard/store"
	"mahajodi/dashboard/utils"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

//VerifyMobile verifies if any admin is present with the provided mobile
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
	if err != nil {
		logrus.Debug(err)
		if isAdmin != true {
			utils.ErrorResponse(w, "no admin with provided email", 404)
			return
		}
		utils.ErrorResponse(w, "internal server error", 500)
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
	// sending otp over sms
	otpErr := store.DBState.SendOTPAWS(req.Email, otp)
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
		OTP    string `json:"otp"`
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
	removeErr := store.DBState.DeleteOTP(req.Email)
	if removeErr != nil {
		logrus.Debug("error while deleting otp: ", removeErr)
		utils.ErrorResponse(w, "internal server error", 500)
		return
	}
	//sending response
	utils.JsonResponse(w, token, 200)
}
