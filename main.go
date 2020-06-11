package main

import (
	"flag"
	"fmt"

	"github.com/kawain/gomaze/ana"
	"github.com/kawain/gomaze/bfs"
)

func main() {
	// 作成
	var c = flag.Int("create", 0, "5以上101以下の奇数を指定")
	// 解答
	var s = flag.Bool("solve", false, "迷路があれば解く")

	flag.Parse()

	if *c > 0 {
		err := ana.Ana(*c)
		if err != nil {
			fmt.Println(err)
		}
	}

	if *s {
		err := bfs.Bfs()
		if err != nil {
			fmt.Println(err)
		}
	}
}
