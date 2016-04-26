package dao

import "gopkg.in/mgo.v2"

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
