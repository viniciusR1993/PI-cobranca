package main

import (
	"fatec/routes"
	"fatec/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func init() {
	utils.LoadTemplates()
}

func main() {
	rtr := mux.NewRouter()
	routes.LoadRoutes(rtr)
	http.ListenAndServe(":8000", rtr)
}
