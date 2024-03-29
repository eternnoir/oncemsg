package main

import (
	"github.com/codegangsta/martini-contrib/binding"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	ctrl "oncemsg/controls"
	"oncemsg/utility"
	"os"
)

type ViewMessage struct {
	Content string `form:"Content"`
}

func main() {
	m := martini.Classic()
	m.Use(martini.Static("assets"))
	m.Use(render.Renderer(render.Options{
		Layout:     "layout",
		Extensions: []string{".html"},
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", "")
	})

	m.Post("/u", binding.Bind(ViewMessage{}), func(msg ViewMessage, r render.Render, req *http.Request) {
		hosturl := os.Getenv("URL")
		str, err := ctrl.SaveMsg(msg.Content, "text")
		if err != nil {
			r.HTML(404, "error", "")
			ctrl.LogError(err)
			return
		}
		ctrl.LogInfo("NEW MSG:" + str)
		r.HTML(200, "url", hosturl+"/r/"+str)
	})

	m.Get("/r/:unid", func(params martini.Params, r render.Render) {
		unid := params["unid"]
		msg, err := ctrl.GetSecMsg(unid)
		if err != nil {
			ctrl.LogError(err)
			r.HTML(404, "error", "")
			return
		}
		if msg == nil {
			ctrl.LogWarn("NO MSG")
			r.HTML(404, "error", "")
			return
		}
		ctrl.LogInfo("Get MSG: " + unid)
		demsg, _ := utility.DeCryAse(msg.Content)
		r.HTML(200, "read", demsg)
	})

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), m))
}
