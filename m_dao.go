package dao

import "labix.org/v2/mgo"

//------------------------------------------------------------
// Abstract DAO
//------------------------------------------------------------

type DAO struct {
	server   string
	dbname   string
	collname string
	indexes  *CollIndexes
	sess     *mgo.Session
	Coll     *mgo.Collection
}
