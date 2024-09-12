package dao

import (
	"context"
	"douya/admin/entity"
	"douya/components"
	"douya/model"
)

func GetDramaList(req *entity.DramaListReq) ([]*entity.DramaListResp, error) {
	var list []*entity.DramaListResp
	var cn int64
	db := components.GetMysqlClient(context.Background())
	tx := db.Model(&model.DramaList{})
	if req.DramaTitle != "" {
		tx.Where("drama_title =?", req.DramaTitle)
	}
	if req.CreateTime != "" {
		tx.Where(" create_time=?", req.CreateTime)
	}
	if req.Page <= 0 {
		req.Page = 1
	}
	if req.Size <= 0 {
		req.Size = 10
	}
	if req.Sort != "desc" && req.Sort != "asc" {
		req.Sort = "asc"
	}
	err := tx.Count(&cn).
		Offset((req.Page - 1) * req.Size).Limit(req.Size).Order("id " + req.Sort).
		Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}
