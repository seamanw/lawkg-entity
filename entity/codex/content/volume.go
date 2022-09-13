package content

import "github.com/seamanw/lawkg-entity/entity/codex/meta"

/*
Volume: 编，法律文件中的一部分
*/
type Volume struct {
	*meta.Order
	*meta.Name
	*meta.Enforce
	// 是否使用了章这种目录结构
	IsUseChapter bool       `json:"is_use_chapter"`
	Data         *[]Chapter `json:"data"`
}
