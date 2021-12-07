package storage

import (
	"gql_app/graph/model"
)

func (db *Psql) SelectProducts() (products []*model.Product, err error) {
	err = db.DB.Model(&products).Select()
	if err != nil {
		return
	}

	return
}
