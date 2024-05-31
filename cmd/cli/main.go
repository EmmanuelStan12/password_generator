package main

import (
	"flag"
	"fmt"

	"github.com/EmmanuelStan12/password_generator/generator"
)

var (
	length    = flag.Int(generator.LENGTH, 16, "Length of the password")
	lowercase = flag.Bool(generator.LOWERCASE, true, "Include lowercase letters")
	uppercase = flag.Bool(generator.UPPERCASE, true, "Include uppercase letters")
	symbols   = flag.Bool(generator.SYMBOLS, true, "Include symbols")
	numbers   = flag.Bool(generator.NUMBERS, true, "Include numbers")
)

func main() {
	flag.Parse()
	defer func() {
		error := recover()
		if error != nil {
			fmt.Println(error)
		}
	}()
	password := generator.GeneratePassword(*length, *lowercase, *uppercase, *symbols, *numbers)
	fmt.Printf("Password: %s\n", password)
}
