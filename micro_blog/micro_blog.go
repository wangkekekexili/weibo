package micro_blog

import (
	"encoding/json"
	"fmt"
)

type MicroBlog struct {
	ID string

	Text       string `json:"longTextContent"`
	NumThumbUp int    `json:"attitudes_count"`
	NumComment int    `json:"comments_count"`
	NumRepost  int    `json:"reposts_count"`
}

func NewFromJSON(id, j string) (*MicroBlog, error) {
	var m MicroBlog
	err := json.Unmarshal([]byte(j), &m)
	if err != nil {
		return nil, fmt.Errorf("error decoding into micro blog: %v", err)
	}
	m.ID = id
	return &m, nil
}

func (m *MicroBlog) String() string {
	return fmt.Sprintf("%s: %s", m.ID, m.Text)
}
