package main
import (
	"pricehistory/crawler"
	"pricehistory/service"
)

func main() {
	menu := crawler.GetCatalogs("http://rozetka.com.ua/fatmenu/action=getMenu;ssl=0/")
	service.SaveLinks(menu)
}
