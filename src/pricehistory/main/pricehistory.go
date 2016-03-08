package main

import (
	"pricehistory/controller"
)

func main() {
	//	crawler.Crawl()
	//	crawler.GetMenuLinks("http://rozetka.com.ua/fatmenu/action=getMenu;ssl=0/")
//		links := database.GetLinks()
//		fmt.Println(links)
//		fmt.Println(len(links))
//		for _,link := range links {
//			fmt.Println("PROCESSING CATALOG - " + link.LinkText + " = " + link.LinkHref)
//			crawler.ProcessCatalog(link.LinkHref)
//		}
	//	processID, linkHref := database.GetUnprocessedLink()
	//	fmt.Println(strconv.Itoa(processID) + " : " + linkHref)
//		process.InitLinkProcesses()
//		process.Process()
//	crawler.ProcessCatalog("http://rozetka.com.ua/4625469/c4625469/tip-76329=347907/")
//
//	product := database.GetProductWithPrices("5942505")
//	fmt.Println(product.ProductID)
//	fmt.Println(product.ProductName)
//
//	for _, price := range product.ProductPrices {
//
//		fmt.Print(price.Date)
//		fmt.Print(" - ")
//		fmt.Println(price.Price)
//
//	}

	controller.Run()

}
