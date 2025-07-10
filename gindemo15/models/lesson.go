package models

type Lesson struct {
	Id        int
	ClassName string
	Student []Student `gorm:"many2many:lesson_student"`
}

func (Lesson) TableName() string {
	return "lesson"
}
