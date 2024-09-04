package model

type User struct {
	Id   int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name string `gorm:"column:name;type:varchar(255);NOT NULL" json:"name"`
}

func (m *User) TableName() string {
	return "user"
}
