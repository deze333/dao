package dao_app

import (
	"labix.org/v2/mgo/bson"
)

//------------------------------------------------------------
// User model
//------------------------------------------------------------

type User struct {
	Id    *bson.ObjectId `bson:"_id,omitempty"         json:"id,omitempty"`
	Email string         `bson:"email,omitempty"       json:"email,omitempty"`
	Name  string         `bson:"name,omitempty"        json:"name,omitempty"`
}
