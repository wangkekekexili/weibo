package micro_blog

import (
	"time"
)

type Record struct {
	ID        int64     `db:"int64"`
	Text      string    `db:"text"`
	UserID    int64     `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}
