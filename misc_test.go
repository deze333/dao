package dao

import (
	"fmt"
	"testing"

	"labix.org/v2/mgo/bson"
)

func TestSave(t *testing.T) {
	var err error

	// Nil must produce error
	err = saveExp(nil, true, nil)
	if err != nil {
		fmt.Println("Nil produces error:", err)
	}

	// As pointer, ptr ID
	type Obj struct {
		Id   *bson.ObjectId `bson:"_id"`
		Name string
	}

	obj := Obj{
		Name: "pointer ID, ensure ID",
	}

	fmt.Println("\n---------------")
	err = saveExp(nil, true, &obj)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(obj)

	obj = Obj{
		Name: "pointer ID, don't ensure ID",
	}

	fmt.Println("\n---------------")
	err = saveExp(nil, false, &obj)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(obj)

	// As struct, non ptr ID
	// Doesn't make much sense but seems to work
	type ObjNoPtr struct {
		Id   bson.ObjectId `bson:"_id"`
		Name string
	}

	objNoPtr := ObjNoPtr{
		Name: "non-pointer ID",
	}

	fmt.Println("\n---------------")
	err = saveExp(nil, true, objNoPtr)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(objNoPtr)

	return

	objs := []*Obj{
		&Obj{nil, "b"},
		&Obj{nil, "c"},
		&Obj{nil, "d"},
	}

	err = saveExp(nil, true, objs)
	if err != nil {
		t.Error(err)
	}

	err = saveExp(nil, true, obj, obj, obj)
	if err != nil {
		t.Error(err)
	}

	/*
		coll := mgo.Collection{}
		coll.Insert(objs, objs)
		coll.Insert(objs)
		coll.Insert(objs...)
	*/
}
