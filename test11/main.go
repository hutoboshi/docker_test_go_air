package main

import "fmt"

// メモ化を使ったフィボナッチ数の計算
func fibonacciMemo(n int, memo map[int]int) int {
	if n <= 1 {
		return n
	}
	if val, exists := memo[n]; exists {
		return val
	}
	memo[n] = fibonacciMemo(n-1, memo) + fibonacciMemo(n-2, memo)
	return memo[n]
}

func main() {
	var n int
	fmt.Print("計算したいフィボナッチ数の位置を入力してください: ")
	fmt.Scan(&n)
	memo := make(map[int]int)
	fmt.Printf("Fibonacci(%d) = %d\n", n, fibonacciMemo(n, memo))
}
