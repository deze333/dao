package dao

//------------------------------------------------------------
// DAO quirky and unsual methods for testing
//------------------------------------------------------------

// Gets random one.
func (dao *DAO) Quirk_GetOneAs(obj interface{}, fields ...string) (err error) {

	err = dao.Coll.Find(M{}).Select(M{}.Select(fields...)).One(obj)
	return
}
