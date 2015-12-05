package transporter

import "database/sql"

//Tx is our implementation of the sql.Tx, we did this to make the sql errors
//visible to transporter user when the databse operation fails.
type Tx struct {
	*sql.Tx
	err error
}

//Exec is a wrapper to the Exec function inside sql.Tx.
func (tx *Tx) Exec(query string) (sql.Result, error) {
	result, err := tx.Tx.Exec(query)
	if err != nil {
		tx.err = err
	}

	return result, err
}
