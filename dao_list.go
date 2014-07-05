package dao

//------------------------------------------------------------
// DAO get list methods
//------------------------------------------------------------

// Gets all documents.
// Fields is an array of fields to be fetched.
func (dao *DAO) GetAll(fields ...string) (res []map[string]interface{}, err error) {

	err = dao.Coll.Find(M{}).Select(M{}.Select(fields...)).All(&res)
	return
}
