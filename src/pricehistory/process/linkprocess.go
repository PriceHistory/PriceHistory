package process

import (
	"log"
	"pricehistory/crawler"
	"pricehistory/database"
	"pricehistory/process/status"
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
		database.UpdateLinkProcessStatus(processID, status.InProgress)
		if processID == 0 {
			log.Println("Finished processing. No more links")
			break
		}
		crawler.ProcessCatalog(link)
		database.UpdateLinkProcessStatus(processID, status.Processed)
	}

}
