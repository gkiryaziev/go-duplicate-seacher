package operations

import (
	"bufio"
	"fmt"
	"os"

	"github.com/cheggaaa/pb"
	//"time"

	s "github.com/gkiryaziev/go-duplicate-seacher/service"
)

// DoDuplicate search duplicates
func DoDuplicate(filesList []string, newFile string) error {

	m := map[uint64]bool{}
	readed := 0
	added := 0

	out, err := os.Create(newFile)
	if err != nil {
		return err
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	for _, srcFile := range filesList {

		total, err := s.CalculateLines(srcFile)
		if err != nil {
			return err
		}

		in, err := os.Open(srcFile)
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
			lineHash := s.GetHashFvn64(line)

			readed++

			if _, seen := m[lineHash]; !seen {
				fmt.Fprintln(writer, line)
				m[lineHash] = true
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
	for _, srcFile := range filesList {
		fmt.Println(srcFile)
	}

	fmt.Println("-------------------------------------------")
	fmt.Printf("|%-20s|%20d|\n", "Readed", readed)
	fmt.Printf("|%-20s|%20d|\n", "Removed", (readed - added))
	fmt.Printf("|%-20s|%20d|\n", "Result", added)
	fmt.Println("-------------------------------------------")
	fmt.Println()

	return nil
}
