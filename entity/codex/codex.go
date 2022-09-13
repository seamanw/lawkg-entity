/*
法律法规
*/
package codex

import (
	"github.com/seamanw/lawkg-entity/entity/codex/content"
	"github.com/seamanw/lawkg-entity/entity/codex/meta"
)

/*
《中华人民共和国立法法》
http://www.npc.gov.cn/zgrdw/npc/dbdhhy/12_3/2015-03/18/content_1930713.htm
第六十一条 法律根据内容需要，可以分编、章、节、条、款、项、目
*/
type Codex struct {
	Id string `json:"id"`
	/*
		法律法规的元数据
	*/
	*meta.Name
	*meta.Enforce
	*meta.Source
	*meta.Structure
	/*
		法规的数据
	*/
	Preface     *content.Preface   // 法规前言数据
	VolumeData  *[]content.Volume  `json:"volume_data"` // 法规主体数据：编、章或条的切片
	ChapterData *[]content.Chapter `json:"chapter_data"`
	ClauseData  *[]content.Clause  `json:"clause_data"`
	Appendix    *content.Appendix  // 法规附则数据
}

func NewCodex() Codex {
	return Codex{}
}
