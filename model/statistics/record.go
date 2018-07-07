package statistics

import "time"

type Record struct {
	ID           int64     `db:"id"`
	MicroBlogID  int64     `db:"micro_blog_id"`
	ObservedTime time.Time `db:"observed_time"`
	NumThumbUp   int       `db:"num_thumb_up"`
	NumComment   int       `db:"num_comment"`
	NumRepost    int       `db:"num_repost"`
}
