package dao

import (
	"gopkg.in/mgo.v2/bson"
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

// Gets many documents matched by IDs into provided array.
// Fields is an array of fields to be fetched.
// Objs must be a pointer to an empty array of structs.
func (dao *DAO) GetManyAs(objs interface{}, ids []bson.ObjectId, fields ...string) (err error) {

	q := M{"_id": M{"$in": ids}}
	err = dao.Coll.Find(q).Select(M{}.Select(fields...)).All(objs)
	return
}

// Gets all documents.
// Fields is an array of fields to be fetched.
// Objs must be a pointer to an empty array of structs.
func (dao *DAO) GetAllAs(objs interface{}, fields ...string) (err error) {

	err = dao.Coll.Find(M{}).Select(M{}.Select(fields...)).All(objs)
	return
}

// Gets all documents.
// Fields is an array of fields to be fetched.
func (dao *DAO) GetAllAsMap(fields ...string) (res []map[string]interface{}, err error) {

	err = dao.Coll.Find(M{}).Select(M{}.Select(fields...)).All(&res)
	return
}
