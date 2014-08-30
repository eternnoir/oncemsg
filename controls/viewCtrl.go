package controls

import "github.com/martini-contrib/render"

func IndexView(r render.Render) {
	r.HTML(200, "index", "")
}
