package dto

type BookDTO struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	ISBN            string `json:"isbn"`
	PublicationYear string `json:"publicationYear"`
	Language        string `json:"language"`
	Pages           int    `json:"pages"`
	Genre           string `json:"genre"`
}
