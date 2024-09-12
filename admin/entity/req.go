package entity

type DramaListReq struct {
	DramaTitle string `json:"drama_title" form:"drama_title"`
	CreateTime string `json:"create_time" form:"create_time"`
	*PageReq
}

// PageReq 分页结构体
type PageReq struct {
	Size int    `json:"size" form:"size" uri:"size"`
	Page int    `json:"page" form:"page" uri:"page"`
	Sort string `json:"sort" form:"sort" uri:"sort"`
}
type LoginPassWordReq struct {
	Username string `json:"username" form:"username" uri:"username" binding:"required"` //登录用户名
	Password string `json:"password" form:"password" uri:"password" binding:"required"` //登录密码
}
