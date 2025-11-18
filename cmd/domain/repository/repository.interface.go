package repository

type Repository interface {
	Get(id string) (interface{}, error)
	Create(entity interface{}) error
	Update(entity interface{}) error
	Delete(id int) error
}
