package main

import (
	"fmt"
	"net/http"

	"assignment/user/app/controller"
	"assignment/user/app/routes"
	"assignment/user/app/service"
)

func main() {
	goc, err := service.GetSession()
	if err != nil {
		fmt.Println("mongodb connection error")
		panic(err)
	}
	defer goc.Close()
	userDB := service.NewUserDbClient(goc)
	ctrl := controller.NewUserRepository(userDB)
	rout := routes.NewRouter(ctrl)
	fmt.Println("server started on PORT 8070")
	http.ListenAndServe(":8070", rout)
}
