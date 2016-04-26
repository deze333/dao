package dao_app

import (
	adao "github.com/deze333/dao"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//------------------------------------------------------------
// DB 'Example One': User DAO
//------------------------------------------------------------

type UserDao struct {
	adao.DAO
}

//------------------------------------------------------------
// Collection, Indexes
//------------------------------------------------------------

const (
	COLL_USERS = "users"
)

var _idxsUsers = adao.BuildIndexes(
	mgo.Index{
		Key: []string{"email"},
		//Unique: true,
		//DropDups: true,
		//Background: true,
		//Sparse: true,
	},
)

//------------------------------------------------------------
// Methods
//------------------------------------------------------------

// Open.
func (dao *UserDao) Open() {
	dao.DAO.Open(SERVER, DB, COLL_USERS, _idxsUsers)
}

//------------------------------------------------------------
// Methods
//------------------------------------------------------------

// Get
func (dao *UserDao) Get(id bson.ObjectId, fields ...string) (res *User, err error) {

	err = dao.Coll.FindId(id).Select(M{}.Select(fields...)).One(&res)
	return
}

// Find by email.
func (dao *UserDao) FindByEmail(email string) (res *User, err error) {

	// Query
	q := bson.M{"email": email}

	// Execute
	err = dao.Coll.Find(q).One(&res)
	return
}

// Get by ids.
func (dao *UserDao) GetByIds(ids []bson.ObjectId, fields ...string) (res []*User, err error) {

	// Result
	res = []*User{}

	q := M{"_id": M{"$in": ids}}

	err = dao.Coll.Find(q).Select(M{}.Select(fields...)).All(&res)
	return
}
