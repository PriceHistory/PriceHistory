package crawler

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
``)

type Menu map[string]MenuItem

type MenuItem struct {
	Children []Child `json:"children"`
	Banner   Banner  `json:"banner"`
}

type Child struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Href    string   `json:"href"`
	Tag     string   `json:"tag"`
	Columns []Column `json:"columns"`
	Banners []Banner `json:"banners"`
}

type Column struct {
	Children []Child `json:"children"`
}

type Banner struct {
	Image string `json:"image"`
	Href  string `json:"href"`
}

func GetCatalogs(initialURL string) Menu {
	document, err := goquery.NewDocument(initialURL)
	if err != nil {
		log.Fatal(err)
	}
	var menu Menu
	json.Unmarshal([]byte(document.Text()), &menu)
	return menu
}
