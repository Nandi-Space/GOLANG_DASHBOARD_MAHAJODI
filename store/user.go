package store

import "Mahajodi_GOLANG_Dashboard/models"

func (state *State) GetAdmin(adminID int64) (models.Admin, error) {
	var user models.Admin
	err := state.db.QueryRow(`SELECT id,username,email,
								FROM tbl_admin
					where id=?`, adminID).
		Scan(&user.Id, &user.UserName, &user.Email)
	return user, err
}
