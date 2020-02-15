package controller

import (
	"fmt"
	"fsdemo-priceservice/service"
	"log"
	"net/http"
)

const baseURL = "/priceservice"

func NewBaseController() *PSController {
	//mux := http.NewServeMux()
	//mux.Handle("/")
	fmt.Println("hello!")
	log.Print("The logs here.")
	log.Print("Base URL is ", baseURL)
	svc := service.PSService{}
	log.Print(svc)
	//controller.mux = http.NewServeMux()
	//newController := PSController{*http.NewServeMux(), 2, "hello"}
	newController := &PSController{}
	newController.mux = http.NewServeMux()
	newController.num = 10
	newController.str = "testing here..."
	log.Print(*newController)
	return newController
}

type PSController struct {
	mux *http.ServeMux
	num int
	str string
}
