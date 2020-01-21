package main

import (
	"fmt"

	"github.com/odysseus/vigenere"
)

func main() {
	msg := ""
	key := ""
	fmt.Println("Digite a menssagem criptografada")
	fmt.Scanf("%s", &msg)
	fmt.Println("Digite a chave para descriptografar")
	fmt.Scanf("%s", &key)
	decoded := vigenere.Decipher(msg, key)
	fmt.Println(decoded)

}
