package model

import (
	"time"
)

type DramaList struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	DramaTitle string    `gorm:"column:drama_title;type:varchar(255);comment:短剧名;NOT NULL" json:"drama_title"`
	Picture    string    `gorm:"column:picture;type:varchar(255);comment:图片" json:"picture"`
	Status     int       `gorm:"column:status;type:tinyint(4);comment:状态;NOT NULL" json:"status"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:创建时间;NOT NULL" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;comment:更新时间;NOT NULL" json:"update_time"`
	Text       string    `gorm:"column:text;type:text;comment:简介" json:"text"`
}

func (m *DramaList) TableName() string {
	return "drama_list"
}
