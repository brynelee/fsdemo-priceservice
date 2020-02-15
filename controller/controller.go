package controller

import (
	"fsdemo-priceservice/service"
	"log"
	"net/http"
)

const baseURL = "/priceservice"

type PSController struct {
	Mux *http.ServeMux
}

func NewBaseController() *PSController {

	log.Print("Calling into the NewBaseController...")

	newController := &PSController{
		http.NewServeMux(),
	}

	//handle "/priceservice"
	newController.Mux.HandleFunc(baseURL, service.HelloRequest)

	return newController
}
