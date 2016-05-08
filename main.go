package main

import (
	"fmt"
	"os"
	"time"

	o "github.com/gkiryaziev/go-duplicate-seacher/operations"
	s "github.com/gkiryaziev/go-duplicate-seacher/service"
)

func main() {

	// variables
	newFile := "Dict_New.dic"
	fileExt := ".dic"
	version := "0.1.4"

	// args
	for k, arg := range os.Args {
		switch arg {
		case "-h":
			s.Usage()
			return
		case "-v":
			fmt.Println(version)
			return
		case "-new":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			newFile = os.Args[k+1]
		case "-ext":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			fileExt = "." + os.Args[k+1]
		}
	}

	// start time
	start := time.Now()

	filesList, err := s.SearchFilesInDir(fileExt, "./")
	s.CheckError(err)

	fmt.Println()
	fmt.Println(len(filesList), "files found.")
	fmt.Println()

	err = o.DoDuplicate(filesList, newFile)
	s.CheckError(err)

	// elapsed time
	elapsed := time.Since(start)
	fmt.Println("\nElapsed time: ", elapsed)
}
