package storage

import "gql_app/graph/model"

func SelectProducts() (products []*model.Product, err error) {
	err = DB.Model(&products).Select()
	if err != nil {
		return
	}
	return
}
