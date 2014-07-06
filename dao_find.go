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

// Finds many objects matching dateField to the specified time period.
// Objs must be a pointer to an empty array of structs.
func (dao *DAO) FindManyByPeriodAs(objs interface{}, dateField string, ps, pe time.Time, fields ...string) (err error) {

	q := M{dateField: M{"$gte": ps, "$lt": pe}}
	err = dao.Coll.Find(q).Select(M{}.Select(fields...)).All(objs)
	return
}
