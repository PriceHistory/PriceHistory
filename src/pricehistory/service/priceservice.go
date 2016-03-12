package service
import (
	"pricehistory/entity"
	"pricehistory/database"
)

func GetProduct(productID string) entity.Product {
	return database.GetProductWithPrices(productID)
}
