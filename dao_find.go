package dao

import "time"

//------------------------------------------------------------
// DAO find methods
//------------------------------------------------------------

// Finds an object from collection by given equals 'and' criteria.
// Obj must be a pointer to a struct.
func (dao *DAO) FindAs(obj interface{}, equals map[string]interface{}, fields ...string) (err error) {

	err = dao.Coll.Find(equals).Select(M{}.Select(fields...)).One(obj)
	return
}

// Finds many objects matching equals 'and' criteria.
// Objs must be a pointer to an empty array of structs.
func (dao *DAO) FindManyAs(objs interface{}, equals map[string]interface{}, fields ...string) (err error) {

	err = dao.Coll.Find(equals).Select(M{}.Select(fields...)).All(objs)
	return
}

// Finds many objects matching dateKey to the specified time period.
// Note that interval is [ps, pe) meaning pe is EXCLUDED.
// Objs must be a pointer to an empty array of structs.
func (dao *DAO) FindManyByIntervalAs(objs interface{}, dateKey string, ps, pe time.Time, fields ...string) (err error) {

	q := M{dateKey: M{"$gte": ps, "$lt": pe}}
	err = dao.Coll.Find(q).Select(M{}.Select(fields...)).All(objs)
	return
}

// Finds many objects matching 'ps' and 'pe' field keys to be inside the specified time period.
// Objs must be a pointer to an empty array of structs.
func (dao *DAO) FindManyByPeriodAs(objs interface{}, psKey, peKey string, ps, pe time.Time, fields ...string) (err error) {

	q := M{psKey: M{"$gte": ps}, peKey: M{"$lte": pe}}
	err = dao.Coll.Find(q).Select(M{}.Select(fields...)).All(objs)
	return
}
