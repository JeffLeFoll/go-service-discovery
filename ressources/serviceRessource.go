package ressources

import (
	"encoding/json"
	"fmt"
	"go-service-discovery/datastores"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const endpoint = "/service"

//service Structure representing the Service endpoint
type service struct {
	dataStore datastores.ServiceDatastore
}

// NewServiceRessource provide a new instance of Service
func NewServiceRessource() *service {
	log.Println("NewServiceRessource")
	return &service{datastores.NewServiceDatastore()}
}

// Register ressource's endpoint to the httprouter
func (s *service) RegisterRessource(router *httprouter.Router) {
	router.GET(endpoint, s.getAllServicesInstances)
	router.GET(endpoint+"/:name", s.getServiceInstanceByName)
	router.POST(endpoint, s.registerServiceInstance)
	router.PUT(endpoint+"/:id", s.updateServiceInstance)
}

func (s *service) getAllServicesInstances(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	instances := s.dataStore.GetAllServicesInstances()
	json.NewEncoder(w).Encode(instances)
}

func (s *service) getServiceInstanceByName(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	instance := s.dataStore.GetServiceInstanceByName(params.ByName("name"))
	json.NewEncoder(w).Encode(instance)
}

func (s *service) registerServiceInstance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var instance datastores.ServiceInstance
	err := decoder.Decode(&instance)
	if err != nil {
		panic(err)
	}

	log.Println("AddServiceInstance " + instance.Name)

	s.dataStore.AddServiceInstance(instance)
	w.WriteHeader(http.StatusOK)
}

func (s *service) updateServiceInstance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprint(w, "updateServiceInstance ! "+params.ByName("id"))
}
