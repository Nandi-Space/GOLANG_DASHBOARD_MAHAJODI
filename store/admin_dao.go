package store

import (
	"Mahajodi_GOLANG_Dashboard/models"
)

const (
	//query check admin
	queryCheckAdmin = `SELECT id FROM tbl_admin WHERE email=?`

	//Query to get admin user using moile number
	queryGetAdmin = `SELECT id,	username, email, phone,	otp, FROM tbl_admin WHERE email=?`

	//query save otp
	querySaveOTP = `UPDATE tbl_admin SET otp=? WHERE id=?`

	//query emove otp
	queryRemveOTP = `UPDATE tbl_admin SET otp="" WHERE id=?`
)

//IsPresent returns true if the phone is pesent else returns false
func (state *State) IsPresent(email string) (int64, bool, error) {
	var admin models.Admin
	err := state.db.QueryRow(queryGetAdmin, email).Scan(&admin.ID)
	if err != nil {
		return 0, false, err
	}
	return admin.ID, true, nil
}

//GetAdmin returns admin based on phone number provided
func (state *State) GetAdmin(id int64) (models.Admin, error) {
	var admin models.Admin
	err := state.db.QueryRow(queryGetAdmin, id).Scan(&admin.ID, &admin.Phone, &admin.OTP)
	return admin, err
}

//SaveOTP saves the OTP based on phone number provided
func (state *State) SaveOTP(id int64, otp string) error {
	_, err := state.db.Exec(querySaveOTP, otp, id)
	return err
}

//DeleteOTP removes otp code
func (state *State) DeleteOTP(id int64) error {
	_, err := state.db.Exec(queryRemveOTP, id)
	return err
}
