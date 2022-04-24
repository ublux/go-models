package gomodels

import (
	"fmt"

	"github.com/ublux/go-models/models"
)

func Add(a, b int) int {
	return a + b
}

func foo() {
	var c1 models.Contact
	fmt.Println(c1)
}
