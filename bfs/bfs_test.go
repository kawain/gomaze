package bfs

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestCreation(t *testing.T) {
	b, _ := ioutil.ReadFile("../maze.txt")

	maze := arrayCreation(b)

	fmt.Printf("%v\n", maze)

	t.Fatalf("正解は %v\n", 1)
}
