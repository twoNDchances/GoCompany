package models

import "fmt"

type StringOuter interface {
	ToString() string
}

func PrintEntity(entity StringOuter) {
	fmt.Println(entity.ToString())
}
