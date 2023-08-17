package handler

import (
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
  addTodo(todo)
  r.Redirect("/todos/")
}

func DeleteTask(todo model.Todo, r render.Render) {
  deleteTodo(todo)
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

func deleteTodo(todo model.Todo)  {
  println(todo.Name)
}
