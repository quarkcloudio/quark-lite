package resource

import (
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
	"github.com/quarkcms/quark-lite/model"
)

type Demo struct {
	adminresource.Template
}

// 初始化
func (p *Demo) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "Demo"

	// 模型
	p.Model = &model.Demo{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *Demo) Fields(ctx *builder.Context) []interface{} {
	field := &adminresource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("name", "名称"),

		field.Datetime("created_at", "创建时间"),
	}
}

// 搜索
func (p *Demo) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("name", "名称"),
	}
}

// 行为
func (p *Demo) Actions(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&actions.CreateLink{}).Init(p.Title),
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Disable{}).Init("批量禁用"),
		(&actions.Enable{}).Init("批量启用"),
		(&actions.EditLink{}).Init("编辑"),
		(&actions.Delete{}).Init("删除"),
		(&actions.FormSubmit{}).Init(),
		(&actions.FormReset{}).Init(),
		(&actions.FormBack{}).Init(),
		(&actions.FormExtraBack{}).Init(),
	}
}
