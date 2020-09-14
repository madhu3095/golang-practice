package main

import (
	"awesomeProject/controllers"
	"net/http"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(
		":3000",
		nil,
	)
}
