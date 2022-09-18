package content

import "github.com/seamanw/lawkg-entity/entity/codex/meta"

/*
章，法律文件中的一部分
*/
type Chapter struct {
	*meta.Order
	*meta.Name
	*meta.ForceOfTime
	// 是否使用了节这种目录结构
	IsUseProvision bool         `json:"is_use_provision"`
	Data           *[]Provision `json:"data"`
}
