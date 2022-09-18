/*
法律法规抽象
*/
package entity

import (
	"time"

	"github.com/seamanw/lawkg-entity/entity/codex/content"
	"github.com/seamanw/lawkg-entity/entity/codex/meta"
)

/*
法律法规的元数据
*/
type CodexMeta struct {
	*meta.Name
	*meta.ForceOfTime
	*meta.Source
}

func NewCodexMeta() *CodexMeta {
	return &CodexMeta{
		Name: &meta.Name{
			FullName:  "",
			BriefName: "",
		},
		ForceOfTime: &meta.ForceOfTime{
			IsEnforceNow: true,
			EnforceDate:  time.Date(1949, time.Month(10), 1, 0, 0, 0, 0, time.UTC),
			ExpiresDate:  time.Date(2999, time.Month(12), 23, 59, 59, 999, 999, time.UTC),
		},
		Source: &meta.Source{
			URI: []string{""},
		},
	}
}

/*
法律法规的内容数据
《中华人民共和国立法法》：http://www.npc.gov.cn/zgrdw/npc/dbdhhy/12_3/2015-03/18/content_1930713.htm
第六十一条 法律根据内容需要，可以分编、章、节、条、款、项、目
*/
type CodexData struct {
	*meta.Structure                    // 法规内容结构
	Preface         *content.Preface   // 法规的前言
	VolumeData      *[]content.Volume  // 法规的编
	ChapterData     *[]content.Chapter // 法规的章
	ClauseData      *[]content.Clause  // 法规的条
	Appendix        *content.Appendix  // 法规的附则
}

func NewCodex() *CodexData {
	return &CodexData{}
}
