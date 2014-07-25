package dao

//------------------------------------------------------------
// DAO upsert methods
//------------------------------------------------------------

// Upserts.
func (dao *DAO) Upsert(selector, update interface{}) (err error) {

	_, err = dao.Coll.Upsert(selector, update)
	return
}
