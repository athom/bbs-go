package api

import (
	"github.com/kataras/iris/context"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"

	"github.com/mlogclub/bbs-go/controllers/render"
	"github.com/mlogclub/bbs-go/model"
	"github.com/mlogclub/bbs-go/services2"
	"github.com/mlogclub/bbs-go/services2/collect/oschina"
	"github.com/mlogclub/bbs-go/services2/collect/studygolang"
)

type ProjectController struct {
	Ctx context.Context
}

func (this *ProjectController) GetCollect1() *simple.JsonResult {
	user := services2.UserTokenService.GetCurrent(this.Ctx)
	if user == nil || user.Id != 1 {
		return simple.JsonErrorMsg("无权限")
	}
	go func() {
		for i := 48; i >= 1; i-- {
			studygolang.GetStudyGoLangPage(i, func(url string) {
				p := studygolang.GetStudyGolangProject(url)
				if p != nil {
					temp := services2.ProjectService.Take("name = ?", p.Name)
					if temp == nil {
						logrus.Info("采集项目：" + p.Name + ", " + url)
						_, _ = services2.ProjectService.Publish(2, p.Name, p.Title, p.Logo, p.Url, p.DocUrl, p.DownloadUrl,
							model.ContentTypeMarkdown, p.Content)
					} else {
						logrus.Warn("项目已经存在：" + temp.Name)
					}
				} else {
					logrus.Warn("项目采集失败：" + url)
				}
			})
		}
	}()
	return simple.JsonSuccess()
}

func (this *ProjectController) GetCollect2() *simple.JsonResult {
	user := services2.UserTokenService.GetCurrent(this.Ctx)
	if user == nil || user.Id != 1 {
		return simple.JsonErrorMsg("无权限")
	}
	go func() {
		for i := 76; i >= 1; i-- {
			urls := oschina.GetPage(i)
			if len(urls) == 0 {
				continue
			}
			for _, url := range urls {
				p := oschina.GetProject(url)
				if p == nil {
					continue
				}
				temp := services2.ProjectService.Take("name = ?", p.Name)
				if temp != nil {
					logrus.Warn("项目已经存在：" + temp.Name)
					continue
				}
				logrus.Info("采集项目：" + p.Name + ", " + url)
				_, _ = services2.ProjectService.Publish(2, p.Name, p.Title, p.Logo, p.Url, p.DocUrl, p.DownloadUrl,
					model.ContentTypeHtml, p.Content)
			}
		}
	}()
	return simple.JsonSuccess()
}

func (this *ProjectController) GetBy(projectId int64) *simple.JsonResult {
	project := services2.ProjectService.Get(projectId)
	if project == nil {
		return simple.JsonErrorMsg("项目不存在")
	}
	return simple.JsonData(render.BuildProject(project))
}

func (this *ProjectController) GetProjects() *simple.JsonResult {
	page := simple.FormValueIntDefault(this.Ctx, "page", 1)

	projects, paging := services2.ProjectService.Query(simple.NewParamQueries(this.Ctx).
		Page(page, 20).Desc("id"))

	return simple.JsonPageData(render.BuildSimpleProjects(projects), paging)
}
