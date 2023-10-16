package model

type Tag struct {
	Model

	Name        string `gorm:"column:tag_name" json:"tag_name"`
	Description string `gorm:"column:tag_description" json:"tag_description"`
}

func GetAllTags() (tags []Tag, err error) {
	err = client.Table("tags").Find(&tags).Error
	return
}

func CreateNewTag(name, description string) (bool, error) {
	err := client.Table("tags").Create(&Tag{
		Name:        name,
		Description: description,
	}).Error

	return true, err
}
