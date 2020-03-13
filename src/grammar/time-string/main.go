package main

import (
	"fmt"
	"time"
)

func main() {
	createTime := time.Now()
	fmt.Println(createTime)
	fmt.Println(createTime.String())
	fmt.Println(createTime.String()[:19])
}
