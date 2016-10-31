package admin

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/session"
	"gopkg.in/macaron.v1"
)

func ClearCache(ctx *macaron.Context, cache cache.Cache) {
	cache.Flush()
	ctx.Redirect("/admin")
}

func ClearSessionCache(ctx *macaron.Context, sess session.Store) {
	user := sess.Get("USER")
	sess.Flush()
	sess.Set("USER", user)
	ctx.Redirect("/admin")

}
