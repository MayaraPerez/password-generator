package password

import (
	"fmt"
	"math/rand"
	"time"
)

func RandonPassword(tipo string, comprimento int) (string, error) {
	var charSet string
	switch tipo {
	case "letras":
		charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	case "números":
		charSet = "0123456789"
	case "símbolos especiais":
		charSet = "!@#$%^&*()_-+=<>?/|"
	case "letras, números e símbolos especiais":
		charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_-+=<>?/|"
	default:
		return "", fmt.Errorf("tipo inválido")
	}

	rand.Seed(time.Now().UnixNano())

	password := ""

	for i:=1; i<= comprimento; i++ {
		randomIndex := rand.Intn(len(charSet))
		password += string(charSet[randomIndex])
	}
	return password, nil
	
}