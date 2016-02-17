package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strconv"
	"pricehistory/database"
	"pricehistory/util"
	"strings"
	"log"
)

func processCatalogPage(catalogPage *goquery.Document) map[string]string {
	fmt.Println("Processing page: " + catalogPage.Url.String())
	prices := make(map[string]string)
	selection := catalogPage.Find(".g-i-tile")
	nodes := selection.Nodes
	for item := range nodes {
		productDocument := goquery.NewDocumentFromNode(nodes[item])
		priceSelection := productDocument.Find(".g-price-uah")
		if (len(priceSelection.Nodes) == 0) {
			continue
		}
		titleSelection := productDocument.Find(".g-i-tile-i-title")
		idSelection := productDocument.Find(".g-id")
		for j := range idSelection.Nodes {
			id := idSelection.Eq(j).Text()
			prices[id] = priceSelection.Text()
			price := priceSelection.Text()
			fmt.Println(price)
			convertedPrice := util.ConvertPrice(price)
			title := strings.Trim(titleSelection.Text(), "\n	")
			fmt.Println(id + ": " + priceSelection.Text() + " (" + title + ")")
			database.Save(id, title, convertedPrice)
		}
	}
	return prices
}

func ProcessCatalog(catalogFirstPageURL string) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered nextCatalogPageURL", r)
			panic(catalogFirstPageURL)
		}
	}()
	var document *goquery.Document
	for nextPageURL := catalogFirstPageURL; nextPageURL != ""; nextPageURL = nextCatalogPageURL(document) {
		var err error
		document, err = goquery.NewDocument(nextPageURL)
		if (err != nil) {
			fmt.Errorf("Error: " + err.Error())
			panic(catalogFirstPageURL)
		}
		prices := processCatalogPage(document)
		fmt.Println(prices)
	}
}

func nextCatalogPageURL(currentCatalogPage *goquery.Document) string {
	selection := currentCatalogPage.Find(".paginator-catalog-l-i.active")
	if (len(selection.Nodes) == 0) {
		fmt.Println("Paginator not found")
		panic("Error while getting nextCatalogPageURL")
		return ""
	}
	currentPageID, exists := selection.Attr("id")
	if (!exists) {
		fmt.Println("Current page id not found")
		panic("Error while getting nextCatalogPageURL")
		return ""
	}
	nextPageID := getNextPageID(currentPageID)
	pageSelection := currentCatalogPage.Find("#" + nextPageID)
	nextPageLink := pageSelection.Find("a")
	link, _ := nextPageLink.Attr("href")
	return link;
}

func getNextPageID(pageID string) string {
	id := pageID[4:]
	nextID,_ := strconv.Atoi(id)
	nextID++
	return "page" + strconv.Itoa(nextID)
}
