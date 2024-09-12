package model

type User struct {
	Id       int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"column:username;type:varchar(255);comment:用户名;NOT NULL" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);comment:密码;NOT NULL" json:"password"`
}

func (m *User) TableName() string {
	return "user"
}
