package admin

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/mlogclub/simple"

	"github.com/mlogclub/bbs-go/services2"
)

type SubjectContentController struct {
	Ctx iris.Context
}

func (this *SubjectContentController) GetBy(id int64) *simple.JsonResult {
	t := services2.SubjectContentService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *SubjectContentController) AnyList() *simple.JsonResult {
	list, paging := services2.SubjectContentService.Query(simple.NewParamQueries(this.Ctx).EqAuto("subject_id").
		EqAuto("entity_type").EqAuto("entity_id").EqAuto("deleted").PageAuto().Desc("id"))

	var itemList []map[string]interface{}
	for _, v := range list {
		subject := services2.SubjectService.Get(v.SubjectId)
		itemList = append(itemList, simple.NewRspBuilder(v).Put("subject", subject).Build())
	}
	return simple.JsonData(&simple.PageResult{Results: itemList, Page: paging})
}
