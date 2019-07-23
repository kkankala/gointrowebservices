package main

import (
	"fmt"
	"net/http"

	"github.com/kkankala/gointrowebservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	fmt.Println("server listening at 3000")
	http.ListenAndServe(":3000", nil)
}
