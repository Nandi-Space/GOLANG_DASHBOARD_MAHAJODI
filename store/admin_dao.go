package store

import (
	"Mahajodi_GOLANG_Dashboard/models"
	"log"
)

const (
	//query check admin
	queryCheckAdmin = `SELECT id FROM tbl_admin WHERE email=?;`

	//Query to get admin user using moile number
	queryGetAdmin = `SELECT id,	username, email, otp FROM tbl_admin WHERE email=?;`
	queryGetAdmin2 = `SELECT id,	username, email, otp FROM tbl_admin WHERE id=?;`

	//query save otp
	querySaveOTP = `UPDATE tbl_admin SET otp=? WHERE email=?;`

	//query emove otp
	queryRemveOTP = `UPDATE tbl_admin SET otp="" WHERE email=?;`
)

//IsPresent returns true if the email is pesent else returns false
func (state *State) IsPresent(email string) ( bool, error) {
	
	_,err := state.db.Exec(queryCheckAdmin, email)
	if err != nil {
		log.Println(err)
		return  false, err
	}
	return true, nil
}
//GetAdmin returns admin based on id provided
func (state *State) GetAdmin2(id int64) (models.Admin, error) {
	var admin models.Admin
	err := state.db.QueryRow(queryGetAdmin2, id).Scan(&admin.ID, &admin.UserName,&admin.Email, &admin.OTP)
	return admin, err
}


//GetAdmin returns admin based on email provided
func (state *State) GetAdmin(email string) (models.Admin, error) {
	var admin models.Admin
	err := state.db.QueryRow(queryGetAdmin, email).Scan(&admin.ID, &admin.UserName,&admin.Email, &admin.OTP)
	return admin, err
}

//SaveOTP saves the OTP based on email provided
func (state *State) SaveOTP(email string, otp string) error {
	_, err := state.db.Exec(querySaveOTP, otp, email)
	return err
}

//DeleteOTP removes otp code
func (state *State) DeleteOTP(email string) error {
	_, err := state.db.Exec(queryRemveOTP, email)
	return err
}
