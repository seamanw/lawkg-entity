package content

import "github.com/seamanw/lawkg-entity/entity/codex/meta"

/*
Paragraph: 款，条的基本组成单元
*/
type Paragraph struct {
	*meta.Order
	*meta.Name
	*meta.Enforce
	// 款下是否有不同的项
	IsUseSubParagraph bool            `json:"is_use_subparagraph"`
	Data              *[]SubParagraph `json:"data"`
}
