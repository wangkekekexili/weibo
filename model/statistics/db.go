package statistics

import (
	"time"

	"github.com/jmoiron/sqlx"
)

func Update(db *sqlx.DB, microBlogID int64, numThumbUp, numComment, numRepost int) error {
	_, err := db.Exec("INSERT INTO statistics (`micro_blog_id`, `observed_time`, `num_thumb_up`, `num_comment`, num_repost) VALUES (?, ?, ?, ?, ?)",
		microBlogID, time.Now(), numThumbUp, numComment, numRepost)
	return err
}
