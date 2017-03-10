package ressources

import (
	"encoding/json"
	"go-service-discovery/datastores"
	"net/http"

	"time"

	"log"

	"github.com/julienschmidt/httprouter"
)

const endpoint = "/service"

//service Structure representing the Service endpoint
type service struct {
	dataStore datastores.ServiceDatastore
}

// NewServiceRessource provide a new instance of Service
func NewServiceRessource() *service {
	return &service{datastores.NewServiceDatastore()}
}

// Register ressource's endpoint to the httprouter
func (s *service) RegisterRessource(router *httprouter.Router) {
	router.GET(endpoint, s.getAllServicesInstances)
	router.GET(endpoint+"/:name", s.getServiceInstanceByName)
	router.POST(endpoint, s.registerServiceInstance)
	router.PUT(endpoint+"/:id", s.updateServiceInstance)

	go doEvery(10*time.Minute, s.cleanUpDataStore)
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

	if instance.TimestampRegistry.IsZero() {
		instance.TimestampRegistry = time.Now()
	}

	s.dataStore.AddServiceInstance(instance)
	w.WriteHeader(http.StatusOK)
}

func (s *service) updateServiceInstance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var instance datastores.ServiceInstance
	err := decoder.Decode(&instance)
	if err != nil {
		panic(err)
	}

	instance.TimestampRegistry = time.Now()

	s.dataStore.UpdateServiceInstance(instance)
	w.WriteHeader(http.StatusOK)
}

func (s *service) cleanUpDataStore() {
	log.Println("Cleaning round")
	instances := s.dataStore.GetAllServicesInstances()

	for position, instance := range instances {
		difference := instance.TimestampRegistry.Sub(time.Now())

		if difference > 59 {
			log.Printf("Cleaning instance ID: %s", instance.ID)
			s.dataStore.RemoveServiceInstance(position)
		}
	}
}

func doEvery(d time.Duration, functionToExecute func()) {
	for range time.Tick(d) {
		functionToExecute()
	}
}
