package repository

import "models"

type InMemoryRepo struct {
	carts map[string]models.Cart
}

func NewInMemoryRepo() *InMemoryRepo {
	m := make(map[string]models.Cart)
	i := &InMemoryRepo{carts: m}
	return i
}

func (i *InMemoryRepo) Add(cart models.Cart) {
	if len(cart.Id) == 0 {
		panic("Invalid id passed to Add cart")
	}
	i.carts[cart.Id] = cart
}

func (i *InMemoryRepo) Get(id string) (models.Cart, bool) {
	cart, ok := i.carts[id]
	return cart, ok
}

func (i *InMemoryRepo) Delete(id string) bool {
	_, ok := i.carts[id]
	delete(i.carts, id)
	return ok
}
