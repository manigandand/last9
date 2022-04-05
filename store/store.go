package store

import (
	"last9/config"
	"last9/store/adaptee/sqlite"
	"last9/store/adapter"
	"log"
)

// Store global store connection interface
var Store adapter.Store

// Init loads the sample data and prepares the store layer
func Init() {
	// store inmemory adapter ...
	switch config.DBType {
	case "sqlite":
		Store = sqlite.NewAdapter()
		// case "cloudsqlpostgres":
	}
	if Store == nil {
		log.Fatalf("🦠store initialize failed 👎")
	}
	log.Println("Inited Store...👍")
}
