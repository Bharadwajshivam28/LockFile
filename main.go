package main

import (
	"bytes"
	"fmt"
	"os"
    "github.com/Bharadwajshivam28/LockFile/filecrypt"
	"golang.org/x/term"
	"github.com/fatih/color"
)

const (
	textBold = "\033[1m"
)

func main () {
	message := fmt.Sprintf("%s %s %s", textBold + color.RedString("👋 Hello, I am LockFile"), textBold + color.YellowString("I will Lock Your Files"), textBold + color.GreenString("for safety.. 🔥"))

	fmt.Println(message)
	if len (os.Args) < 2{
		printHelp()
		os.Exit(0)
	}
	function := os.Args[1]

	switch function {
	case "help":
		printHelp()
    case "encrypt":
	    encryptHandle()
	case "decrypt":
	    decryptHandle()
	default:
		fmt.Println(textBold + color.HiYellowString("Run encrypt to encrypt a file and decrypt to decrypt a file.... 🏃"))
		os.Exit(1)
    }
}

func printHelp() {
	fmt.Println("")
	fmt.Println(textBold + color.YellowString("Format to run LockFile is...⬇️"))
	fmt.Println("")
	fmt.Println(textBold + color.WhiteString("\tgo run .encrypt /path/to/your/file"))
	fmt.Println("")
	fmt.Println(textBold + color.YellowString("Available commands are:"))
	fmt.Println(textBold + color.WhiteString("\t encrypt\tEncrypts a file by a password... 🔒"))
	fmt.Println(textBold + color.WhiteString("\t decrypt\tDecrypt a file using a password... 🔑"))
	fmt.Println(textBold + color.WhiteString("\t help\t\tDisplays help text... 😊"))
	fmt.Println("")
}

func encryptHandle() {
    if len(os.Args) < 3 {
		println("")
		println(textBold + color.CyanString("Missing the Path file 😔 run " + color.GreenString("go run . help") + color.CyanString(" to see the Commands")))
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file){
		fmt.Println(textBold + color.YellowString("File not found...✖"))
		os.Exit(0)
	}
	password := getPassword()
	fmt.Println(textBold + color.GreenString("\nEncrypting...⏳"))
	filecrypt.Encrypt(file, password)
	fmt.Println(textBold + color.GreenString("\n file successfully protected...✅"))
}

func decryptHandle() {
	if len(os.Args) < 3 {
		println("")
		println(textBold + color.CyanString("Missing the Path file 😔 run " + color.GreenString("go run . help") + color.CyanString(" to see the Commands")))
		os.Exit(0)
	}
	file := os.Args[2]
	if !validateFile(file){
		fmt.Println(textBold + color.YellowString("File not found...✖"))
		os.Exit(0)
	}
	fmt.Print(textBold + color.GreenString("\nEnter password 🙏"))
	password, _ := term.ReadPassword(0)
	fmt.Println(textBold + color.GreenString("\nDecrypting...⏳"))
	filecrypt.Decrypt(file, password)
	fmt.Print("")
	fmt.Print(textBold + color.GreenString("\n file successfully decrypted...✅"))
	fmt.Print("")
}

func getPassword() []byte{
	fmt.Print("")
	fmt.Print(textBold + color.GreenString("Enter password 🙏"))
	password, _ := term.ReadPassword(0)
	fmt.Print("")
	fmt.Print(textBold + color.GreenString("\nConfirm Password 🙏"))
	password2, _ := term.ReadPassword(0)
	if !validatePassword(password, password2){
		fmt.Print("")
		fmt.Print(textBold + color.GreenString("\nPassword do not match. Please try again 🔁\n"))
		return getPassword()
	}
	return password
}

func validatePassword(password1 []byte, password2 []byte) bool{
	if !bytes.Equal(password1, password2){
		return false
	}
	return true
}

func validateFile(file string) bool{
	if _, err := os.Stat(file); os.IsNotExist(err){
		return false
	}
	return true
}