package entities

type Book struct {
	Id              string `gorm:"type:uuid;primaryKey" json:"id,omitempty"`
	Title           string `gorm:"size:100" json:"title,omitempty"`
	Author          string `gorm:"size:100" json:"author,omitempty"`
	ISBN            string `gorm:"size:100;unique" json:"isbn,omitempty"`
	PublicationYear string `gorm:"size:100" json:"year,omitempty"`
	Language        string `gorm:"size:100" json:"language,omitempty"`
	Pages           int    `gorm:"default:0" json:"pages,omitempty"`
	Genre           string `gorm:"size:100" json:"genre,omitempty"`
}

func NewBook(id string, title string, author string, isbn string, publicationYear string, language string, pages int, genre string) Book {
	return Book{
		Id:              id,
		Title:           title,
		Author:          author,
		ISBN:            isbn,
		PublicationYear: publicationYear,
		Language:        language,
		Pages:           pages,
		Genre:           genre,
	}
}
