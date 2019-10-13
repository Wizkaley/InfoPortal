package main

import (
	"RESTApp/controller"
	"RESTApp/dao"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

// func main() {
// 	dao.Init()
// 	fmt.Printf("Server Listening on //localhost:%d\n", 8081)
// 	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
// 	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS", "PUT"})
// 	originOk := handlers.AllowedOrigins([]string{"http://localhost:3001"})
// 	//http.ListenAndServe(":8081",controller.Handlers())
// 	http.ListenAndServe(":8081", handlers.CORS(originOk, headersOk, methodsOk)(controller.Handlers()))
// }

func main() {
	session, err := dao.Init("localhost:27017")
	defer session.Close()
	if err != nil {
		log.Printf("Master DB Con Error : %v ", err)
	}
	fmt.Printf("Server Listening on //localhost:%d\n", 8081)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "OPTIONS", "PUT"})
	originOk := handlers.AllowedOrigins([]string{"http://localhost:3001"})
	//http.ListenAndServe(":8081",controller.Handlers())
	http.ListenAndServe(":8081", handlers.CORS(originOk, headersOk, methodsOk)(controller.Handlers(session)))

	//http.ListenAndServe(":8081", controller.Handlers(session))
}
