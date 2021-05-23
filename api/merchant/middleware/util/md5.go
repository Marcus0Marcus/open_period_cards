package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func Md5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func GenRandSalt() string {
	alpha := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "+", "-", "*", "/", "!", "@", "#", "$"}
	resLen := 6
	res := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < resLen; i++ {
		idx := rand.Intn(len(alpha))
		res = res + alpha[idx]
	}
	return res
}
