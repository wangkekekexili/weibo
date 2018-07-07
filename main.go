package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/tidwall/gjson"
	"github.com/wangkekekexili/weibo/model/micro_blog"
	"github.com/wangkekekexili/weibo/model/statistics"
	"github.com/wangkekekexili/weibo/model/user"
)

var DB *sqlx.DB

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	DB = sqlx.MustConnect("mysql", "root:@tcp(localhost:3306)/weibo?parseTime=true")
}

func main() {
	b, err := get("https://m.weibo.cn/api/container/getIndex?containerid=102803&openApp=0")
	if err != nil {
		log.Fatal(err)
	}
	ticker := time.NewTicker(100 * time.Millisecond)
	for _, card := range gjson.Get(string(b), "data.cards").Array() {
		<-ticker.C
		if card.Get("card_type").Int() != 9 {
			continue
		}

		microBlogID := card.Get("mblog.id").Int()
		userID := card.Get("mblog.user.id").Int()
		userScreenName := card.Get("mblog.user.screen_name").String()

		b, err := get(fmt.Sprintf("https://m.weibo.cn/statuses/extend?id=%d&standalone=0", microBlogID))
		if err != nil {
			log.Fatal(err)
		}
		data := gjson.Get(string(b), "data")
		text := stripHTML(data.Get("longTextContent").String())
		numThumbUp := int(data.Get("attitudes_count").Int())
		numComment := int(data.Get("comments_count").Int())
		numRepost := int(data.Get("reposts_count").Int())

		err = user.Update(DB, userID, userScreenName)
		if err != nil {
			log.Println(err)
		}

		err = micro_blog.Update(DB, microBlogID, text, userID)
		if err != nil {
			log.Println(err)
		}

		err = statistics.Update(DB, microBlogID, numThumbUp, numComment, numRepost)
		if err != nil {
			log.Println(err)
		}
	}
}

func get(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error calling Do: %v", err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading data: %v", err)
	}
	resp.Body.Close()
	return b, nil
}

func stripHTML(input string) string {
	// Remove <br /> <br> and \n.
	input = strings.NewReplacer("<br />", "", "<br>", "", "\n", "").Replace(input)

	// Remove <a> tags.
	output := strings.Builder{}
	pos := 0
	inTagA := false
	inTagSpan := false
	runes := []rune(input)
	for pos < len(runes) {
		if inTagA {
			if runes[pos] == '<' && pos+3 < len(runes) &&
				runes[pos+1] == '/' && runes[pos+2] == 'a' && runes[pos+3] == '>' {
				pos += 4
				inTagA = false
			} else {
				pos++
			}
		} else if inTagSpan {
			if runes[pos] == '<' && pos+6 < len(runes) &&
				runes[pos+1] == '/' && runes[pos+2] == 's' && runes[pos+3] == 'p' &&
				runes[pos+4] == 'a' && runes[pos+5] == 'n' && runes[pos+6] == '>' {
				pos += 7
				inTagA = false
			} else {
				pos++
			}
		} else {
			if runes[pos] == '<' && pos+2 < len(runes) &&
				runes[pos+1] == 'a' && runes[pos+2] == ' ' {
				pos += 3
				inTagA = true
			} else if runes[pos] == '<' && pos+5 < len(runes) &&
				runes[pos+1] == 's' && runes[pos+2] == 'p' &&
				runes[pos+3] == 'a' && runes[pos+4] == 'n' &&
				runes[pos+5] == ' ' {
				pos += 6
				inTagSpan = true
			} else {
				output.WriteRune(runes[pos])
				pos++
			}
		}
	}

	return output.String()
}
