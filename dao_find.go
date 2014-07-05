package dao

//------------------------------------------------------------
// DAO find methods
//------------------------------------------------------------

// Deletes object from collection by given criteria.
// Kvals is an array of key-value pairs like so:
// "name", "Joe", "age", 99, ...
func (dao *DAO) FindAs(obj interface{}, criteria map[string]interface{}, fields ...string) (err error) {

	err = dao.Coll.Find(criteria).Select(M{}.Select(fields...)).One(obj)
	return
}
