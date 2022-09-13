package content

import "github.com/seamanw/lawkg-entity/entity/codex/meta"

/*
节，法律文件中的一部分
*/
type Provision struct {
	*meta.Order
	*meta.Name
	*meta.Enforce
	// 是否使用了条这种目录结构
	IsUseClause bool      `json:"is_use_clause"`
	Data        *[]Clause `json:"data"`
}
