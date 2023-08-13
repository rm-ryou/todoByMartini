package handler

import (
  "github.com/martini-contrib/render"
)

func Hello(r render.Render) {
  r.HTML(200, "hello", nil)
}
