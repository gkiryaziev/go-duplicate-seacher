package operations

import (
	"bufio"
	"fmt"
	"github.com/cheggaaa/pb"
	"os"
	//"time"

	s "github.com/gkiryaziev/go-duplicate-seacher/service"
)

// DoDuplicate search duplicates
func DoDuplicate(files_list []string, new_file string) error {

	m := map[uint64]bool{}
	var readed int64 = 0
	var added int64 = 0

	out, err := os.Create(new_file)
	if err != nil {
		return err
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	for _, src_file := range files_list {

		total, err := s.CalculateLines(src_file)
		if err != nil {
			return err
		}

		in, err := os.Open(src_file)
		if err != nil {
			return err
		}
		defer in.Close()

		scanner := bufio.NewScanner(in)

		// Progress Bar
		bar := pb.New64(total)
		bar.ShowPercent = true
		bar.ShowBar = true
		bar.ShowCounters = true
		bar.ShowTimeLeft = true
		//bar.SetRefreshRate(time.Millisecond * 100)
		//bar.Format("<.- >")
		bar.Start()

		for scanner.Scan() {
			line := scanner.Text()
			line_hash := s.GetHashFvn64(line)

			readed++

			if _, seen := m[line_hash]; !seen {
				fmt.Fprintln(writer, line)
				m[line_hash] = true
				added++
			}
			bar.Increment()
		}
		bar.Finish()

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	fmt.Println("\nProcessed files:")
	fmt.Println("-------------------------------------------")
	for _, src_file := range files_list {
		fmt.Println(src_file)
	}

	fmt.Println("-------------------------------------------")
	fmt.Printf("|%-20s|%20d|\n", "Readed", readed)
	fmt.Printf("|%-20s|%20d|\n", "Removed", (readed - added))
	fmt.Printf("|%-20s|%20d|\n", "Result", added)
	fmt.Println("-------------------------------------------\n")

	return nil
}
