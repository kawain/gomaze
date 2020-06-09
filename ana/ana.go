package ana

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

type mazeType [][]int

type worker struct {
	row, col int
	stack    [][]int
}

func new(num int, maze mazeType) *worker {
	var row, col int
	// 候補の数字配列作成
	var tmp []int
	for i := 0; i < num; i++ {
		if i%2 != 0 {
			tmp = append(tmp, i)
		}
	}
	// 候補配列の長さ
	l := len(tmp)
	// ランダムで奇数ポイントを作成
	rand.Seed(time.Now().UnixNano())
	row = tmp[rand.Intn(l)]
	col = tmp[rand.Intn(l)]

	// 最外周のために2つ増やしているので、1加算(忘却ポイント)
	row++
	col++

	w := &worker{
		row:   row,
		col:   col,
		stack: [][]int{{row, col}},
	}

	// スタート地点を0にする
	maze[row][col] = 0

	return w
}

func (w *worker) digging(maze mazeType) int {
	// 現在の位置から上下左右を見て、2つ先が1なら移動
	var row1, col1, row2, col2 int
	var canGo bool
	// 上下左右配列
	t := []string{"上", "下", "左", "右"}

	for {
		// 上下左右配列の長さ
		l := len(t)
		if l == 0 {
			// この場合、行ける場所なし
			canGo = false
			break
		}

		// ランダムで上下左右を見る
		i := rand.Intn(l)

		switch t[i] {
		case "上":
			row2 = w.row - 2
			col2 = w.col
			row1 = w.row - 1
			col1 = w.col
		case "下":
			row2 = w.row + 2
			col2 = w.col
			row1 = w.row + 1
			col1 = w.col
		case "左":
			row2 = w.row
			col2 = w.col - 2
			row1 = w.row
			col1 = w.col - 1
		case "右":
			row2 = w.row
			col2 = w.col + 2
			row1 = w.row
			col1 = w.col + 1
		}

		// 行けない時は上下左右の配列のiを削除
		if maze[row2][col2] == 0 {
			t = remove(t, i)
		} else {
			// この場合、どこかに行けた
			canGo = true
			break
		}
	}

	if canGo {
		// どこかに行けたので、更新する
		maze[row2][col2] = 0
		maze[row1][col1] = 0
		w.row = row2
		w.col = col2
		w.push([]int{row2, col2})
	} else {
		// 行けないのでスタックをポップ
		p, _ := w.pop()
		w.row = p[0]
		w.col = p[1]
	}

	return len(w.stack)
}

func (w *worker) push(num []int) {
	w.stack = append(w.stack, num)
}

func (w *worker) pop() ([]int, error) {
	var r []int
	if len(w.stack) > 0 {
		n := len(w.stack) - 1
		r = w.stack[n]
		w.stack = w.stack[:n]
		return r, nil
	}
	return r, errors.New("スタックは空です")
}

// スライスの要素を削除する関数
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func inputCheck(num int) bool {
	// 5 - 101 までの奇数判定
	if num < 5 || num > 101 {
		return false
	}

	if num%2 == 0 {
		return false
	}

	return true
}

func arrayCreation(num int) mazeType {
	// 2次元配列を作成 最外周のために2つ増やす
	num += 2
	maze := make(mazeType, num)
	for i := 0; i < num; i++ {
		maze[i] = make([]int, num)
	}

	// 最外周以外、初期値を1にする
	for i := 0; i < num; i++ {
		for n := 0; n < num; n++ {
			if i != 0 && i != num-1 && n != 0 && n != num-1 {
				maze[i][n] = 1
			}
		}
	}

	return maze
}

func printMaze(maze mazeType) {
	// 見やすくプリント
	mazeText := ""
	l := len(maze)
	for i := 0; i < l; i++ {
		for n := 0; n < l; n++ {
			// 最外周は描画しない
			if i != 0 && i != l-1 && n != 0 && n != l-1 {
				if maze[i][n] == 1 {
					// 壁
					mazeText += "■"
				} else {
					// 通路
					mazeText += "□"
				}
			}
		}
		// 忘却ポイント
		if i != 0 && i != l-1 {
			mazeText += "\n"
		}
	}

	fmt.Println(mazeText)

	// 文字列置換
	mazeText = strings.Replace(mazeText, "■", "#", -1)
	mazeText = strings.Replace(mazeText, "□", " ", -1)

	// ファイルに書き込む
	err := ioutil.WriteFile("maze.txt", []byte(mazeText), 0666)
	if err != nil {
		fmt.Println(err)
	}
}

// Ana エントリーポイント
func Ana(num int) error {

	if !inputCheck(num) {
		err := errors.New("5以上101以下の奇数を指定してください")
		return err
	}

	m := arrayCreation(num)

	w := new(num, m)

	for {
		i := w.digging(m)
		if i == 0 {
			break
		}
	}

	printMaze(m)

	return nil
}
