package datastores

import "log"

//ServiceInstance Structure representing a Service Instance
type ServiceInstance struct {
	ID   string
	Name string
	URL  string
}

// ServiceDatastore structure representing the datastore
type ServiceDatastore struct {
	repository []ServiceInstance
}

// NewServiceDatastore provide a new instance of ServiceDatastore
func NewServiceDatastore() ServiceDatastore {
	log.Println("NewServiceDatastore")
	return ServiceDatastore{make([]ServiceInstance, 0)}
}

// GetServiceInstanceByName return the Service Instance with the given name
func (d *ServiceDatastore) GetServiceInstanceByName(name string) (instance ServiceInstance) {
	for _, instance := range d.repository {
		if instance.Name == name {
			return instance
		}
	}
	return
}

// GetAllServicesInstances return all the services instances
func (d *ServiceDatastore) GetAllServicesInstances() []ServiceInstance {
	instances := make([]ServiceInstance, len(d.repository))
	copy(instances, d.repository)

	return instances
}

//AddServiceInstance Add a new ServiceInstance in the datastore
func (d *ServiceDatastore) AddServiceInstance(instance ServiceInstance) {
	d.repository = append(d.repository, instance)
}
