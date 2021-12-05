package resolvers

func (r *Resolver) Insert(data interface{}) error {
	_, err := r.DB.Model(data).OnConflict("(id) do update").Insert()
	if err != nil {
		return err
	}
	return nil
}
