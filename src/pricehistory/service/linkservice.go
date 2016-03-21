package service
import (
	"pricehistory/crawler"
	"pricehistory/database"
	"log"
)

func SaveLinks(menu crawler.Menu) {
	for _, v := range menu {
		for _, c := range v.Children {
			if len(crawler.Child(c).Columns) == 0 {
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
