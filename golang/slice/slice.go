package main

import (
	"fmt"
)

func main() {
	a := []int{0, 1, 2, 3, 4}
	b := []int{5, 6, 7, 8, 9}
	i := 2
	j := 4

	fmt.Println(a)
	fmt.Println(a[1])

	fmt.Println("接続")
	fmt.Println(append(a, b...))

	fmt.Println("コピー")
	d := make([]int, len(a))
	copy(d, a)
	fmt.Println(d)

	fmt.Println("カット i番目番目までとj番目移行番目以降になる")
	fmt.Println(append(a[:i], a[j:]...))

	a = []int{0, 1, 2, 3, 4}
	fmt.Println("順番が保存されない削除")
	fmt.Println("a[2]を消す")
	a[2] = a[len(a)-1]
	a = a[:len(a)-1]
	fmt.Println(a)

	a = []int{0, 1, 2, 3, 4}
	fmt.Println("間を増やす i番目からj個個の0を入れる")
	fmt.Println(append(a[:i], append(make([]int, j), a[i:]...)...))

	fmt.Println("後ろを伸ばす")
	fmt.Println(append(a, make([]int, j)...))

	fmt.Println("挿入")
	fmt.Println(append(a[:i], append([]int{10}, a[i:]...)...))

	fmt.Println("挿入2")
	b = append(a, 0)
	copy(b[i+1:], b[i:])
	b[i] = 10
	fmt.Println(b)

	fmt.Println("sliceの挿入")
	fmt.Println(append(a[:i], append([]int{10, 11, 12}, a[i:]...)...))

	fmt.Println("push 最後に追加")
	fmt.Println(append(a, 10))

	fmt.Println("pop 最後を取り出す")
	x, a := a[len(a)-1], a[:len(a)-1]
	fmt.Println(x, a)

	fmt.Println("push front/unshift 前に追加")
	fmt.Println(append([]int{10}, a...))

	fmt.Println("pop front/shift 前を削除してずらす")
	x, a = a[0], a[1:]
	fmt.Print(x, a)
}
