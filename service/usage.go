package service

import (
	"fmt"
	"os"
	"path/filepath"
)

// Usage menu
func Usage() {
	a := filepath.Base(os.Args[0])
	fmt.Println()
	fmt.Println("Usage:", a, "[OPTIONS]")
	fmt.Println()
	fmt.Println("    Duplicate seacher.")
	fmt.Println()
	fmt.Println("    Search duplicates in all files in current directory.")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("    -new  STR        New wordlist file. [Dict_New.dic]")
	fmt.Println("    -ext  STR        File extension. [dic]")
	fmt.Println()
	fmt.Println("    -h               This help.")
	fmt.Println("    -v               Print version.")
}
