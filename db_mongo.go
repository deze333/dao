package dao

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2"
)

//------------------------------------------------------------
// DB Mongo connector
//------------------------------------------------------------

const (
	// Parameter keys
	KEY_DB_SERVER_ID = "id"
	KEY_DB_SERVER    = "server"
	//KEY_DB_PORT       = "port"
	KEY_DB_OPTIONS = "options"
	KEY_DB_MODE    = "mode"
	KEY_DB_LOG     = "log"
)

//------------------------------------------------------------
// Vars
//------------------------------------------------------------

var (
	_dbServers = map[string]*mgo.Session{}
)

//------------------------------------------------------------
// Implementation
//------------------------------------------------------------

// Session getter returns a copy of the master session.
func getSession(servername, dbname, collname string) (sess *mgo.Session, coll *mgo.Collection) {

	if masterSess, ok := _dbServers[servername]; ok {
		sess = masterSess.Copy()
		coll = sess.DB(dbname).C(collname)
	}
	return
}

// Session getter returns a copy of the master session.
func getDb(servername, dbname string) (db *mgo.Database) {

	if masterSess, ok := _dbServers[servername]; ok {
		db = masterSess.Copy().DB(dbname)
	}
	return
}

// Creates master DB session to given server and stores it for future reference.
// Use getSession() to get a connection to a specific DB/collection.
func connectToServer(params map[string]string) (err error) {

	// Server ID, ensure at least default
	id := params[KEY_DB_SERVER_ID]
	if id == "" {
		id = "_"
	}

	// Is server with this ID already connected?
	if _, ok := _dbServers[id]; ok {
		err = fmt.Errorf("Mongo DB server already connected: %v", id)
		return
	}

	// Params
	server := params[KEY_DB_SERVER]
	options := params[KEY_DB_OPTIONS]
	mode := params[KEY_DB_MODE]
	log := params[KEY_DB_LOG]

	// Validate
	if server == "" {
		err = fmt.Errorf("Mongo DB init aborted due to empty 'server' key")
		return
	}

	// Build URL
	url := "mongodb://" + server
	if options != "" {
		url += "/?" + options
	}

	fmt.Println("[dao] MongoDB URL      :", url)

	// Enable logging ?
	if log != "" {
		mgo.SetDebug(true)
		mgo.SetLogger(&MongoLogger{})
	}

	// Dial server
	var sess *mgo.Session
	if sess, err = mgo.Dial(url); err != nil {
		err = fmt.Errorf("Error connecting to MongoDB: %v", err)
		return
	}

	// Set Mongo mode
	switch mode {
	case "eventual":
		sess.SetMode(mgo.Eventual, true)
		fmt.Println("[dao] MongoDB Mode     : Eventual")

	case "monotonic":
		sess.SetMode(mgo.Monotonic, true)
		fmt.Println("[dao] MongoDB Mode     : Monotonic")

	case "strong":
		sess.SetMode(mgo.Strong, true)
		fmt.Println("[dao] MongoDB Mode     : Strong")

	default:
		fmt.Println("[dao] MongoDB Mode     : Undefined")
	}

	// Remember connection
	_dbServers[id] = sess
	return
}

// Disconnects all servers.
func disconnectServers() {
	for _, sess := range _dbServers {
		sess.Close()
	}
	_dbServers = map[string]*mgo.Session{}
}

//------------------------------------------------------------
// Collections operations
//------------------------------------------------------------

// Lists all collections present.
func GetCollections(servername, dbname string) (collnames []string, err error) {
	if db := getDb(servername, dbname); db != nil {
		collnames, err = db.CollectionNames()
	} else {
		err = errors.New("Cannot acquire Mongo DB")
	}
	return
}

// Checks if given collection exists.
func IsCollectionExists(servername, dbname, collname string) (exists bool, err error) {
	var colls []string
	colls, err = GetCollections(servername, dbname)
	if err != nil {
		return
	}
	for _, name := range colls {
		if name == collname {
			exists = true
			return
		}
	}
	return
}

//------------------------------------------------------------
// Optional Mgo logger
//------------------------------------------------------------

// Logger for Mgo
type MongoLogger struct{}

// Mgo MongoLogger interface implementation.
func (l *MongoLogger) Output(calldepth int, s string) error {
	fmt.Println("[MGO]", calldepth, s)
	return nil
}
