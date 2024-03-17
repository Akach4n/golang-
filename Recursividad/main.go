package main

import "fmt"

func main() {
	recursi(10)
}

func recursi(num int) {
	if num >= 0 {
		fmt.Println(num)
		recursi(num - 1)
	}

}
