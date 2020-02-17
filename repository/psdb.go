package repository

import (
	"database/sql"
	"fmt"
	"fsdemo-priceservice/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var MysqlDb *sql.DB
var err error

const (
	USER_NAME = "xiaodong"
	PASS_WORD = "time4@FUN"
	DBHOST    = "mysqldb"
	DBPORT    = "3306"
	DATABASE  = "productprice"
	CHARSET   = "utf8mb4"
)

func init() {

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true", USER_NAME, PASS_WORD, DBHOST, DBPORT, DATABASE, CHARSET)

	MysqlDb, err = sql.Open("mysql", dbDSN)
	if err != nil {
		log.Println("dbDSN: ", dbDSN)
		panic("数据源配置不正确: " + err.Error())
	}

	//最大连接数
	MysqlDb.SetMaxOpenConns(100)
	//闲置连接数
	MysqlDb.SetMaxIdleConns(20)
	//最大连接周期
	MysqlDb.SetConnMaxLifetime(100 * time.Second)

	if err = MysqlDb.Ping(); err != nil {
		panic("数据库连接失败： " + err.Error())
	}
}

func GetProductPriceByIDAndName(productId int, productName string) (productPrice float64, errCode error) {

	log.Println("repository: GetProductPriceByIDAndName called with productID: ", productId, " and productName: ", productName)

	queryStr := "select product_price from productpricetable where product_id=? and product_name=?"
	result := MysqlDb.QueryRow(queryStr, productId, productName)
	if errCode = result.Scan(&productPrice); errCode != nil {
		log.Println("repository: GetProductPriceByIDAndName: query error for product price with ID and name: ", productId, productName)
		productPrice = 0
	}

	log.Println("repository: GetProductPriceByIDAndName: got product price: ", productPrice)

	return
}

func GetAllProductPrice() (productPriceList []model.ProductPrice, errCode error) {

	log.Println("repository: GetAllProductPrice called.")

	queryStr := "select * from productpricetable"

	results, err := MysqlDb.Query(queryStr)

	defer results.Close()

	if err != nil {
		log.Println("repository: getAllProductPrice: query error for all product price.")
		errCode = err
		return
	}

	for results.Next() {

		pp := model.ProductPrice{}
		var nullPrice sql.NullFloat64
		var nullTime sql.NullTime

		errCode = results.Scan(&pp.ProductId, &pp.ProductName, &nullPrice, &nullTime)
		if errCode != nil {
			log.Println("repository: getAllProductPrice: Result scan failure with error code: ", errCode.Error())
			return
		}
		if nullPrice.Valid {
			pp.Price = nullPrice.Float64
		}
		if nullTime.Valid {
			pp.UpdateTime = nullTime.Time
		}

		productPriceList = append(productPriceList, pp)
	}

	log.Println("repository: getAllProductPrice: query for all product price succeeded with result as below: ")
	log.Println(productPriceList)
	return
}
