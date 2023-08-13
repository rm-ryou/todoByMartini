package main

import (

  "example.com/todoByMartini/app/handler"
  "example.com/todoByMartini/app/model"

  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/binding"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/", handler.Hello)

  m.Group("/todos", func(r martini.Router) {
    r.Get("/", handler.GetTasks)
    r.Get("/new", handler.NewTask)
    r.Post("/new", binding.Form(model.Todo{}), handler.GetNewTask)
    // r.Delete("/delete", handler.DeleteTask)
  })

  m.Run()
}
