/*
法律法规的目录结构构造
*/
package meta

// 元数据：
type Structure struct {
	IsUsePreface   bool
	IsUseAppendix  bool
	IsUseVolume    bool
	IsUseChapter   bool
	IsUseProvision bool
	IsUseClause    bool
}
