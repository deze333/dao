package dao

import (
	"fmt"

	"labix.org/v2/mgo/bson"
)

//------------------------------------------------------------
// DAO delete methods
//------------------------------------------------------------

// Deletes object from collection.
func (dao *DAO) Delete(id bson.ObjectId) (err error) {

	err = dao.Coll.RemoveId(id)
	return
}

// Deletes object from collection by given criteria.
// Kvals is an array of key-value pairs like so:
// "name", "Joe", "age", 99, ...
func (dao *DAO) DeleteBy(kvals ...interface{}) (err error) {

	q := M{}

	num := len(kvals)
	for i := 0; i < num; i += 2 {
		q[fmt.Sprint(kvals[i])] = kvals[i+1]
	}

	err = dao.Coll.Remove(q)
	return
}
