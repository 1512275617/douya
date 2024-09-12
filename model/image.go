package model

type Image struct {
	Id      int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Title   string `gorm:"column:title;type:varchar(255);comment:名称;NOT NULL" json:"title"`
	Picture string `gorm:"column:picture;type:varchar(255);comment:轮播图;NOT NULL" json:"picture"`
	Url     string `gorm:"column:url;type:varchar(255);comment:跳转链接;NOT NULL" json:"url"`
	Sort    int    `gorm:"column:sort;type:int(11);comment:排序;NOT NULL" json:"sort"`
	Status  int    `gorm:"column:status;type:tinyint(4);comment:状态;NOT NULL" json:"status"`
}

func (m *Image) TableName() string {
	return "image"
}
