package main

import (
	ctrl "oncemsg/controls"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Message struct {
	Content string `form:"Content"`
}

func main() {
	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Get("/", ctrl.IndexView)
	m.Post("/save", binding.Bind(Message{}), func(msg Message, r render.Render) string {
		return msg.Content
	})
	m.Run()
}
