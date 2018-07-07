package micro_blog

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

func Update(db *sqlx.DB, microBlogID int64, text string, userID int64) error {
	var dummy int
	err := db.QueryRow("SELECT 1 FROM micro_blogs WHERE id = ?", microBlogID).Scan(&dummy)
	if err == nil {
		return nil
	}
	if err != sql.ErrNoRows {
		return err
	}
	_, err = db.Exec("INSERT INTO micro_blogs (`id`, `text`, `user_id`, `created_at`) VALUES (?, ?, ?, ?)",
		microBlogID, text, userID, time.Now())
	if err != nil {
		return err
	}
	return nil
}
