package crawler

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"fmt"
)

type Menu map[string]MenuItem

type MenuItem struct {
	Children Child    `json:"children"`
	Banner   Banner   `json:"banner"`
}

type Child struct {
	Popular Popular  `json:"popular"`
	Columns Column `json:"columns"`
}

type Popular struct {
	Links []Link `json:"links"`
}

type Link struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Href  string `json:"href"`
}

type Column map[string][]ColumnItem

type ColumnItem struct {
	Title Link `json:"title"`
	Links []Link `json:"links"`
}

type Banner struct {
	Image string `json:"image"`
	Href  string `json:"href"`
}

func GetCatalogs(initialURL string) Menu {
	document, err := goquery.NewDocument(initialURL)
	fmt.Println(document.Text())
	if err != nil {
		log.Fatal(err)
	}
	var menu Menu
	json.Unmarshal([]byte(document.Text()), &menu)
	fmt.Println(menu)
	return menu
}
