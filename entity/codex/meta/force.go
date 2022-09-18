/*
法律法规效力状态的元表示
*/
package meta

import "time"

/*
《英汉法律词典（第五版）》：https://www.dedao.cn/ebook/detail?id=YxRj1dbLGNgA2oaD6VBevmQZ7rnbYWmBQa3kyOKRMzJpX9lP5dxqE41j8De2pmOP
validity of law 法律效力
extraterritorial effect 治外法权效力；域外的效力
inure（=enure） v. 1. 生效，有效力，适用；有用 2. 使坚强，使习惯于（不利条件）
inurement n. 生效，适用
dubious validity 未定效力，可疑的效力，含糊的效力，半信半疑的效力
material sphere of validity 属事效力范围，实质的效力范围
effectiveness n. 管辖效力
force of law 法的效力（又译：法律效力）
*/

/*
时效性：当前是否有效，生效、失效时间
*/
type ForceOfTime struct {
	IsEnforceNow bool      `json:"is_enforce_now"`
	EnforceDate  time.Time `json:"enforce_date"`
	ExpiresDate  time.Time `json:"expires_date"`
}

/*
地效性：生效地区
*/
