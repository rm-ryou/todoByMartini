package main

import (

  "example.com/todoByMartini/handler"
  "example.com/todoByMartini/model"

  "github.com/go-martini/martini"
  "github.com/martini-contrib/render"
  "github.com/martini-contrib/binding"
)

func main() {
  m := martini.Classic()
  m.Use(render.Renderer())

  m.Get("/", handler.Hello)

  m.Get("/todos", handler.ShowTodos)
  m.Post("/todos", binding.Form(model.Todo{}), handler.GetTodoRequest)

  m.Run()
}
