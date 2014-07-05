package dao

import (
	"labix.org/v2/mgo/bson"
)

//------------------------------------------------------------
// DAO get methods
//------------------------------------------------------------

// Gets document into provided object.
// Fields is an array of fields to be fetched.
func (dao *DAO) GetAs(obj interface{}, id bson.ObjectId, fields ...string) (err error) {

	err = dao.Coll.FindId(id).Select(M{}.Select(fields...)).One(obj)
	return
}

// Gets document as a map.
// Fields is an array of fields to be fetched.
func (dao *DAO) GetAsMap(id bson.ObjectId, fields ...string) (obj map[string]interface{}, err error) {

	err = dao.Coll.FindId(id).Select(M{}.Select(fields...)).One(&obj)
	return
}
