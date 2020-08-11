package main

import (
	"fmt"
	"math/rand"
  "time"
)

func random() int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 255
	reslt := rand.Intn(max - min + 1) + min 
	return reslt
}

func main() {
	concatenated := fmt.Sprintf("rgb(%d, %d, %d)", random(), random(), random())
	fmt.Println(concatenated)
}
