package storage

func Insert(data interface{}) error {
	_, err := DB.Model(data).Insert()
	if err != nil {
		return err
	}
	return nil
}
