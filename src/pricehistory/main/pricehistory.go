package main

import (
	"pricehistory/process"
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
		process.InitLinkProcesses()
		process.Process()
//	crawler.ProcessCatalog("http://rozetka.com.ua/4625469/c4625469/tip-76329=347907/")
}
