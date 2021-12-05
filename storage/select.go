package storage

import (
	"gql_app/graph/model"

	"github.com/go-pg/pg/v10"
)

func SelectProducts() (products []*model.Product, err error) {
	err = DB.Model(&products).Select()
	if err != nil {
		return
	}
	return
}

func SelectUserByID(userID int) (user model.User, err error) {
	user.ID = userID
	err = DB.Model(&user).WherePK().Select()
	if err != nil {
		return
	}
	return
}

func SelectUserByPhone(phone string) (user model.User, err error) {
	err = DB.Model(&user).Where("phone=?", phone).Select()
	if err == pg.ErrNoRows {
		user.Phone = phone
		err = Insert(&user)
		if err != nil {
			return
		}
	} else {
		return
	}
	return
}
