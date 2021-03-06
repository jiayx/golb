package admin

import (
	"blog/models"
	"blog/modules/context"
)

// Index 后台管理首页
func Index(ctx *context.Context) {

	ctx.Data["Title"] = "网站概要"

	postCount, err := models.PostsCount(true)
	if err != nil {
		ctx.Handle(500, "Internal Server Error", err)
		return
	}
	ctx.Data["PostCount"] = postCount
	catesCount, err := models.CatesCount()
	if err != nil {
		ctx.Handle(500, "Internal Server Error", err)
		return
	}
	ctx.Data["CountCate"] = catesCount

	TagsCount, err := models.TagsCount()
	ctx.Data["CountTag"] = TagsCount

	ctx.Data["CountComment"] = 132
	// postCount, err := models.CountComment()

	latestPosts, err := models.LatestPosts(6)
	// fmt.Println(latestPosts)
	if err != nil {
		ctx.Handle(500, "Internal Server Error", err)
		return
	}
	ctx.Data["LatestPosts"] = *latestPosts

	ctx.Data["Scripts"] = []string{"admin/js/index.js"}
	ctx.Data["MainActive"] = "active"
	ctx.Data["ConsoleActive"] = "active toggle"

	ctx.HTMLSet(200, "admin", "index")
}

// NotFound 404
func NotFound(ctx *context.Context) {
	ctx.Handle(404, "Page Not Found", nil)
	return
}
