package models

type Nav struct {
	Id     int
	Title  string
	Url    string
	Status int
	Sort   int
}

func (Nav) TableName() string {
	return "nav"
}
