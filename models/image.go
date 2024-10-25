package models

type Image struct {
	Key string `gorm:"key"`
	Url string `gorm:"url"`
}

// TableName returns the corresponding database table name for this struct.
func (m Image) TableName() string {
	return "image"
}
