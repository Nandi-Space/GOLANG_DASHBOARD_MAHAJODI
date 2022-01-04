package store

import "Mahajodi_GOLANG_Dashboard/models"

const (
	//query for payment Details
	queryGetPayment = `SELECT 
	payment.id,
	payment.method,
	payment.transaction_id,
	payment.amount,
	payment.plan_id,
	payment.pay_at,
	tbl_user.name,
	tbl_user.email
	FROM  payment
	LEFT JOIN  tbl_user ON ( payment.user_id = tbl_user.id )
	ORDER BY payment.pay_at DESC;`
)

//GetPayment retrieves data of payment users
func (state *State) GetPayment() ([]models.Payment, error) {

	stmt, err := state.db.Prepare(queryGetPayment)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	results := make([]models.Payment, 0)
	for rows.Next() {
		var payment models.Payment
		if err := rows.Scan(&payment.ID, &payment.Method, &payment.TransactionId, &payment.Amount, &payment.PlanId, &payment.PayAt, &payment.Name, &payment.Email); err != nil {
			return nil, err
		}
		results = append(results, payment)
	}
	return results, nil
}
