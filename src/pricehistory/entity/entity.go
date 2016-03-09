package entity
import "time"

type Product struct {
	ProductID int
	ProductName string
	ProductPrices []ProductPrice
}

type ProductPrice struct {
	Price int
	Date time.Time
}

type Link struct {
	LinkID   int
	LinkHref string
	LinkText string
}

