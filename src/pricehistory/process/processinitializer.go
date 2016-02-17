package process
import (
	"github.com/vgotsuliak/pricehistory/pricehistory/database"
	"log"
	"strconv"
)

const unprocessedStatus int = 0

func InitLinkProcesses() {
	clearOldProcesses()
	initNewProcesses()
}

func clearOldProcesses() {
	log.Println("Clearing old link processes")
	database.ClearLinkProcesses();
	log.Println("Finished clearing")
}

func initNewProcesses() {
	links := database.GetLinks()
	for _,link := range links {
		database.AddLinkProcess(link.LinkID, unprocessedStatus)
	}
	log.Println("Finished creating processes. Created " + strconv.Itoa(len(links)) + " new processes.")
}