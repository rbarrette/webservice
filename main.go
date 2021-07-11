package main

import (
	"net/http"

	"github.com/rbarrette/webservice/controllers"
)

func main() {

	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)

}
