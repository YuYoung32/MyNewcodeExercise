package main

import (
	"fmt"
	"regexp"
)

func main() {
	example := `http:////\\//\\\\//\\\\/www.good\\//\\.com`
	reg, _ := regexp.Compile(":[\\\\/]*")
	res := reg.ReplaceAllString(example, "://")
	fmt.Println(res)
}
