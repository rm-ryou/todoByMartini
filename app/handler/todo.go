package handler

import (
  "fmt"

  "example.com/todoByMartini/app/model"

  "github.com/martini-contrib/render"
)

var todos []model.Todo

func GetTasks(r render.Render) {
  r.HTML(200, "todos/index", todos)
}

func NewTask(r render.Render) {
  r.HTML(200, "todos/new", nil)
}

func GetNewTask(todo model.Todo, r render.Render) {
  todo.ID = len(todos)
  fmt.Printf("len todo = %d\n", todo.ID)
  addTodo(todo)
  r.Redirect("/todos/")
}

func DeleteTask(todo model.Todo, r render.Render) {
  fmt.Println("this is delete task func.")
  fmt.Printf("must delete %s\n", todo.Name)
  fmt.Printf("number is %d\n", todo.ID)
  r.HTML(200, "todos/index", todos)
}

func GetTodoRequest(todo model.Todo, r render.Render) {
  r.HTML(200, "todos", todos)
}

func ShowTodos(r render.Render) {
  r.HTML(200, "todos", nil)
}

func addTodo(todo model.Todo) {
  todos = append(todos, todo)
}
