package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=55642325 dbname=gotest sslmode=disable")
	if (err != nil) {
		log.Fatal(err)
	}
}

func Save(id string, title string, price int) {
	insertedProductPK := saveProduct(id, title)
	if (insertedProductPK == 0) {
		fmt.Println("Failed while saving product")
		return
	}
	savePrice(insertedProductPK, price, db)
	fmt.Println("SUCCESS")
}

func saveProduct(id string, title string) int {
	var productPK int
	db.QueryRow("SELECT ProductPK FROM product WHERE ProductOuterID = $1", id).Scan(&productPK)
	if (productPK != 0) {
		return productPK
	}
	err := db.QueryRow("INSERT INTO product(ProductOuterID, ProductTitle) values($1, $2) RETURNING ProductPK", id, title).Scan(&productPK)
	if (err != nil) {
		log.Fatal(err)
	}
	return productPK;
}

func savePrice(productID int, price int, db *sql.DB) {
	_, err := db.Exec("INSERT INTO price(ProductFK, PriceDate, Price) values($1, now(), $2)", productID, price)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveLink(href string, text string) {
	_, err := db.Exec("INSERT INTO link(LinkHref, LinkText) values($1, $2)", href, text)
	if (err != nil) {
		log.Println(err)
	}
	fmt.Println("Saved link. href: " + href + " text: " + text)
}

type Link struct {
	LinkID int
	LinkHref string
	LinkText string
}

func GetLinks() []Link {
	var linkPK int
	var linkHref string
	var linkText string
	rows, err := db.Query("SELECT LinkPK, LinkHref, LinkText FROM link")
	if (err != nil) {
		log.Fatal(err)
	}
	var links []Link
	for rows.Next() {
		rows.Scan(&linkPK, &linkHref, &linkText)
		currentLink := Link{linkPK, linkHref, linkText}
		links = append(links, currentLink)
	}
	return links
}

func ClearLinkProcesses() {
	_, err := db.Exec("DELETE FROM LinkProcess")
	if (err != nil) {
		log.Println(err)
	}
}

func AddLinkProcess(linkID int, status int) {
	_, err := db.Exec("INSERT INTO LinkProcess(LinkFK, Status) VALUES($1, $2)", linkID, status)
	if (err != nil) {
		log.Println(err)
	}
}

func GetUnprocessedLink() (int, string) {
	var linkProcessID int
	var linkHref string
	db.QueryRow("select lp.LinkProcessPK, l.LinkHref from LinkProcess lp INNER JOIN link l ON lp.LinkFK = l.LinkPK WHERE Status = 0 LIMIT 1;").Scan(&linkProcessID, &linkHref)
	return linkProcessID, linkHref;
}

func UpdateLinkProcessStatus(linkProcessID int, status int) {
	db.Exec("UPDATE LinkProcess SET Status = $1 WHERE LinkProcessPK = $2", status, linkProcessID)
}