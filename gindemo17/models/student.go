package models

type Student struct {
	Id      int
	Number  int
	Name    string
	LessonId int
	Lesson	[]Lesson `gorm:"many2many:lesson_student"`
}

func (Student) TableName() string {
	return "student"
}
