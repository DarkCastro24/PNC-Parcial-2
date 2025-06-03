package repositories

import "PNC-Parcial-2/entities"

type UserRepository interface {
	GetAll() ([]entities.Book, error)
	GetById(id string) (entities.Book, error)
	Create(user entities.Book) error
	Update(id, title, author string) error
	Delete(id string) error
}
