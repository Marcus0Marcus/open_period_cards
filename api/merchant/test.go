package main

import (
	"fmt"
	"merchant/middleware/util"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(util.GenRandSalt())
	}

}
