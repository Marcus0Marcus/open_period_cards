package main

import (
	"fmt"
	"open_period_cards/middleware/util"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(util.GenRandSalt())
	}

}
