package src

import (
	"math/rand"
	"strings"
	"time"
)

func GenerateToken (name string) string {
	var builder strings.Builder
	allstr := "1234567890qwertyuiopasdfghjklzxcvbnm"
	rand.Seed(time.Now().Unix())
	for i := 0; i < 5; i++ {
		n := rand.Intn(len(allstr))
		builder.WriteString(allstr[n:n+1])
	}
	token := builder.String()
	return token
}