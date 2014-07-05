package dao

import "labix.org/v2/mgo/bson"

//------------------------------------------------------------
// DAO update array methods
//------------------------------------------------------------

// Adds element to array.
// Key identifies array.
func (dao *DAO) Update_ArrayAdd(id bson.ObjectId, key string, obj interface{}) (err error) {

	err = dao.Coll.UpdateId(id, M{"$push": M{key: obj}})
	return
}
