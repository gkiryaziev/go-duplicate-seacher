package main

import (
	"fmt"
	"os"
	"time"

	o "./operations"
	s "./service"
)

func main() {

	// variables
	var new_file string = "Dict_New.dic"
	var file_ext = ".dic"
	var version = "0.1.2"

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
			new_file = os.Args[k+1]
		case "-ext":
			err := s.CheckArgs(len(os.Args), k)
			s.CheckError(err)
			file_ext = "." + os.Args[k+1]
		}
	}

	// start time
	start := time.Now()

	files_list, err := s.SearchFilesInDir(file_ext, "./")
	s.CheckError(err)

	fmt.Println()
	fmt.Println(len(files_list), "files found.")

	err = o.DoDuplicate2(files_list, new_file)
	s.CheckError(err)

	// elapsed time
	elapsed := time.Since(start)
	fmt.Println("\nElapsed time: ", elapsed)
}
