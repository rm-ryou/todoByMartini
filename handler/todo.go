package handler

import (
  "fmt"

  "example.com/todoByMartini/model"

  "github.com/martini-contrib/render"
)

var todos []model.Todo

func GetTodoRequest(todo model.Todo, r render.Render) {
  fmt.Println(todo.Name)
  addTodo(todo)
  r.HTML(200, "todos", todos)
}

func ShowTodos(r render.Render) {
  r.HTML(200, "todos", nil)
}

func addTodo(newTodo model.Todo) {
  todos = append(todos, newTodo)
}
