package service
import (
	"pricehistory/crawler"
	"pricehistory/database"
)

func SaveLinks(links []crawler.Link) {
	for _, link := range links {
		database.SaveLink(link.Href, link.Title)
	}
}
