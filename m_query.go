package dao

//------------------------------------------------------------
// Query convenience type, instead of bson.M
//------------------------------------------------------------

type M map[string]interface{}
type MS []M

//------------------------------------------------------------
// Query helper functions
//------------------------------------------------------------

func (m M) Select(fs ...string) (q M) {

	q = M{}
	for _, f := range fs {
		q[f] = 1
	}
	return
}

func (m M) Unselect(fs ...string) (q M) {

	q = M{}
	for _, f := range fs {
		q[f] = 0
	}
	return
}
