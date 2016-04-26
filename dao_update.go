package dao

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

//------------------------------------------------------------
// DAO update methods
//------------------------------------------------------------

// Sets single field to given value.
// Provide empty string to unset the field regardless of field's type.
func (dao *DAO) Update_Set(id bson.ObjectId, key string, obj interface{}) (err error) {

	rval := reflect.ValueOf(obj)
	switch rval.Kind() {

	case reflect.String:
		// String: "" means unset
		str := rval.String()
		//fmt.Println("Str =", str, ", len =", len(str))
		if str == "" {
			//fmt.Println("*** DAO SET = UNSET", dao.Coll.FullName, key)
			err = dao.Coll.UpdateId(id, M{"$unset": M{key: ""}})
		} else {
			err = dao.Coll.UpdateId(id, M{"$set": M{key: obj}})
		}

	default:
		err = dao.Coll.UpdateId(id, M{"$set": M{key: obj}})
	}
	return
}

// Updates (set/unset) by map.
// Nil means don't touch.
//
// For string values:
// Empty string means unset.
// Not empty string means set.
//
// Objects passed as they are.
func (dao *DAO) Update(id bson.ObjectId, params map[string]interface{}) (err error) {

	sets := M{}
	unsets := M{}

	fmt.Println("*** DAO UPDATE", dao.Coll.FullName)
	for key, val := range params {

		// NOTE: ObjectId is actually a string! We need explicit cast.

		// If that's ObjectId, we're done
		if id, ok := val.(bson.ObjectId); ok {
			//fmt.Println("Key = ", key, "ObjectId =", id)
			sets[key] = id
			continue
		}

		// If that's pointer to ObjectId, we're done
		if id, ok := val.(*bson.ObjectId); ok && id != nil {
			//fmt.Println("Key = ", key, "ObjectId =", id)
			sets[key] = *id
			continue
		}

		// Extensive checking for nil is required b/c interface{} is never nil
		// http://golang.org/doc/faq#nil_error

		rval := reflect.ValueOf(val)
		//fmt.Println("Key = ", key, "Val =", val, "Rval =", rval, "Kind =", rval.Kind())

		switch rval.Kind() {

		case reflect.String:
			// String: "" means unset
			str := rval.String()
			//fmt.Println("Str =", str, ", len =", len(str))
			if str == "" {
				unsets[key] = ""
			} else {
				sets[key] = str
			}

		case reflect.Ptr:
			// Pointer: nil means don't update
			elem := rval.Elem()
			if !elem.IsValid() {
				//fmt.Println("--- Ignoring NIL for", key)
				continue
			} else if elem.Kind() == reflect.String {
				// Pointer to string ?
				str := elem.String()
				if str == "" {
					unsets[key] = ""
				} else {
					sets[key] = str
				}
			} else {
				// Pointer to something else, retrieve object
				//fmt.Println("--- Pointer to NOT string", key, "=", val, "=", elem.Interface())
				sets[key] = elem.Interface()
			}

		default:
			// Pass as is
			sets[key] = val
		}
	}

	q := M{}
	if len(sets) > 0 {
		q["$set"] = sets
		fmt.Println("+++ sets:", sets)
	}
	if len(unsets) > 0 {
		q["$unset"] = unsets
		fmt.Println("xxx unsets:", unsets)
	}

	if len(q) > 0 {
		err = dao.Coll.UpdateId(id, q)
	}
	return
}
