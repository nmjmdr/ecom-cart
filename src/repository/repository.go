package repository

import "models"

type Repo interface {
	Add(cart models.Cart)
	Get(id string) (models.Cart, bool)
	Delete(id string) bool
}

var inMemoryRepo Repo

func Get() Repo {
	if inMemoryRepo == nil {
		inMemoryRepo = NewInMemoryRepo()
	}
	return inMemoryRepo
}
