package layout

import (
	"time"

	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/template/layout"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
)

type Index struct {
	layout.Template
}

// 初始化
func (p *Index) Init(ctx *builder.Context) interface{} {

	// layout 的左上角 的 title
	p.Title = "QuarkLite"

	// layout 的左上角 的 logo
	p.Logo = false

	// layout 的头部行为
	p.Actions = nil

	// layout 的菜单模式,side：右侧导航，top：顶部导航，mix：混合模式
	p.Layout = "mix"

	// layout 的菜单模式为mix时，是否自动分割菜单
	p.SplitMenus = false

	// layout 的内容模式,Fluid：定宽 1200px，Fixed：自适应
	p.ContentWidth = "Fluid"

	// 主题色,"#1890ff"
	p.PrimaryColor = "#1890ff"

	// 是否固定 header 到顶部
	p.FixedHeader = true

	// 是否固定导航
	p.FixSiderbar = true

	// 使用 IconFont 的图标配置
	p.IconfontUrl = "//at.alicdn.com/t/font_1615691_3pgkh5uyob.js"

	// 当前 layout 的语言设置，'zh-CN' | 'zh-TW' | 'en-US'
	p.Locale = "zh-CN"

	// 侧边菜单宽度
	p.SiderWidth = 208

	// 网站版权 time.Now().Format("2006") + " QuarkGo"
	p.Copyright = time.Now().Format("2006") + " QuarkGo"

	// 友情链接
	p.Links = []map[string]interface{}{
		{
			"key":   "1",
			"title": "Quark",
			"href":  "http://quarkcloud.io/",
		},
		{
			"key":   "2",
			"title": "爱小圈",
			"href":  "http://www.ixiaoquan.com",
		},
		{
			"key":   "3",
			"title": "Github",
			"href":  "https://github.com/quarkcloudio",
		},
	}

	return p
}
