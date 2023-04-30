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

	"github.com/SkyAPM/go2sky"
)

const baseURL = "/priceservice"

var psCtrler *PSController

type PSController struct {
	Mux    *http.ServeMux
	err    error
	tracer *go2sky.Tracer
}

func NewBaseController(tracer *go2sky.Tracer) *PSController {

	log.Print("controller: NewBaseController: Calling into the NewBaseController...")

	newController := &PSController{
		http.NewServeMux(),
		nil,
		tracer,
	}

	//handle "/priceservice"
	newController.Mux.HandleFunc(baseURL, helloRequest)

	//handle "/priceservice/prices?productid=xxx&productname=xxx"
	newController.Mux.HandleFunc(baseURL+"/prices", handlePricesRequest)

	//handle "/priceservice/pricelist
	newController.Mux.HandleFunc(baseURL+"/pricelist", handleGetPriceList)

	psCtrler = newController

	return newController
}

// handle GET "/priceservice"
// input: none
// output: single line of text
func helloRequest(writer http.ResponseWriter, request *http.Request) {
	_, _ = fmt.Fprintf(writer, "Hello, this is from Price Service, request url is: ", request.URL.Path[1:])
}

// handle request to the "/priceservice/prices"
func handlePricesRequest(writer http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		_ = r.ParseForm()
		log.Println("controller: handlGetPrice: got request Header with ", r.Header)
		log.Println("controller: handlGetPrice: got request Body with ", r.Body)
		log.Println("controller: handlGetPrice: got request Form with ", r.Form)

		if len(r.Form) > 0 {
			/*subSpan, newCtx, err := psCtrler.tracer.CreateLocalSpan(ctx)
			if err != nil {
				log.Fatalf("error happened in handlePricesRequest handleGetPrice create local span \n")
			}*/
			// single product price query
			handleGetPrice(writer, r)
		} else {
			// query for all product price
			handleGetPriceList(writer, r)
		}

	} else {
		handlePostPrices(writer, r)
	}

}

// handle GET "/priceservice/prices?productid=xxx&productname=xxx"
// input: productid int, productname string
// output: struct { productid int, productname string, productprice big.Float}
func handleGetPrice(writer http.ResponseWriter, r *http.Request) {

	span, _, err := psCtrler.tracer.CreateEntrySpan(r.Context(),
		"handleGetPrice",
		func(key string) (string, error) {
			return r.Header.Get(key), nil
		})
	if err != nil {
		log.Fatalf("error happened in handleGetPrice create local span \n")
	}
	//span.SetOperationName("handleGetPrice request")
	span.Tag("handleGetPrice", "in")

	productId, err := strconv.Atoi(r.FormValue("productid"))
	if err != nil {
		log.Println("productid processing failure")
	}
	productName := r.FormValue("productname")

	prodPrice, err := service.GetProductPriceByIDAndName(productId, productName)
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

	span.End()

	return
}

// handle POST "/priceservice/prices"
// input:
// output: struct { productid int, productname string, productprice big.Float}
func handlePostPrices(writer http.ResponseWriter, r *http.Request) {

	span, _, errSky := psCtrler.tracer.CreateEntrySpan(r.Context(),
		"handlePostPrices",
		func(key string) (string, error) {
			return r.Header.Get(key), nil
		})
	if errSky != nil {
		log.Fatalf("error happened in handlePostPrices create local span \n")
	}
	//span.SetOperationName("handlePostPrices request")
	span.Tag("handlePostPrices", "in")

	var err error

	bodyLength := r.ContentLength
	body := make([]byte, bodyLength)
	_, err = r.Body.Read(body)
	if err != nil {
		log.Println("controller: handlePostPrices: error got from r.Body.Read(body) and error is: ", err.Error())
	}

	var ppuList []model.ProductPrice

	err = json2.Unmarshal(body, &ppuList)

	if err != nil {
		log.Println("controller: handlePostPrices: error got from json2. Unmarshal: " + err.Error())
	}

	log.Println("controller: handlePostPrices: got POST request with payload: ", ppuList)

	updatedPPUList, err := service.GetProductPriceList(ppuList)

	if err != nil {
		log.Println("controller: handlePostPrices: service.GetProductPriceList called failed with error: ", err.Error())
		writer.WriteHeader(500)
		_, _ = fmt.Fprintf(writer, "查询价格失败。错误返回信息： %s", err.Error())
		return
	}

	//format data to json format
	writer.Header().Set("Content-Type", "application/json")
	json, _ := json2.Marshal(updatedPPUList)
	_, _ = writer.Write(json)

	span.End()

	return

}

// handle GET "/priceservice/pricelist"
// input:
// output:
func handleGetPriceList(writer http.ResponseWriter, r *http.Request) {

	span, _, errSky := psCtrler.tracer.CreateEntrySpan(r.Context(),
		"handleGetPriceList",
		func(key string) (string, error) {
			return r.Header.Get(key), nil
		})
	if errSky != nil {
		log.Fatalf("error happened in handleGetPriceList create local span \n")
	}
	//span.SetOperationName("handleGetPriceList request")
	span.Tag("handleGetPriceList", "in")

	log.Println("controller: handlGetPrice: got request Header with ", r.Header)
	log.Println("controller: handlGetPrice: got request Body with ", r.Body)
	log.Println("controller: handlGetPrice: got request Form with ", r.Form)

	//get query results
	prodPriceList, err := service.GetAllProductPrice()
	if err != nil {
		log.Println("controller: handleGetPriceList: got error with code: ", err.Error())
		return
	}

	log.Println("controller: handleGetPriceList: got productPriceList ", prodPriceList)

	//format data to json format
	writer.Header().Set("Content-Type", "application/json")
	json, _ := json2.Marshal(prodPriceList)
	_, _ = writer.Write(json)

	span.End()

	return
}
