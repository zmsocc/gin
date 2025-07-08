package models

type Article struct {
	Id     int
	Title  string
	CateId int
	State  int
}

func (Article) TableName() string {
	return "article"
}
