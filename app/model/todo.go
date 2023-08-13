package model

type Todo struct {
  ID    uint
  Name  string  `form:"name"`
}
