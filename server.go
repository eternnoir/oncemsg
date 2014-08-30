package main

import (
	msgctrl "oncemsg/controls"

	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type Message struct {
	Content string `form:"Content"`
}

func main() {
	m := martini.Classic()

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", "")
	})
	m.Post("/save", binding.Bind(Message{}), func(msg Message, r render.Render) string {
		id, err := msgctrl.SaveMsg("aa")
		if err == nil {
			return msg.Content
		} else {
			return id
		}
	})
	m.Run()
}
