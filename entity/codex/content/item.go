package content

import "github.com/seamanw/lawkg-entity/entity/codex/meta"

/*
Item: 目，项的基本组成单元
*/
type Item struct {
	*meta.Order
	*meta.Name
	*meta.Enforce
	Data string
}
