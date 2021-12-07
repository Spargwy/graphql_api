package storage

import (
	"gql_app/graph/model"

	"github.com/go-pg/pg/v10"
)

func (db *Psql) Insert(data interface{}) error {
	_, err := db.DB.Model(data).OnConflict("(id) do update").Insert()
	if err != nil {
		return err
	}
	return nil
}

func (db *Psql) SelectUserByPhone(phone string) (user model.User, err error) {
	err = db.DB.Model(&user).Where("phone=?", phone).Select()
	if err == pg.ErrNoRows {
		user.Phone = phone
		err = db.Insert(&user)
		if err != nil {
			return
		}
	} else {
		return
	}
	return
}

func (db *Psql) SelectUserByID(userID int) (user model.User, err error) {
	user.ID = userID
	err = db.DB.Model(&user).WherePK().Select()
	if err != nil {
		return
	}
	return
}
