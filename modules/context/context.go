package context

import (
	"fmt"
	"time"

	"github.com/Unknwon/log"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"

	"blog/modules/setting"
)

// Context 自定义的 Context
type Context struct {
	*macaron.Context
	Session session.Store
	Flash   *session.Flash
}

// Contexter 初始化一个自定义Context
// 这里设置了一些公共信息 扩展一些ctx的方法
func Contexter() macaron.Handler {
	return func(c *macaron.Context, sess session.Store, f *session.Flash) {
		ctx := &Context{
			Context: c,
			Session: sess,
			Flash:   f,
		}
		ctx.Data["PageStartTime"] = time.Now()

		log.Debug("Session ID: %s", sess.ID())
		ctx.Data["Version"] = "0.1"
		ctx.Data["VersionDate"] = "2016-09-05"
		c.Map(ctx)
	}
}

// Handle 处理错误
func (ctx *Context) Handle(status int, title string, err error) {
	if err != nil {
		log.Error("%s: %v", title, err)
		if macaron.Env != macaron.PROD {
			ctx.Data["ErrorMsg"] = err
		}
	}

	switch status {
	case 404:
		ctx.Data["Title"] = "Page Not Found"
		ctx.Data["ErrMsg"] = "Page Not Found."
	case 500:
		ctx.Data["Title"] = "Internal Server Error"
		ctx.Data["ErrMsg"] = "Internal Server Error."
	}
	// ctx.HasTemplate("")
	ctx.HTMLSet(status, "admin", "error")
}

func (ctx *Context) HasTemplate(t ...string) bool {
	set := ""
	name := ""
	switch len(t) {
	case 1:
		set = "/"
		name = t[0]
	default:
		set = "/" + t[0] + "/"
		name = t[1]
	}
	tPath := setting.AppPath + "templates" + set + name
	fmt.Println(tPath)
	return true

}
