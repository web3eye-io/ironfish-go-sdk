package main

import (
	"fmt"
	"time"
)

const ConnectTimeout = time.Second * 3

func main() {

	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixMilli())
}
