package dao

//------------------------------------------------------------
// DAO count methods
//------------------------------------------------------------

// Counts all collections documents.
func (dao *DAO) CountAll() (n int, err error) {

	n, err = dao.Coll.Find(M{}).Count()
	return
}
