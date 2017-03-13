package datastores

import "time"

//ServiceInstance Structure representing a Service Instance
type ServiceInstance struct {
	ID                string
	Name              string
	URL               string
	TimestampRegistry time.Time
}

// ServiceDatastore structure representing the datastore
type ServiceDatastore struct {
	repository []ServiceInstance
}

// NewServiceDatastore provide a new instance of ServiceDatastore
func NewServiceDatastore() ServiceDatastore {
	return ServiceDatastore{make([]ServiceInstance, 0)}
}

// GetServiceInstanceByName return all the Services Instances with the given name
func (d *ServiceDatastore) GetServiceInstanceByName(name string) (instancesSelected []ServiceInstance) {
	instancesSelected = make([]ServiceInstance, 0)

	for _, instance := range d.repository {
		if instance.Name == name {
			instancesSelected = append(instancesSelected, instance)
		}
	}
	return instancesSelected
}

// GetServiceInstanceByID return the Service Instance with the given id
func (d *ServiceDatastore) GetServiceInstanceByID(id string) (instance ServiceInstance) {
	for _, instance := range d.repository {
		if instance.ID == id {
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

//RemoveServiceInstance Remove a service instance from the data store
func (d *ServiceDatastore) RemoveServiceInstance(position int) {
	d.repository = append(d.repository[:position], d.repository[position+1:]...)
}

//UpdateServiceInstance Add a new ServiceInstance in the datastore
func (d *ServiceDatastore) UpdateServiceInstance(id string, updatedInstance ServiceInstance) {

	for position, instance := range d.repository {
		if instance.ID == id {
			d.repository[position] = updatedInstance
			return
		}
	}
}
