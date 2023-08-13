package handler

import (
  "fmt"

  "example.com/todoByMartini/app/model"

  "github.com/martini-contrib/render"
)

var todos []model.Todo

func GetTasks(r render.Render) {
  fmt.Println("hoge")
  r.HTML(200, "todos/index", todos)
}

func NewTask(r render.Render) {
  r.HTML(200, "todos/new", nil)
}

func GetNewTask(todo model.Todo, r render.Render) {
  fmt.Printf("new task is : %s\n", todo.Name)
  addTodo(todo)
  r.Redirect("/todos/")
}

// func DeleteTask(r render.Render) {
// }

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
