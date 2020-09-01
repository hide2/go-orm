package main

import (
	"fmt"
	"strings"
)

func main() {
	a := []string{"aaa", "bbb"}
	fmt.Println(strings.Join(a, ","))
}
