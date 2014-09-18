package dao

import (
	"fmt"

	"labix.org/v2/mgo"
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

// Creates master DB session to given server and stores it for future reference.
// Use getSession() to get a connection to a specific DB/collection.
func connectToServer(params map[string]string) (err error) {

	// Server ID, ensure at least default
	id := params[KEY_DB_SERVER_ID]
	if id == "" {
		id = "_"
	}

	// Params
	server := params[KEY_DB_SERVER]
	options := params[KEY_DB_OPTIONS]
	mode := params[KEY_DB_MODE]
	log := params[KEY_DB_LOG]

	// Validate
	if server == "" {
		err = fmt.Errorf("Mongo DB init aborted due to empty 'server'")
		return
	}

	// Build URL
	url := "mongodb://" + server
	if options != "" {
		url += "/?" + options
	}

	fmt.Println("MongoDB URL      :", url)

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
		fmt.Println("MongoDB Mode     : Eventual")

	case "monotonic":
		sess.SetMode(mgo.Monotonic, true)
		fmt.Println("MongoDB Mode     : Monotonic")

	case "strong":
		sess.SetMode(mgo.Strong, true)
		fmt.Println("MongoDB Mode     : Strong")

	default:
		fmt.Println("MongoDB Mode     : Undefined")
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
func GetCollections(sess *mgo.Session, dbname string) (collnames []string) {
	var err error
	collnames, err = sess.DB(dbname).CollectionNames()
	if err != nil {
		panic(fmt.Sprintf("MongoDB cannot list collection names for '%v' due to error: %v", dbname, err.Error()))
	}
	return
}

// Checks if given collection exists.
func IsCollectionExists(sess *mgo.Session, dbname, collname string) bool {
	colls := GetCollections(sess, dbname)
	for _, name := range colls {
		if name == collname {
			return true
		}
	}
	return false
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
