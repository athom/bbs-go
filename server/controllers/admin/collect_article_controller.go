package admin

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/mlogclub/simple"

	"github.com/mlogclub/bbs-go/controllers/render"
	"github.com/mlogclub/bbs-go/services2"
)

type CollectArticleController struct {
	Ctx iris.Context
}

func (this *CollectArticleController) GetBy(id int64) *simple.JsonResult {
	t := services2.CollectArticleService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *CollectArticleController) AnyList() *simple.JsonResult {
	list, paging := services2.CollectArticleService.Query(simple.NewParamQueries(this.Ctx).
		EqAuto("rule_id").
		EqAuto("link_id").
		EqAuto("article_id").
		EqAuto("status").
		PageAuto().Desc("id"))
	var results []map[string]interface{}
	for _, article := range list {
		item := simple.NewRspBuilderExcludes(article, "content").Build()
		item["user"] = render.BuildUserDefaultIfNull(article.UserId)
		results = append(results, item)
	}

	return simple.JsonData(&simple.PageResult{Results: results, Page: paging})
}
