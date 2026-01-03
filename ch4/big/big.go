package main

import (
	"fmt"
	"math/big"
)

// 금융 프로그램이란 math/big 에서 제공하는 Float 객체로 정밀도를 높이는게 좋음.
func main() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))
}
