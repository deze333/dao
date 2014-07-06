package dao

import (
	"fmt"
	"reflect"

	"labix.org/v2/mgo/bson"
)

//------------------------------------------------------------
// DAO update methods
//------------------------------------------------------------

// Sets single field to given value.
func (dao *DAO) Update_Set(id bson.ObjectId, key string, obj interface{}) (err error) {

	err = dao.Coll.UpdateId(id, M{"$set": M{key: obj}})
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

	//fmt.Println("*** UPDATE")
	for key, val := range params {

		// Extensive checking for nil is required b/c interface{} is never nil
		// http://golang.org/doc/faq#nil_error

		rval := reflect.ValueOf(val)
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
			}

		default:
			// Pass as is
			sets[key] = val
		}
	}

	fmt.Println("+++ sets:\n", sets)
	fmt.Println("xxx unsets:\n", unsets)

	err = dao.Coll.UpdateId(id, M{"$set": sets, "$unset": unsets})
	return
}
