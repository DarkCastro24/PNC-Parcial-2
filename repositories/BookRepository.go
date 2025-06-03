package repositories

import (
	"PNC-Parcial-2/entities"
	"gorm.io/gorm"
)

type bookRepositoryGORM struct {
	db *gorm.DB
}

func NewBookRepositoryGORM(db *gorm.DB) UserRepository {
	return &bookRepositoryGORM{db: db}
}

func (r *bookRepositoryGORM) Create(book entities.Book) error {
	return r.db.Create(&book).Error
}

func (r *bookRepositoryGORM) GetAll() ([]entities.Book, error) {
	var books []entities.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepositoryGORM) GetById(id string) (entities.Book, error) {
	var book entities.Book
	err := r.db.First(&book, "id = ?", id).Error
	return book, err
}

func (r *bookRepositoryGORM) Update(id, title, language string) error {
	return r.db.Model(&entities.Book{}).Where("id = ?", id).
		Updates(map[string]interface{}{
			"Title":    title,
			"Language": language,
		}).Error
}

func (r *bookRepositoryGORM) Delete(id string) error {
	return r.db.Delete(&entities.Book{}, "id = ?", id).Error
}
