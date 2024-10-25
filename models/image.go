package models

type Image struct {
	ID   uint   `gorm:"primaryKey"` // 主键
	Key  string `gorm:"key"`
	Name string `gorm:"name"`
	Url  string `gorm:"url"`
}

// TableName returns the corresponding database table name for this struct.
func (m Image) TableName() string {
	return "image"
}
