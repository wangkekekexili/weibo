package user

import (
	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Update(db *sqlx.DB, userID int64, userScreenName string) error {
	_, err := db.Exec("INSERT INTO users (`id`, `screen_name`) VALUES (?, ?)", userID, userScreenName)
	if err != nil {
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == 1062 {
			return nil
		}
		return err
	}
	return nil
}
