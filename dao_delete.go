package dao

import (
	"fmt"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//------------------------------------------------------------
// DAO delete methods
//------------------------------------------------------------

// Deletes single object from collection.
func (dao *DAO) Delete(id bson.ObjectId) (err error) {

	err = dao.Coll.RemoveId(id)
	return
}

// Deletes many objects by IDs.
func (dao *DAO) DeleteMany(ids []bson.ObjectId) (count int, err error) {

	q := M{"_id": M{"$in": ids}}

	var info *mgo.ChangeInfo
	info, err = dao.Coll.RemoveAll(q)
	if info != nil {
		count = info.Removed
	}
	return
}

// Deletes object from collection by given criteria.
// Kvals is an array of key-value pairs like so:
// "name", "Joe", "age", 99, ...
func (dao *DAO) DeleteBy(kvals ...interface{}) (count int, err error) {

	q := M{}

	num := len(kvals)
	for i := 0; i < num; i += 2 {
		q[fmt.Sprint(kvals[i])] = kvals[i+1]
	}

	var info *mgo.ChangeInfo
	info, err = dao.Coll.RemoveAll(q)
	if info != nil {
		count = info.Removed
	}
	return
}

// Deletes all objects from collection.
func (dao *DAO) DeleteAll(are, you, sure bool) (count int, err error) {

	var info *mgo.ChangeInfo
	info, err = dao.Coll.RemoveAll(M{})
	if info != nil {
		count = info.Removed
	}
	return
}
