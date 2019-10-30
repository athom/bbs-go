package admin

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/mlogclub/simple"

	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/services2"
	"github.com/mlogclub/bbs-go/services2/collect"
)

type CollectRuleController struct {
	Ctx iris.Context
}

func (this *CollectRuleController) GetBy(id int64) *simple.JsonResult {
	t := services2.CollectRuleService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *CollectRuleController) AnyList() *simple.JsonResult {
	list, paging := services2.CollectRuleService.Query(simple.NewParamQueries(this.Ctx).LikeAuto("title").EqAuto("status").PageAuto().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}

func (this *CollectRuleController) PostCreate() *simple.JsonResult {
	t := &model.CollectRule{}
	err := this.Ctx.ReadForm(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	if len(t.Rule) == 0 {
		return simple.JsonErrorMsg("请输入采集规则")
	}

	t.CreateTime = simple.NowTimestamp()
	t.UpdateTime = simple.NowTimestamp()

	err = services2.CollectRuleService.Create(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *CollectRuleController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt64(this.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	t := services2.CollectRuleService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("entity not found")
	}

	err = this.Ctx.ReadForm(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	if len(t.Rule) == 0 {
		return simple.JsonErrorMsg("请输入采集规则")
	}

	t.UpdateTime = simple.NowTimestamp()

	err = services2.CollectRuleService.Update(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *CollectRuleController) GetRun() *simple.JsonResult {
	id, err := simple.FormValueInt64(this.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	maxDepth := simple.FormValueIntDefault(this.Ctx, "maxDepth", 0)
	go func() {
		collect.Start(id, maxDepth)
	}()
	return simple.JsonSuccess()
}
