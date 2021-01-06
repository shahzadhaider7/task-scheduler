package db

import (
	"log"

	"github.com/shahzadhaider7/task-scheduler/models"
)

// We want to add the functionality of : add, update, get, remove, list
// DataStore is an interface for query ops
type DataStore interface {
	AddTask(task *models.Task) (string, error)
	GetTask(id string) (*models.Task, error)
	DeleteTask(id string) error
	UpdateTask(task *models.Task) error
}

// Option holds configuration for data store clients
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
