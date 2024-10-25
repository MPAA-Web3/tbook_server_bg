package models

type Tasks struct {
	ID    uint   `gorm:"primaryKey"` // 主键
	Name  string `gorm:"size:255;not null"`
	Url   string `gorm:"size:255;not null"`
	Type  string `gorm:"size:255;not null"`
	State string `gorm:"size:255;not null"`
}

// TableName returns the corresponding database table name for this struct.
func (m Tasks) TableName() string {
	return "tasks"
}
