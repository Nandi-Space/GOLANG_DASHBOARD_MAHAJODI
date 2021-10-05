package store

import (
	"mahajodi/dashboard/models"
)

const (
	//query check admin
	queryCheckAdmin = `SELECT id FROM tbl_admin WHERE mobile=?`

	//Query to get admin user using moile number
	queryGetAdmin = `SELECT id,	mobile,	isVerified,	otp, FROM tbl_admin WHERE mobile=?`

	//query save otp
	querySaveOTP = `UPDATE tbl_admin SET otp=? WHERE mobile=?`

	//query emove otp
	queryRemveOTP = `UPDATE tbl_admin SET otp="" WHERE mobile=?`
)

//IsPresent returns true if the mobile is pesent else returns false
func (state *State) IsPresent(mobile string) (bool, error) {
	var admin models.Admin
	err := state.db.QueryRow(queryGetAdmin, mobile).Scan(&admin.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

//GetAdmin returns admin based on mobile number provided
func (state *State) GetAdmin(mobile string) (models.Admin, error) {
	var admin models.Admin
	err := state.db.QueryRow(queryGetAdmin, mobile).Scan(&admin.ID, &admin.Mobile, &admin.IsVerified, &admin.OTP)
	return admin, err
}

//SaveOTP saves the OTP based on mobile number provided
func (state *State) SaveOTP(mobile, otp string) error {
	_, err := state.db.Exec(querySaveOTP, otp, mobile)
	return err
}

//DeleteOTP removes otp code
func (state *State) DeleteOTP(mobile string) error {
	_, err := state.db.Exec(queryRemveOTP, mobile)
	return err
}
