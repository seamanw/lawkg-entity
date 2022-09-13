/*
法律法规的目录结构构造
*/
package meta

// 元数据：
type Structure struct {
	IsUsePreface   bool `json:"is_use_preface"`
	IsUseAppendix  bool `json:"is_use_appendix"`
	IsUseVolume    bool `json:"is_use_volume"`
	IsUseChapter   bool `json:"is_use_chapter"`
	IsUseProvision bool `json:"is_use_provision`
	IsUseClause    bool `json:"is_use_clause"`
}
