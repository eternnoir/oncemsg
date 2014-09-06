package main

import (
	"log"
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
	m.Use(martini.Static("assets"))
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", "")
	})
	m.Post("/u", binding.Bind(ViewMessage{}), func(msg ViewMessage, r render.Render, req *http.Request) {
		hosturl := os.Getenv("URL")
		str, err := ctrl.SaveMsg(msg.Content, "text")
		if err != nil {
			r.HTML(404, "error", "")
			return
		}
		r.HTML(200, "url", hosturl+"/r/"+str)
	})

	m.Get("/r/:unid", func(params martini.Params, r render.Render) {
		unid := params["unid"]
		msg, err := ctrl.GetSecMsg(unid)
		if err != nil {
			r.HTML(404, "error", "")
			return
		}
		if msg == nil {
			r.HTML(404, "error", "")
			return
		}
		r.HTML(200, "read", msg.Content)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), m))
}
