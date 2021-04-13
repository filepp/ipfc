package inspection

import "ipfc/dbstore/ds"

type Inspector struct {
}

func NewInspector(peerId, addr string, db *ds.DbStore) *Inspector {
	return &Inspector{}
}

func (i *Inspector) Run() {

}

func (i *Inspector) Stop() {

}
