package dao

import "labix.org/v2/mgo"

//------------------------------------------------------------
// Collection Index model
//------------------------------------------------------------

type CollIndexes struct {
	indexes   []mgo.Index
	isIndexed bool
}

//------------------------------------------------------------
// Methods
//------------------------------------------------------------

// Builds struct contaning an array of Mongo indexes.
// Indexes set to un-initialized.
func BuildIndexes(midxs ...mgo.Index) (idxs *CollIndexes) {

	idxs = &CollIndexes{
		indexes: make([]mgo.Index, len(midxs)),
	}

	for i, midx := range midxs {
		idxs.indexes[i] = midx
	}

	return
}

func (ci *CollIndexes) Indexes() []mgo.Index {
	return ci.indexes
}
