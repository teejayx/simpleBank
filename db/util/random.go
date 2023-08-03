package util

import (
	"math/rand"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"


func init() {
	rand.Seed(time.Now().UnixNano())
}


//randomin generats a random integer between min and max

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // max-min+1 is the range
}

func RandomString(n int) string {
	var letters = []rune(alphabet)

	s := make([]rune, n)

	for i:= range s {
		s[i] = letters[rand.Intn(len(letters))]
	}

	return string(s)
}

func RandomOwner() string {	
	return RandomString(6)
}

func RandomMoney() int64 {	
	return randomInt(0, 1000)
}	

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]							
}

func RandomEmail() string {
	return RandomString(6) + "@gmail.com"
}		





