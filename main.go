package main

import(
	//"RestApp/model"
	"RestApp/dao"
	"RestApp/controller"
	"net/http"
	"fmt"
)

func main(){
	dao.Init()
	http.ListenAndServe(":8081",controller.Handlers())
	fmt.Printf("Server Listening on //localhost:%d\n",8081)
}