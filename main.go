package main

import(
	//"RestApp/model"
	"RestApp/dao"
	"RestApp/controller"
	"github.com/gorilla/handlers"
	"net/http"
	"fmt"
	//"os"
)

func main(){
	dao.Init()
	fmt.Printf("Server Listening on //localhost:%d\n",8081)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization"})
	methodsOk := handlers.AllowedMethods([]string{"GET","POST","DELETE","OPTIONS","PUT"})
	originOk := handlers.AllowedOrigins([]string{"http://localhost:3001"})
	//http.ListenAndServe(":8081",controller.Handlers())
	http.ListenAndServe(":8081",handlers.CORS(originOk,headersOk,methodsOk)(controller.Handlers()))
}

// fsdfsdf

// 