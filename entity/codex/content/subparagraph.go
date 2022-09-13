package content

import "github.com/seamanw/lawkg-entity/entity/codex/meta"

/*
SubParagraph: 项，款的基本组成单元
*/
type SubParagraph struct {
	*meta.Order
	*meta.Name
	*meta.Enforce
	// 项下是否有不同的目
	IsUseItem bool    `json:"is_use_item"`
	Data      *[]Item `json:"data"`
}
