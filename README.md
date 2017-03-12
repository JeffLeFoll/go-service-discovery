# go-service-discovery
A naive implementation of the Service Discovery pattern in Go.  
Store in-memory a list of service instance.  
Every minute the server clean up service instance that are more than 59 seconds old.  

## Service instance
The registry manage Service instance.
A service instance carry the following information :
```Go
type ServiceInstance struct {
	ID                string  // UUID v4
	Name              string
	URL               string
	TimestampRegistry time.Time
}
```
ID and TimestampRegistry are automatically set at creation time if not provided.  


## Endpoints
Currently the registry offer the following endpoints :  
GET /service  => Return all services instances registered  
GET /service/:name => Return services instances with the given name  
POST /service => register a new service instance, return the newly created ID or the provided ID if it's a valid UUID v4  
PUT /service/:id => update the TimestampRegistry of an already existing service instance
