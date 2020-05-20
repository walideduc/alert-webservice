package main

import (
	"fmt"
	"net/http"

	"gitlab.accor.net/accor/sf/monitoring/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
	fmt.Println("Yes")
}
