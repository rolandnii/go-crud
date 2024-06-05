package services

import (
	"fmt"
	"math/rand"

)


func GenerateToken(num int) string {

	var token string

	for i := 0; i < num; i++ {
		token  = token + fmt.Sprint((rand.Intn(9)))
	}

	return token
}




