package main

import (
	"flag"
	"fmt"

	"github.com/kawain/gomaze/ana"
)

func main() {
	var i = flag.Int("make", 0, "5以上101以下の奇数を指定してください")

	flag.Parse()

	if *i > 0 {
		err := ana.Ana(*i)
		if err != nil {
			fmt.Println(err)
		}
	}
}
