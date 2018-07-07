package user

type Record struct {
	ID         int64  `db:"id"`
	ScreenName string `db:"screen_name"`
}

func New(id int64, screenName string) *Record {
	return &Record{
		ID:         id,
		ScreenName: screenName,
	}
}
