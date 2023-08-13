package model

type Todo struct {
  ID    int
  Name  string  `form:"name"`
}
