package dao

import (
	"errors"
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
// XXX Still not 100% how objs... work. Better to save one obj at a time.
func (dao *DAO) Save(forceID bool, objs ...interface{}) (err error) {
	if forceID {
		err = saveExp(dao.Coll, forceID, objs...)
	} else {
		for _, obj := range objs {
			err = dao.Coll.Insert(obj)
		}
	}
	return
}

// Experimental implementation of save.
func saveExp(coll *mgo.Collection, forceID bool, objs ...interface{}) (err error) {

	for _, obj := range objs {

		switch reflect.ValueOf(obj).Kind() {
		case reflect.Ptr:
			err = setObjId(obj, forceID)

		case reflect.Struct:
			err = errors.New("Save currently only supports pointers to structs")
			//err = setObjId(&obj, forceID)

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

func setObjId(obj interface{}, forceID bool) (err error) {

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
			setIdField(val.Field(fn), forceID)
		} else {
			return fmt.Errorf("Field with tag bson:\"_id\" not found")
		}

	case reflect.Struct:
		typ := reflect.TypeOf(obj)
		fn := findIdField(typ)
		if fn != -1 {
			setIdField(rval.Field(fn), forceID)
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

func setIdField(val reflect.Value, forceID bool) {

	id := bson.NewObjectId()
	switch val.Kind() {

	case reflect.Ptr:
		//fmt.Println("\t\tField is a ptr, isNil =", val.IsNil())
		// Only set if value is nil and force ID creation was requested
		if val.IsNil() && forceID {
			val.Set(reflect.ValueOf(&id))
		}

	case reflect.Struct:
		// Only as an exercise, not really supported
		fmt.Println("\t\tField is a struct")
		val.Set(reflect.ValueOf(id))
	}
}
