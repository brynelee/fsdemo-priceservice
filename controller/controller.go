package controller

import (
	json2 "encoding/json"
	"fmt"
	"fsdemo-priceservice/model"
	"fsdemo-priceservice/service"
	"log"
	"net/http"
	"strconv"
	"time"
)

const baseURL = "/priceservice"

type PSController struct {
	Mux *http.ServeMux
	err error
}

func NewBaseController() *PSController {

	log.Print("controller: NewBaseController: Calling into the NewBaseController...")

	newController := &PSController{
		http.NewServeMux(),
		nil,
	}

	//handle "/priceservice"
	newController.Mux.HandleFunc(baseURL, helloRequest)

	//handle "/priceservice/prices?productid=xxx&productname=xxx"
	newController.Mux.HandleFunc(baseURL+"/prices", handleGetPrice)

	//handle "/priceservice/pricelist
	newController.Mux.HandleFunc(baseURL+"/pricelist", handleGetPriceList)

	return newController
}

// handle GET "/priceservice"
// input: none
// output: single line of text
func helloRequest(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "Hello, this is from Price Service, request url is: ", request.URL.Path[1:])
}

// handle GET "/priceservice/prices?productid=xxx&productname=xxx"
// input: productid int, productname string
// output: struct { productid int, productname string, productprice big.Float}
func handleGetPrice(writer http.ResponseWriter, r *http.Request) {

	_ = r.ParseForm()
	log.Println("controller: handlGetPrice: got request Header with ", r.Header)
	log.Println("controller: handlGetPrice: got request Body with ", r.Body)
	log.Println("controller: handlGetPrice: got request Form with ", r.Form)

	productId, err := strconv.Atoi(r.FormValue("productid"))
	if err != nil {
		log.Println("productid processing failure")
	}
	productName := r.FormValue("productname")

	prodPrice, err := service.GetProudctByIDAndName(productId, productName)
	if err != nil {
		log.Println("controller: handleGetPrice: got error with code: ", err.Error())
		writer.WriteHeader(501)
		_, _ = fmt.Fprintf(writer, "没有找到对应产品的价格。错误返回信息： %s", err.Error())
		return
	}

	log.Println("controller: handleGetPrice: got productPrice: ", prodPrice)

	//respData := model.ProductPrice{productId,productName, productPrice}
	//fmt.Fprintf(writer, )
	writer.Header().Set("Content-Type", "application/json")
	pp := model.ProductPrice{productId, productName, prodPrice, time.Now()}
	json, _ := json2.Marshal(pp)
	_, _ = writer.Write(json)
	return
}

// handle GET "/priceservice/pricelist"
// input:
// output:
func handleGetPriceList(writer http.ResponseWriter, r *http.Request) {

	log.Println("controller: handlGetPrice: got request Header with ", r.Header)
	log.Println("controller: handlGetPrice: got request Body with ", r.Body)
	log.Println("controller: handlGetPrice: got request Form with ", r.Form)

	prodPriceList, err := service.GetAllProductPrice()
	if err != nil {
		log.Println("controller: handleGetPriceList: got error with code: ", err.Error())
		return
	}

	log.Println("controller: handleGetPriceList: got productPriceList ", prodPriceList)

	//respData := model.ProductPrice{productId,productName, productPrice}
	//fmt.Fprintf(writer, )
	writer.Header().Set("Content-Type", "application/json")
	json, _ := json2.Marshal(prodPriceList)
	_, _ = writer.Write(json)
	return
}
