package models

import "fmt"

type IModel interface {
	ToString() string
}

type Model struct {
	Id   int64
	Name string
}

func (this *Model) setName(name string) {
	this.Name = name
}
func (this *Model) setId(id int64) {
	this.Id = id
}

func (this *Model) ToString() string {
	return fmt.Sprintf("Model is %s:")
}
