package main

import "fmt"

type EntityType string

const (
	Beast EntityType = "Beast"
	Item  EntityType = "Item"
)

type entity struct {
	name       string
	entityType EntityType
}

func (e entity) Name() string {
	return e.name
}

func (e entity) IsBeast() bool {
	return e.entityType == Beast
}

func (e entity) IsItem() bool {
	return e.entityType == Item
}

type beast struct {
	*entity
}

func NewBeast(name string) *beast {
	ent := entity{
		name:       name,
		entityType: Beast,
	}
	return &beast{&ent}
}

type item struct {
	*entity
}

func NewItem(name string) *item {
	ent := entity{
		name:       name,
		entityType: Item,
	}
	return &item{&ent}
}

type Entity interface {
	Identify() EntityType
}

func (b beast) Identify() EntityType {
	return b.entity.entityType
}

func (b *beast) ChangeName(nn string) {
	b.name = nn
}

func (i item) Identify() EntityType {
	return i.entity.entityType
}

func NameBeasts(items []Entity) []*beast {
	var list []*beast
	for _, i := range items {
		if i.Identify() == Beast {
			list = append(list, i.(*beast))
		}
	}
	return list
}

func NameItems(items []Entity) []*item {
	var list []*item
	for _, i := range items {
		if i.Identify() == Item {
			list = append(list, i.(*item))
		}
	}
	return list
}

func PrintAnimals(e []Entity) {
	for _, b := range NameBeasts(e) {
		fmt.Println(b.name)
	}
}

func main() {
	animal1 := NewBeast("Lion")
	animal2 := NewBeast("Zebra")
	weapon := NewItem("Sword")
	Entities := []Entity{animal1, weapon, animal2}
	PrintAnimals(Entities)
	for _, b := range NameBeasts(Entities) {
		if b.name == "Lion" {
			b.ChangeName("Dog")
		}
	}
	for _, i := range NameItems(Entities) {
		fmt.Println(i.name)
	}
}
