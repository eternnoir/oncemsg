package main

import (
	"net/http"
	ctrl "oncemsg/controls"
	"os"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type ViewMessage struct {
	Content string `form:"Content"`
}

func main() {
	m := martini.Classic()
	http.ListenAndServe(":"+os.Getenv("PORT"), m)

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Get("/", ctrl.IndexView)
	m.Post("/u", binding.Bind(ViewMessage{}), func(msg ViewMessage, r render.Render) string {
		return msg.Content
	})
	m.Run()
}
