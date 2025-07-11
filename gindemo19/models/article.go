package models

type Article struct {
	Id            int
	Title         string
	ArticleCateId int // 外键
	State         int
	ArticleCate   ArticleCate
}

type ArticleV2 struct {
	Id          int
	Title       string
	CateId      int // 外键
	State       int
	ArticleC ArticleCate `gorm:"foreignKey:CateId"`
}

func (Article) TableName() string {
	return "article"
}

func (ArticleV2) TableName() string {
	return "article"
}
