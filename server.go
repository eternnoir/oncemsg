package main

import (
	"log"
	"net/http"
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
	/*
		err := http.ListenAndServe(":"+os.Getenv("PORT"), m)
		if err != nil {
			panic(err)
		}*/
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", "")
	})
	m.Post("/u", binding.Bind(ViewMessage{}), func(msg ViewMessage, r render.Render) string {
		return msg.Content
	})
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), m))
}
