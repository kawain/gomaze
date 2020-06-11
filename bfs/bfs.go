package bfs

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

var maze [][]int
var start []int
var goal []int
var q [][]int

func openFileAndMakeMaze() error {
	var tmp [][]byte

	f, err := os.Open("maze.txt")
	if err != nil {
		return errors.New("ファイルがありません")
	}

	s := bufio.NewScanner(f)
	for s.Scan() {
		tmp = append(tmp, []byte(s.Text()))
	}

	if s.Err() != nil {
		return errors.New("ファイル読み込み失敗")
	}

	// SとGがなければエラー
	// スタートとゴールを覚える
	sc := 0
	gc := 0
	l := len(tmp)
	for i := 0; i < l; i++ {
		for n := 0; n < l; n++ {
			if tmp[i][n] == 'S' {
				sc++
				start = []int{i, n}
			} else if tmp[i][n] == 'G' {
				gc++
				goal = []int{i, n}
			}
		}
	}
	if sc != 1 || gc != 1 {
		return errors.New("スタートとゴールを記入してください")
	}

	// maze [][]int 作成
	for i := 0; i < l; i++ {
		tmp2 := []int{}
		for n := 0; n < l; n++ {
			if tmp[i][n] == '#' {
				tmp2 = append(tmp2, 0)
			} else {
				tmp2 = append(tmp2, -1)
			}
		}
		maze = append(maze, tmp2)
	}

	return nil
}

// https://ja.wikipedia.org/wiki/%E5%B9%85%E5%84%AA%E5%85%88%E6%8E%A2%E7%B4%A2
func breadthFirstSearch() {
	// スタートに訪問済みの印を付ける
	maze[start[0]][start[1]] = 1
	// Q に追加
	q = append(q, start)

	for {
		if len(q) == 0 {
			break
		}
		// Q から取り出す
		v := append([]int{}, q[0]...)
		q = q[1:]
		step := maze[v[0]][v[1]]
		// v に接続している点
		up := []int{v[0] - 1, v[1]}
		down := []int{v[0] + 1, v[1]}
		left := []int{v[0], v[1] - 1}
		right := []int{v[0], v[1] + 1}
		// 処理する
		for _, d := range [4][]int{up, down, left, right} {
			if maze[d[0]][d[1]] == -1 {
				maze[d[0]][d[1]] = step + 1
				// Q に追加
				q = append(q, d)
			}
		}
	}
}

func printMaze() {
	mazeText := ""
	l := len(maze)
	for i := 0; i < l; i++ {
		for n := 0; n < l; n++ {
			if maze[i][n] == 0 {
				// 壁
				mazeText += "■"
			} else {
				// 通路
				mazeText += "□"
			}
		}
		mazeText += "\n"
	}
	fmt.Println(mazeText)

	// ステップ表示
	// for i := 0; i < l; i++ {
	// 	for n := 0; n < l; n++ {
	// 		fmt.Printf("%03d ", maze[i][n])
	// 	}
	// 	fmt.Println()
	// }
}

func shortestRoute() {
	// ゴールからスタートへたどる
	var shortest [][]int
	shortest = append(shortest, goal)
	step := maze[goal[0]][goal[1]]
	v := goal

	for {
		if step == 1 {
			break
		}
		// v に接続している点
		up := []int{v[0] - 1, v[1]}
		down := []int{v[0] + 1, v[1]}
		left := []int{v[0], v[1] - 1}
		right := []int{v[0], v[1] + 1}
		// 処理する
		for _, d := range [4][]int{up, down, left, right} {
			if maze[d[0]][d[1]] == step-1 {
				shortest = append(shortest, d)
				v = d
				step = step - 1
				break
			}
		}
	}

	mazeText := ""
	l := len(maze)
	for i := 0; i < l; i++ {
		for n := 0; n < l; n++ {
			if maze[i][n] == 0 {
				// 壁
				mazeText += "■"
			} else {
				f := false
				for _, v := range shortest {
					if i == v[0] && n == v[1] {
						f = true
						break
					}
				}
				// 通路
				if f {
					mazeText += "○"
				} else {
					mazeText += "□"
				}
			}
		}
		mazeText += "\n"
	}
	fmt.Println(mazeText)
}

// Bfs エントリーポイント
func Bfs() error {
	err := openFileAndMakeMaze()
	if err != nil {
		return err
	}

	breadthFirstSearch()
	printMaze()
	shortestRoute()

	return nil
}
