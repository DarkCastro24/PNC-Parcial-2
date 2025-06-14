package services

import (
	"PNC-Parcial-2/entities"
	dto "PNC-Parcial-2/entities/dto"
	"PNC-Parcial-2/repositories"
	"errors"
	"github.com/google/uuid"
	"regexp"
)

type BookService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *BookService {
	return &BookService{repo: repo}
}

func validateISBN(isbn string) bool {
	regex := `^(978|979)[0-9]{10}$`
	match, _ := regexp.MatchString(regex, isbn)
	return match
}

func (s *BookService) CreateBook(input dto.BookDTO) (entities.Book, error) {
	if !validateISBN(input.ISBN) {
		return entities.Book{}, errors.New("El formato del ISBN no es valido")
	} else {
		if input.Pages < 10 {
			return entities.Book{}, errors.New("El numero de paginas debe ser minimo 10")
		}
		user := entities.NewBook(
			uuid.New().String(),
			input.Title,
			input.Author,
			input.ISBN,
			input.PublicationYear,
			input.Language,
			input.Pages,
			input.Genre,
		)
		return user, s.repo.Create(user)
	}
}

func (s *BookService) GetAllUsers() ([]entities.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetBook(id string) (entities.Book, error) {
	return s.repo.GetById(id)
}

func (s *BookService) UpdateUser(id string, input dto.BookDTO) error {
	return s.repo.Update(id, input.Title, input.Language)
}

func (s *BookService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
