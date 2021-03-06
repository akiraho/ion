package main

import (
	"github.com/get-ion/ion"
	"github.com/get-ion/ion/context"
)

type mypage struct {
	Title   string
	Message string
}

func main() {
	app := ion.New()

	app.RegisterView(ion.HTML("./templates", ".html").Layout("layout.html"))
	// TIP: append .Reload(true) to reload the templates on each request.

	app.Get("/", func(ctx context.Context) {
		ctx.Gzip(true)
		ctx.ViewData("", mypage{"My Page title", "Hello world!"})
		ctx.View("mypage.html")
		// Note that: you can pass "layout" : "otherLayout.html" to bypass the config's Layout property
		// or view.NoLayout to disable layout on this render action.
		// third is an optional parameter
	})

	// http://localhost:8080
	app.Run(ion.Addr(":8080"))
}
