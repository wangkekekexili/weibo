package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tidwall/gjson"
)

type microBlog struct {
	ID string `json:"bid"`

	NumThumbUp int `json:"attitudes_count"`
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func main() {
	req, err := http.NewRequest(http.MethodGet, "https://m.weibo.cn/api/container/getIndex?containerid=102803&openApp=0", nil)
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	for _, card := range gjson.Get(string(b), "data.cards").Array() {
		if card.Get("card_type").Int() != 9 {
			continue
		}
		fmt.Println(removeTags(card.Get("mblog.text").String()))
	}
}

func removeTags(s string) string {
	var result []rune

	inTag := false
	for _, r := range []rune(s) {
		if inTag {
			if r == '>' {
				inTag = false
			}
			continue
		}
		if r == '<' {
			inTag = true
			continue
		}
		result = append(result, r)
	}

	return string(result)
}
