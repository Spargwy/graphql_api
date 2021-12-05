package resolvers

import (
	"gql_app/graph/model"

	"github.com/go-pg/pg/v10"
)

func (r *Resolver) SelectProducts() (products []*model.Product, err error) {
	err = r.DB.Model(&products).Select()
	if err != nil {
		return
	}
	return
}

func (r *Resolver) SelectUserByID(userID int) (user model.User, err error) {
	user.ID = userID
	err = r.DB.Model(&user).WherePK().Select()
	if err != nil {
		return
	}
	return
}

func (r *Resolver) SelectUserByPhone(phone string) (user model.User, err error) {
	err = r.DB.Model(&user).Where("phone=?", phone).Select()
	if err == pg.ErrNoRows {
		user.Phone = phone
		err = r.Insert(&user)
		if err != nil {
			return
		}
	} else {
		return
	}
	return
}
