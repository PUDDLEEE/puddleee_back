package types

type Repository interface {
	Create()
	FindAll()
	FindOneById()
	Update()
	Delete()
}
