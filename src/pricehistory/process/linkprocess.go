package process

import (
	"log"
	"pricehistory/crawler"
	"pricehistory/database"
)

func Process() {
	log.Println("Start processing")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in f", r)
		}
	}()
	for {
		processID, link := database.GetUnprocessedLink()
		if processID == 0 {
			log.Println("Finished processing. No more links")
			break
		}
		crawler.ProcessCatalog(link)
		database.UpdateLinkProcessStatus(processID, 1)
	}

}
