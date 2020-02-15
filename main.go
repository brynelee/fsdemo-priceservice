package main

import (
	"fsdemo-priceservice/config"
	"fsdemo-priceservice/controller"
	"log"
	"net/http"
	"time"
)

/*func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}*/

var appConfig *config.Configuration

func init() {

	appConfig = config.LoadConfig()

}

func main() {

	log.Println("fsdemo-priceservice is loading in main().")

	cntl := controller.NewBaseController()

	server := &http.Server{
		Addr:           appConfig.Address,
		Handler:        cntl.Mux,
		TLSConfig:      nil,
		ReadTimeout:    time.Duration(appConfig.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration((appConfig.WriteTimeout * int64(time.Second))),
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Price Service will start to on ", appConfig.Address)
	_ = server.ListenAndServe()

}
