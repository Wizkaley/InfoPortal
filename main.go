package main

import (
	"RESTApp/controller"
	"RESTApp/utils"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
)

func main() {
	//session, err := utils.Init("localhost:27017")
	utils.InitConfig()
	session, err := utils.GetDataBaseSession("10.90.31.148:27017")
	if err != nil {
		log.Printf("Master DB Con Error : %v ", err)
	}
	defer session.Close()

	fmt.Print("Server Listening on //" + utils.Config.DatabaseHost + ":" + strconv.Itoa(utils.Config.DatabasePort))
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS", "PUT"})
	originOk := handlers.AllowedOrigins([]string{"http://localhost:3001"})
	//http.ListenAndServe(":8081",controller.Handlers())
	err = http.ListenAndServe(":8081", handlers.CORS(originOk, headersOk, methodsOk)(controller.Handlers(session, "trial")))
	if err != nil {
		fmt.Println("Error Starting Server : ", err.Error())
	}
}
