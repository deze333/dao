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

// Removes pullObj element from pullFrom array.
// Adds pushObj element to pushTo array.
func (dao *DAO) Update_ArraysPullPush(id bson.ObjectId, pullFrom string, pullObj interface{}, pushTo string, pushObj interface{}) (err error) {

	q := M{
		"$pull": M{pullFrom: pullObj},
		"$push": M{pushTo: pushObj},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}

// Removes pullFrom array element that matches pullObj.
func (dao *DAO) Update_ArrayPull(id bson.ObjectId, pullFrom string, pullObj interface{}) (err error) {

	q := M{
		"$pull": M{pullFrom: pullObj},
	}
	err = dao.Coll.UpdateId(id, q)
	return
}
