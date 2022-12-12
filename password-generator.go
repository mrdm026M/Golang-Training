package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func main() {
	rand.Seed(time.Now().Unix())

	var totalPasswordLength int
	fmt.Print("Enter your password length: ")
	fmt.Scan(&totalPasswordLength)

	var minLowerChar int
	fmt.Printf("Enter number of lowercase letter (0 - %d): ", totalPasswordLength)
	fmt.Scan(&minLowerChar)

	if minLowerChar > totalPasswordLength {
		fmt.Printf("Entered number of lowercase letter is greater then total password length \n")
		os.Exit(1)
	}

	var minUpperChar int
	fmt.Printf("Enter number of uppercase letter (0 - %d): ", totalPasswordLength-minLowerChar)
	fmt.Scan(&minUpperChar)

	if minUpperChar > totalPasswordLength-minLowerChar {
		fmt.Printf("Entered number of uppercase letter is greater then total avaiable password length \n")
		os.Exit(1)
	}

	var minSpecialChar int
	fmt.Printf("Enter number of special letter (0 - %d): ", totalPasswordLength-(minLowerChar+minUpperChar))
	fmt.Scan(&minSpecialChar)

	if minSpecialChar > totalPasswordLength-(minLowerChar+minUpperChar) {
		fmt.Printf("Entered number of special letter is greater then total avaiable password length \n")
		os.Exit(1)
	}

	var minNumberSet int
	fmt.Printf("Enter number of numbers (0 - %d): ", totalPasswordLength-(minLowerChar+minUpperChar+minSpecialChar))
	fmt.Scan(&minNumberSet)

	if minNumberSet > totalPasswordLength-(minLowerChar+minUpperChar+minSpecialChar) {
		fmt.Printf("Entered number of numbers is greater then total avaiable password length \n")
		os.Exit(1)
	}

	password := generatePassword(totalPasswordLength, minLowerChar, minUpperChar, minSpecialChar, minNumberSet)
	fmt.Println(password)
}

func generatePassword(totalPasswordLength, minLowerChar, minUpperChar, minSpecialChar, minNumberSet int) string {
	var password strings.Builder

	//Set lowercase
	for i := 0; i < minLowerChar; i++ {
		random := rand.Intn(len(lowerCharSet))
		password.WriteString(string(lowerCharSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperChar; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNumberSet; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	remainingLength := totalPasswordLength - minLowerChar - minUpperChar - minSpecialChar - minNumberSet
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
