package main

import (
	"fmt"
	"go-design-pattern/strategy"
)

func main() {
	strategy.InitMessageStrategy()
	for i := 1; i < 5; i++ {
		fmt.Println(strategy.Message[fmt.Sprintf("%v", i)].Send())
	}

}
