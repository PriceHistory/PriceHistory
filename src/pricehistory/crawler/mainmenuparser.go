package crawler
import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
)

//parses all-categories-goods pages and return all links
func ParseMenuPage(menuPageURL string) []Link {
	document, err := goquery.NewDocument(menuPageURL)
	if (err != nil) {
		panic(menuPageURL)
	}
	return getLinks(document)
}

func getLinks(menuPageDocument *goquery.Document) []Link {
	fmt.Println("Getting links from menu page")
	childLinks := menuPageDocument.Find(".all-cat-b-l-i-link-child")
	var links []Link
	for i := range childLinks.Nodes {
		childLink := childLinks.Eq(i)
		href, _ := childLink.Attr("href")
		title := strings.TrimSpace(childLink.Text());
		link := Link{Href: href, Title: title}
		links = append(links, link)
	}
	return links

}

