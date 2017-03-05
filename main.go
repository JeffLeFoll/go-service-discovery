package main

import (
	"go-service-discovery/ressources"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	restRouter := httprouter.New()

	serviceCtrl := ressources.NewServiceRessource()
	serviceCtrl.RegisterRessource(restRouter)

	log.Println("serving on http://localhost:7777/service")
	http.ListenAndServe("localhost:7777", restRouter)
}
