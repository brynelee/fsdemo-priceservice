package service

import (
	"fsdemo-priceservice/model"
	"fsdemo-priceservice/repository"
	"log"
	"math/big"
)

type ProductPriceService interface {
	GetProductPriceByIDAndName(productId int, productName string) (productPrice big.Float)
}

// mysql implementation
func GetProductPriceByIDAndName(productId int, productName string) (productPrice float64, errCode error) {

	log.Printf("service: GetProudctByIDAndName: productId: %d, product name: %s", productId, productName)
	return repository.GetProductPriceByIDAndName(productId, productName)
}

func GetProductPriceList(ppuList []model.ProductPrice) (ppuListUpdated []model.ProductPrice, errCode error) {

	log.Printf("service: GetProductPriceList: ", ppuList)

	return repository.GetProductPriceList(ppuList)
}

func GetAllProductPrice() (productPriceList []model.ProductPrice, errCode error) {

	log.Printf("service: getAllProductPrice: ", productPriceList)
	return repository.GetAllProductPrice()
}
