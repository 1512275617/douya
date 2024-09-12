package model

import (
	"time"
)

type Jujiliebiao struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key" json:"id"`
	DramaTitle string    `gorm:"column:drama_title;type:varchar(255);comment:短剧名称;NOT NULL" json:"drama_title"`
	Jishu      int       `gorm:"column:jishu;type:int(11);comment:短剧级数;NOT NULL" json:"jishu"`
	Url        string    `gorm:"column:url;type:varchar(255);comment:视频链接;NOT NULL" json:"url"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;comment:创建时间;NOT NULL" json:"create_time"`
}

func (m *Jujiliebiao) TableName() string {
	return "jujiliebiao"
}
