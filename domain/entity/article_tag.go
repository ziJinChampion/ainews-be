package entity

type ArticleTag struct {
	Id        int `gorm:"primaryKey;autoIncrement"`
	TagId     int
	ArticleId int
}
