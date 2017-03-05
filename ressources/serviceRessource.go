package ressources

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const endpoint = "/service"

// Structure representing the Service endpoint
type service struct{}

// NewServiceRessource provide a new instance of Service
func NewServiceRessource() *service {
	return &service{}
}

func (s *service) RegisterRessource(router *httprouter.Router) {
	router.GET(endpoint, getAllServicesInstances)
	router.GET(endpoint+"/:name", getServiceInstanceByName)
	router.POST(endpoint, registerServiceInstance)
	router.PUT(endpoint+"/:id", updateServiceInstance)
}

func getAllServicesInstances(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "getAllServicesInstances ! ")
}

func getServiceInstanceByName(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "getServiceInstanceByName ! "+params.ByName("name"))
}

func registerServiceInstance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "registerServiceInstance ! ")
}

func updateServiceInstance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "updateServiceInstance ! "+params.ByName("id"))
}
