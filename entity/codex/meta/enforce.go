/*
效力状态
*/
package meta

import "time"

// 元数据：效力
type Enforce struct {
	IsEnforce   bool      `json:"is_enforce"`
	EnforceDate time.Time `json:"enforce_date"`
	ExpiresDate time.Time `json:"expires_date"`
}
