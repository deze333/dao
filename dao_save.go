package dao

import (
	"errors"
	"fmt"
	"reflect"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//------------------------------------------------------------
// DAO save methods
//------------------------------------------------------------

// Saves objecst(s). Force ID means generate ID before passing to Mongo.
// Useful if you need this ID right away.
// Method checks if ID already exists and doesn't overwrite it in this case.
// IMPORTANT:
// Force ID mode only supports pointers to structs.
// ObjectId must be a pointer too.
func (dao *DAO) Save(forceId bool, objs ...interface{}) (err error) {
	if forceId {
		err = saveExp(dao.Coll, forceId, objs...)
	} else {
		err = dao.Coll.Insert(objs...)
	}
	return
}

// Experimental implementation of save.
func saveExp(coll *mgo.Collection, forceId bool, objs ...interface{}) (err error) {

	for _, obj := range objs {

		switch reflect.ValueOf(obj).Kind() {
		case reflect.Ptr:
			err = setObjId(obj, forceId)

		case reflect.Struct:
			err = errors.New("Save currently only supports pointers to structs")
			//err = setObjId(&obj, forceId)

		default:
			err = fmt.Errorf("Save does not support objs of type: %v",
				reflect.ValueOf(obj).Kind())
			return
		}
	}

	// Save if all is good
	if coll != nil && err == nil {
		err = coll.Insert(objs...)
	}
	return
}

func setObjId(obj interface{}, forceId bool) (err error) {

	var rval = reflect.ValueOf(obj)
	switch rval.Kind() {

	case reflect.Ptr:
		if !rval.IsValid() {
			err = fmt.Errorf("Element not valid: %v", obj)
			return
		}
		val := rval.Elem()
		typ := rval.Elem().Type()
		fn := findIdField(typ)
		if fn != -1 {
			setIdField(val.Field(fn), forceId)
		} else {
			return fmt.Errorf("Field with tag bson:\"_id\" not found")
		}

	case reflect.Struct:
		typ := reflect.TypeOf(obj)
		fn := findIdField(typ)
		if fn != -1 {
			setIdField(rval.Field(fn), forceId)
		} else {
			return fmt.Errorf("Field with tag bson:\"_id\" not found")
		}
	}

	return
}

func findIdField(typ reflect.Type) int {

	for i := 0; i < typ.NumField(); i++ {
		tag := typ.Field(i).Tag.Get("bson")
		if tag == "_id" || tag == "_id,omitempty" {
			return i
		}
	}
	return -1
}

func setIdField(val reflect.Value, forceId bool) {

	id := bson.NewObjectId()
	switch val.Kind() {

	case reflect.Ptr:
		//fmt.Println("\t\tField is a ptr, isNil =", val.IsNil())
		// Only set if value is nil and force ID creation was requested
		if val.IsNil() && forceId {
			val.Set(reflect.ValueOf(&id))
		}

	case reflect.Struct:
		// Only as an exercise, not really supported
		fmt.Println("\t\tField is a struct")
		val.Set(reflect.ValueOf(id))
	}
}
