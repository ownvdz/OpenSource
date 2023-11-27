package main

import "fmt"

func status(name string) {
	balls := map[string]int{"성기훈": 20, "오일남": 0}
	// ball := balls[name]
	ball, exists := balls[name]
	fmt.Println(ball, exists)

	if !exists {
		fmt.Println(name, "님은 게임에 참여하실 수 없습니다")
	} else if ball < 1 {
		fmt.Println(name, "님은 ", balls[name], "개로 탈락")
	} else {
		fmt.Println(name, "님은 게임에서 승리하셨습니다")
	}
}

func main() {
	status("오일남")
	status("강철")
	status("성기훈")

	// balls := make(map[string]int)
	// // var balls map[string]int

	// fmt.Printf("%#v\n", balls)
	// balls["성기훈"] = 20
	// fmt.Println(balls["성기훈"])
	// fmt.Println(balls["오일남"])

}
