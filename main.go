package main

import (
	"fmt"
	"github.com/model"
	"github.com/satori/go.uuid"
)

func main() {
	uid := uuid.NewV4()
	fmt.Println(uid)
	model.Hello()
}
