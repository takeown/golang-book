package main

import "fmt"

func main() {
	day := "thursday"

	switch day {
	case "monday", "tuesday":
		fmt.Println("월 화는 수업 가는날")
	case "wednesday", "thursday", "friday":
		fmt.Println("수목금은 실습 나가는날")
	}
}
