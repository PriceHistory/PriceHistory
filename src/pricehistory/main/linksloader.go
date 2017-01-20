package main
import (
	"pricehistory/crawler"
//	"fmt"
	"pricehistory/service"
)

func main() {
//	menu := crawler.GetCatalogs("http://rozetka.com.ua/ua/fatmenu/action=getMenu;ssl=0/")
//	fmt.Println("-----------------------------------------")
//	fmt.Println(menu)
//	fmt.Println(len(menu))
//	service.SaveLinks(menu)

	links := crawler.ParseMenuPage("http://rozetka.com.ua/ua/all-categories-goods/")
	service.SaveLinks(links)


}
