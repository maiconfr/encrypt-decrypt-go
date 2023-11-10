package main

import (
	"encrypt-decrypt/functions"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

var MySecret string

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Adicionar o arquivo de configuração .env \n", err.Error())
		panic(err)
	}

	MySecret = os.Getenv("MySecret")

	StringToEncrypt := "Encrypting this string"
	// To encrypt the StringToEncrypt
	encText, err := functions.Encrypt(StringToEncrypt, MySecret, bytes)
	if err != nil {
		fmt.Println("error encrypting your classified text: ", err)
	}
	fmt.Println(encText)
	// To decrypt the original StringToEncrypt
	decText, err := functions.Decrypt("Li5E8RFcV/EPZY/neyCXQYjrfa/atA==", MySecret, bytes)
	if err != nil {
		fmt.Println("error decrypting your encrypted text: ", err)
	}
	fmt.Println(decText)
}
