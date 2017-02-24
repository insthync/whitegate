package mongodb

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2"
	"suriyun.com/suriyun/whitedoor/configurations"
)

var session *mgo.Session
var config *configurations.AppConfiguration

// Init ... Initialize mongodb instance
func Init(initConfig *configurations.AppConfiguration) (*mgo.Session, error) {
	config = initConfig
	dialInfo := mgo.DialInfo{
		Addrs:    []string{config.DatabaseHost},
		Timeout:  30 * time.Second,
		Database: config.DatabaseName,
		Username: config.DatabaseUsername,
		Password: config.DatabasePassword,
	}

	fmt.Println("Connecting to MongoDB server...")
	var err error
	session, err = mgo.DialWithInfo(&dialInfo)
	if err != nil {
		return session, err
	}

	session.SetMode(mgo.Monotonic, true)
	fmt.Println("MongoDB was connected successfully.")
	return session, nil
}

// GetSession ... Get session to access database
func GetSession() *mgo.Session {
	return session
}

// GetDatabase ... Get default database to access collections
func GetDatabase() *mgo.Database {
	return GetSession().DB(config.DatabaseName)
}

// GetCollection ... Get collection to manage data
func GetCollection(name string) *mgo.Collection {
	return GetDatabase().C(name)
}
