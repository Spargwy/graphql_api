package storage

import "gql_app/graph/model"

func SelectUser(userID int) (user model.User, err error) {
	user.ID = userID
	err = DB.Model(user).WherePK().Select()
	if err != nil {
		return
	}
	return
}
