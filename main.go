package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/tidwall/gjson"
	"github.com/wangkekekexili/weibo/micro_blog"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
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
		id := card.Get("mblog.id").String()
		b, err := get(fmt.Sprintf("https://m.weibo.cn/statuses/extend?id=%s&standalone=0", id))
		if err != nil {
			log.Fatal(err)
		}

		m, err := micro_blog.NewFromJSON(id, gjson.Get(string(b), "data").String())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(m)
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
