/*
数据库模板
*/
package data

import "time"

/*
数据库模板字段
*/
type TemplateFields struct {
	ID             string    `json:"id"`
	CreateBy       string    `json:"create_by"`
	CreateTime     time.Time `json:"create_time"`
	LastUpdateBy   string    `json:"last_update_by"`
	LastUpdateTime time.Time `json:"last_update_time"`
	Status         int       `json:"status"`
	IsDelete       int       `json:"is_delete"`
}

/*
新建模板字段结构体
*/
func NewTemplateFields() *TemplateFields {
	return &TemplateFields{
		ID:             "",
		CreateBy:       "",
		CreateTime:     time.Now(),
		LastUpdateBy:   "",
		LastUpdateTime: time.Now(),
		Status:         0,
		IsDelete:       0,
	}
}
