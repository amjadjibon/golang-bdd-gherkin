package models

type BookBase struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (b *BookBase) Validate() bool {
	if b.Title == "" && b.Author == "" {
		return false
	}
	return true
}

type Book struct {
	ID     int    `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
}
