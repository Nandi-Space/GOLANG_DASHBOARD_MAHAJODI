package store

import "Mahajodi_GOLANG_Dashboard/models"

const (
	//query for total users
	queryTotalMembers = `SELECT COUNT(*) FROM tbl_user;`

	//query for total male users
	queryTotalMales = `SELECT COUNT(*) FROM tbl_user LEFT JOIN detail_profile ON (detail_profile.user_id = tbl_user.id) WHERE gender = "Male";`

	//query for total female users
	queryTotalFemales = `SELECT COUNT(*) FROM tbl_user LEFT JOIN detail_profile ON (detail_profile.user_id = tbl_user.id) WHERE gender = "Female";`

	//query for males
	queryGetMales = `SELECT tbl_user.id,
	tbl_user.name,
	detail_profile.community,
	detail_profile.date_of_birth,
	detail_profile.marital_status,
	detail_profile.country,
	detail_profile.district,
	detail_profile.education,
	detail_profile.religion,
	detail_profile.profession
	FROM tbl_user
	LEFT JOIN detail_profile ON
	(detail_profile.user_id = tbl_user.id)
	WHERE gender = "Male"
	ORDER BY created_at DESC;`

	//query for females
	queryGetFemales = `SELECT tbl_user.id,
	tbl_user.name,
	detail_profile.community,
	detail_profile.date_of_birth,
	detail_profile.marital_status,
	detail_profile.country,
	detail_profile.district,
	detail_profile.education,
	detail_profile.religion,
	detail_profile.profession
	FROM tbl_user
	LEFT JOIN detail_profile ON
	(detail_profile.user_id = tbl_user.id)
	WHERE gender = "Female"
	ORDER BY created_at DESC;`

	//query for free users
	queryGetFreeUsers = `SELECT tbl_user.id,
	tbl_user.name,
	detail_profile.community,
	detail_profile.date_of_birth,
	detail_profile.marital_status,
	detail_profile.country,
	detail_profile.district,
	detail_profile.education,
	detail_profile.religion,
	detail_profile.plan_id,
	detail_profile.profession
	FROM tbl_user
	LEFT JOIN detail_profile ON
	(detail_profile.user_id = tbl_user.id)
	WHERE detail_profile.plan_id = 0 `

	//query for basic 3 months users
	queryGetBasic3MonthsUsers = `SELECT tbl_user.id,
	tbl_user.name,
	detail_profile.community,
	detail_profile.date_of_birth,
	detail_profile.marital_status,
	detail_profile.country,
	detail_profile.district,
	detail_profile.education,
	detail_profile.religion,
	detail_profile.plan_id,
	detail_profile.profession
	FROM tbl_user
	LEFT JOIN detail_profile ON
	(detail_profile.user_id = tbl_user.id)
	WHERE detail_profile.plan_id = 1 `

	//query for basic 3 months users
	queryGetBasic6MonthsUsers = `SELECT tbl_user.id,
	tbl_user.name,
	detail_profile.community,
	detail_profile.date_of_birth,
	detail_profile.marital_status,
	detail_profile.country,
	detail_profile.district,
	detail_profile.education,
	detail_profile.religion,
	detail_profile.plan_id,
	detail_profile.profession
	FROM tbl_user
	LEFT JOIN detail_profile ON
	(detail_profile.user_id = tbl_user.id)
	WHERE detail_profile.plan_id = 2 `

	//query for premium users
	queryGetPremium6MonthsUsers = `SELECT tbl_user.id,
	tbl_user.name,
	detail_profile.community,
	detail_profile.date_of_birth,
	detail_profile.marital_status,
	detail_profile.country,
	detail_profile.district,
	detail_profile.education,
	detail_profile.religion,
	detail_profile.plan_id,
	detail_profile.profession
	FROM tbl_user
	LEFT JOIN detail_profile ON
	(detail_profile.user_id = tbl_user.id)
	WHERE detail_profile.plan_id = 3 `
)

//GetFreeUsers returns data of free users
func (state *State) GetFreeUsers() ([]models.User, error) {
	stmt, err := state.db.Prepare(queryGetFreeUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Community, &user.DOB, &user.MaritalStatus, &user.Country, &user.District, &user.Eduction, &user.Religion, &user.PlanId, &user.Profession); err != nil {
			return nil, err
		}
		results = append(results, user)
	}
	return results, nil
}

//GetBasic3MonthsUsers gets data of basic 3 months users
func (state *State) GetBasic3MonthsUsers() ([]models.User, error) {
	stmt, err := state.db.Prepare(queryGetBasic3MonthsUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Community, &user.DOB, &user.MaritalStatus, &user.Country, &user.District, &user.Eduction, &user.Religion, &user.PlanId, &user.Profession); err != nil {
			return nil, err
		}
		results = append(results, user)
	}
	return results, nil
}

//GetBasic6MonthsUsers gets data of basic 6 months users
func (state *State) GetBasic6MonthsUsers() ([]models.User, error) {
	stmt, err := state.db.Prepare(queryGetBasic6MonthsUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Community, &user.DOB, &user.MaritalStatus, &user.Country, &user.District, &user.Eduction, &user.Religion, &user.PlanId, &user.Profession); err != nil {
			return nil, err
		}
		results = append(results, user)
	}
	return results, nil
}

//GetPremium6MonthsUsers gets data of premium 6 months users
func (state *State) GetPremium6MonthsUsers() ([]models.User, error) {
	stmt, err := state.db.Prepare(queryGetPremium6MonthsUsers)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	results := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Community, &user.DOB, &user.MaritalStatus, &user.Country, &user.District, &user.Eduction, &user.Religion, &user.PlanId, &user.Profession); err != nil {
			return nil, err
		}
		results = append(results, user)
	}
	return results, nil
}

//GetTotalMembers returns total number of users
func (state *State) GetTotalMembers() (int, error) {
	var numberCount int
	err := state.db.QueryRow(queryTotalMembers).Scan(&numberCount)
	return numberCount, err
}

//GetTotalMales returns total number of male users
func (state *State) GetTotalMales() (int, error) {
	var numberCount int
	err := state.db.QueryRow(queryTotalMales).Scan(&numberCount)
	return numberCount, err
}

//GetTotalFemales returns total number of female users
func (state *State) GetTotalFemales() (int, error) {
	var numberCount int
	err := state.db.QueryRow(queryTotalFemales).Scan(&numberCount)
	return numberCount, err
}

//GetMales retrieves data of male users
func (state *State) GetMales() ([]models.User, error) {
	stmt, err := state.db.Prepare(queryGetMales)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Community, &user.DOB, &user.MaritalStatus, &user.Country, &user.District, &user.Eduction, &user.Religion, &user.Profession); err != nil {
			return nil, err
		}
		results = append(results, user)
	}

	return results, nil
}

//GetFemales retrieves data of female users
func (state *State) GetFemales() ([]models.User, error) {

	stmt, err := state.db.Prepare(queryGetFemales)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]models.User, 0)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Community, &user.DOB, &user.MaritalStatus, &user.Country, &user.District, &user.Eduction, &user.Religion, &user.Profession); err != nil {
			return nil, err
		}
		results = append(results, user)
	}
	return results, nil
}
