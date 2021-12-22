package main

import "fmt"

func main() {
	res := Solve()
	fmt.Printf("first: %d, last: %d\n", res[0], res[len(res)-1])
}
