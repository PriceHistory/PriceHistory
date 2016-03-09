package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"pricehistory/process/status"
	"pricehistory/entity"
	"github.com/ungerik/go-dry"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "user=postgres password=55642325 dbname=gotest sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
}

func Save(id string, title string, price int) {
	insertedProductPK := saveProduct(id, title)
	if insertedProductPK == 0 {
		log.Println("Failed while saving product")
		return
	}
	savePrice(insertedProductPK, price, db)
}

func saveProduct(id string, title string) int {
	var productPK int
	db.QueryRow("SELECT ProductPK FROM product WHERE ProductOuterID = $1", id).Scan(&productPK)
	if productPK != 0 {
		return productPK
	}
	err := db.QueryRow("INSERT INTO product(ProductOuterID, ProductTitle) values($1, $2) RETURNING ProductPK", id, title).Scan(&productPK)
	if err != nil {
		log.Fatal(err)
	}
	return productPK
}

func savePrice(productID int, price int, db *sql.DB) {
	_, err := db.Exec("INSERT INTO price(ProductFK, PriceDate, Price) values($1, now(), $2)", productID, price)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveLink(href string, text string) {
	_, err := db.Exec("INSERT INTO link(LinkHref, LinkText) values($1, $2)", href, text)
	if err != nil {
		log.Println(err)
	}
	log.Println("Saved link. href: " + href + " text: " + text)
}

func GetLinks() []entity.Link {
	var linkPK int
	var linkHref string
	var linkText string
	rows, err := db.Query("SELECT LinkPK, LinkHref, LinkText FROM link")
	if err != nil {
		log.Fatal(err)
	}
	var links []entity.Link
	for rows.Next() {
		rows.Scan(&linkPK, &linkHref, &linkText)
		currentLink := entity.Link{linkPK, linkHref, linkText}
		links = append(links, currentLink)
	}
	return links
}

func ClearLinkProcesses() {
	_, err := db.Exec("DELETE FROM LinkProcess")
	if err != nil {
		log.Println(err)
	}
}

func AddLinkProcess(linkID int, status int) {
	_, err := db.Exec("INSERT INTO LinkProcess(LinkFK, Status) VALUES($1, $2)", linkID, status)
	if err != nil {
		log.Println(err)
	}
}

func GetUnprocessedLink() (int, string) {
	var linkProcessID int
	var linkHref string
	db.QueryRow("select lp.LinkProcessPK, l.LinkHref from LinkProcess lp INNER JOIN link l ON lp.LinkFK = l.LinkPK WHERE Status = $1 LIMIT 1;", status.NotProcessed).Scan(&linkProcessID, &linkHref)
	return linkProcessID, linkHref
}

func UpdateLinkProcessStatus(linkProcessID int, status int) {
	db.Exec("UPDATE LinkProcess SET Status = $1 WHERE LinkProcessPK = $2", status, linkProcessID)
}

func GetProductWithPrices(productOuterID string) entity.Product {
	var productID int
	var productName string
	var price int
	var priceDate time.Time
	rows, err := db.Query("select pr.productpk, pr.producttitle, p.price, p.pricedate from price p join product pr on p.productfk = pr.productpk where productouterid = $1", productOuterID)
	dry.PanicIfErr(err)
	var product entity.Product
	var productPrices []entity.ProductPrice
	for rows.Next() {
		rows.Scan(&productID, &productName, &price, &priceDate)
		product.ProductID = productID
		product.ProductName = productName
		productPrice := entity.ProductPrice{price, priceDate}
		productPrices = append(productPrices, productPrice)
	}
	product.ProductPrices = productPrices
	return product
}
