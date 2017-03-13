package ressources

import (
	"encoding/json"
	"go-service-discovery/datastores"
	"net/http"

	"time"

	"log"

	"github.com/julienschmidt/httprouter"
	uuid "github.com/satori/go.uuid"
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
	router.DELETE(endpoint+"/:id", s.unRegisterServiceInstance)

	go s.doEvery(1*time.Minute, s.cleanUpDataStore)
}

func (s *service) getAllServicesInstances(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	instances := s.dataStore.GetAllServicesInstances()
	json.NewEncoder(w).Encode(instances)
}

func (s *service) getServiceInstanceByName(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	instance := s.dataStore.GetServiceInstanceByName(params.ByName("name"))
	json.NewEncoder(w).Encode(instance)
}

func (s *service) getServiceInstanceByID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	instance := s.dataStore.GetServiceInstanceByID(params.ByName("name"))
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

	u2, err := uuid.FromString(instance.ID)
	if err != nil {
		u2 = uuid.NewV4()
	}
	instance.ID = u2.String()

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

	s.dataStore.UpdateServiceInstance(params.ByName("id"), instance)
	w.WriteHeader(http.StatusOK)
}

func (s *service) unRegisterServiceInstance(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	instances := s.dataStore.GetAllServicesInstances()

	for position, instance := range instances {
		if instance.ID == params.ByName("id") {
			s.dataStore.RemoveServiceInstance(position)
		}
	}
}

func (s *service) cleanUpDataStore() {
	instances := s.dataStore.GetAllServicesInstances()

	for position, instance := range instances {
		difference := time.Now().Sub(instance.TimestampRegistry) / time.Second

		if difference > 59 {
			log.Printf("Cleaning instance ID: %s, Name: %s", instance.ID, instance.Name)
			s.dataStore.RemoveServiceInstance(position)
		}
	}
}

func (s *service) doEvery(d time.Duration, functionToExecute func()) {
	for range time.Tick(d) {
		functionToExecute()
	}
}
