package main

import (
	"fmt"

	"github.com/lixiangzhong/traceroute"
)

func main() {
	t := traceroute.New("www.github.com")
	t.LocalAddr = "0.0.0.0"
	result, err := t.Do()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range result {
		fmt.Println(v)
	}
}
