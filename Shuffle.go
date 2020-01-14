package main

import (
	"fmt"
	"math/rand"
	"time"
)

var names = []string{
	"三山國王",
	"耶穌",
	"佛祖",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(2))
}
