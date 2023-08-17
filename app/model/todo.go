package model

type Todo struct {
  ID    int     `form:"id"`
  Name  string  `form:"name"`
}
