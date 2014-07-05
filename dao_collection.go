package dao

//------------------------------------------------------------
// DAO collection methods
//------------------------------------------------------------

// Drops collection. No turning back!
func (dao *DAO) DropCollection(are, you, sure bool) (err error) {

	if are && you && sure {
		err = dao.Coll.DropCollection()

		// Make sure indexes will get rebuild next time
		dao.expireIndexes()
		dao.sess.ResetIndexCache()
	}

	return
}
