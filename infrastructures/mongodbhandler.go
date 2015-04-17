package infrastructures

import (
	"gopkg.in/mgo.v2"    
    "log"	
	"time"
)

const (
	MongoDBHosts = "127.0.0.1:27017"	
	Database = "bookmanagement"
	BooksCollection = "books"
	CategoriesCollection = "categories"
)

type MongoDbHandler struct {
	MongoSession *mgo.Session
}

func NewMongoDbHandler() *MongoDbHandler {
	// We need this object to establish a session to our MongoDB.
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{MongoDBHosts},
		Timeout:  60 * time.Second,
		Database: Database,
	}
 
	// Create a session which maintains a pool of socket connections
	// to our MongoDB.
	mongoSession, errMongo := mgo.DialWithInfo(mongoDBDialInfo)
	if errMongo != nil {
		log.Fatalf("CreateSession: %s\n", errMongo)
	}	
	mongoSession.SetMode(mgo.Monotonic, true)

	return &MongoDbHandler{ MongoSession:  mongoSession}
}