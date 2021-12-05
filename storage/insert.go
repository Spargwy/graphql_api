package storage

func Insert(data interface{}) error {
	_, err := DB.Model(data).OnConflict("(id) do update").Insert()
	if err != nil {
		return err
	}
	return nil
}
