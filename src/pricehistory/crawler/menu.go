package crawler

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"pricehistory/database"
)

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

func GetMenuLinks(url string) {
	getCatalogs(url)
}

func getCatalogs(initialURL string) {
	document, err := goquery.NewDocument(initialURL)
	if err != nil {
		log.Fatal(err)
	}
	menu := new(Menu)
	json.Unmarshal([]byte(document.Text()), &menu)
	for _, v := range *menu {
		for _, c := range v.Children {
			if len(Child(c).Columns) == 0 {
				log.Println(c)
				database.SaveLink(c.Href, c.Title)
			} else {
				for _, col := range c.Columns {
					for _, ch := range col.Children {
						database.SaveLink(ch.Href, ch.Title)
					}
				}
			}
		}
	}
}
