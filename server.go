package main

import (

  "example.com/todoByMartini/handler"
  _ "example.com/todoByMartini/model"

  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/", handler.Hello)

  m.Run()
}
