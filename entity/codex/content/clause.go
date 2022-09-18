package content

import "github.com/seamanw/lawkg-entity/entity/codex/meta"

/*
Clause: 条，法律文件的基本组成单元
*/
type Clause struct {
	*meta.Order
	*meta.Name
	*meta.ForceOfTime
	// 条内是否分不同的款
	IsUseParagraph bool         `json:"is_use_paragraph"`
	Data           *[]Paragraph `json:"data"`
}
