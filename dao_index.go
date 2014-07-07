package dao

import (
	"bytes"
	"fmt"

	_D "github.com/deze333/diag"
)

//------------------------------------------------------------
// Collection indexes
//------------------------------------------------------------

// Ensures multiple indexes. Error is suppressed yet reported.
func (dao *DAO) EnsureIndexes() {
	if dao.indexes.isIndexed {
		return
	}

	for _, index := range dao.indexes.indexes {
		err := dao.Coll.EnsureIndex(index)
		if err != nil {
			_D.SOS("db", "Collection set index error", "db", dao.dbname, "coll", dao.collname, "error", err)
		}
	}

	// All indexed
	dao.indexes.isIndexed = true

	// Output indexes
	dao.DebugIndexes()
}

// Sets collection to not indexed. Used in situation of collection drop.
func (dao *DAO) expireIndexes() {
	dao.indexes.isIndexed = false
}

// Debug method that prints active indexes for given collection.
func (dao *DAO) DebugIndexes() {
	indexes, err := dao.Coll.Indexes()
	if err != nil {
		panic(fmt.Sprintf("MongoDB cannot read indexes for '%v' due to error: %v", dao.collname, err.Error()))
		return
	}

	var buf bytes.Buffer
	for _, index := range indexes {
		buf.WriteString(fmt.Sprint(index.Key))
		buf.WriteString(", ")
	}

	// Ideally, indexing only happens once per deployment, hence this NOTE
	_D.NOTE2("DAO "+dao.collname, "idx", buf.String())
}
