package main

import (
  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/", hello)

  m.Run()
}

func hello(r render.Render) {
  r.HTML(200, "hello", nil)
}
