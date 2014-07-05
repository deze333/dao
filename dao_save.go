package dao

//------------------------------------------------------------
// DAO save methods
//------------------------------------------------------------

// Saves element
func (dao *DAO) Save(obj interface{}) (err error) {

	err = dao.Coll.Insert(obj)
	return
}
