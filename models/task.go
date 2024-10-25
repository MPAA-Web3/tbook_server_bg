package models

type Task struct {
	ID       uint   `gorm:"primaryKey"`        // 主键
	Name     string `gorm:"size:255;not null"` // 奖品名称
	ImageUrl string `gorm:"size:255;not null"` //URL
	Value    string `gorm:"size:255;not null"` //Value
	Type     string `gorm:"size:255;not null"` //Type
}

// TableName returns the corresponding database table name for this struct.
func (m Task) TableName() string {
	return "task"
}
