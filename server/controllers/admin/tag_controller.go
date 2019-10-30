package admin

import (
	"github.com/kataras/iris"
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/services2"
	"github.com/mlogclub/simple"
	"strconv"
)

type TagController struct {
	Ctx iris.Context
}

func (this *TagController) GetBy(id int64) *simple.JsonResult {
	t := services2.TagService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *TagController) AnyList() *simple.JsonResult {
	list, paging := services2.TagService.Query(simple.NewParamQueries(this.Ctx).
		LikeAuto("name").
		EqAuto("status").
		PageAuto().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}

func (this *TagController) PostCreate() *simple.JsonResult {
	t := &model.Tag{}
	err := this.Ctx.ReadForm(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	if len(t.Name) == 0 {
		return simple.JsonErrorMsg("name is required")
	}
	if services2.TagService.GetByName(t.Name) != nil {
		return simple.JsonErrorMsg("标签「" + t.Name + "」已存在")
	}

	t.Status = model.TagStatusOk
	t.CreateTime = simple.NowTimestamp()
	t.UpdateTime = simple.NowTimestamp()

	err = services2.TagService.Create(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *TagController) PostUpdate() *simple.JsonResult {
	id, err := simple.FormValueInt64(this.Ctx, "id")
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	t := services2.TagService.Get(id)
	if t == nil {
		return simple.JsonErrorMsg("entity not found")
	}

	err = this.Ctx.ReadForm(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	if len(t.Name) == 0 {
		return simple.JsonErrorMsg("name is required")
	}
	if tmp := services2.TagService.GetByName(t.Name); tmp != nil && tmp.Id != id {
		return simple.JsonErrorMsg("标签「" + t.Name + "」已存在")
	}

	t.UpdateTime = simple.NowTimestamp()
	err = services2.TagService.Update(t)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *TagController) AnyListAll() *simple.JsonResult {
	categoryId, err := strconv.ParseInt(this.Ctx.FormValue("categoryId"), 10, 64)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if categoryId < 0 {
		return simple.JsonErrorMsg("请指定categoryId")
	}
	list, err := services2.TagService.ListAll(categoryId)
	if err != nil {
		return simple.JsonData([]interface{}{})
	}
	return simple.JsonData(list)
}

// 标签数据级联选择器
func (this *TagController) GetCascader() *simple.JsonResult {
	categories, err := services2.CategoryService.GetCategories()
	if err != nil {
		return simple.JsonErrorMsg("数据加载失败")
	}

	var results []map[string]interface{}

	for _, cat := range categories {
		tags, err := services2.TagService.ListAll(cat.Id)
		if err != nil || len(tags) == 0 {
			continue
		}

		var tagOptions []map[string]interface{}
		for _, tag := range tags {
			tagOption := make(map[string]interface{})
			tagOption["value"] = tag.Id
			tagOption["label"] = tag.Name
			tagOptions = append(tagOptions, tagOption)
		}

		option := make(map[string]interface{})
		option["value"] = cat.Id
		option["label"] = cat.Name
		option["children"] = tagOptions

		results = append(results, option)
	}

	return simple.JsonData(results)

}
