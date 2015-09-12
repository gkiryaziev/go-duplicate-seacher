package operations

import (
	"bufio"
	"fmt"
	"os"

	s "../service"
)

func DoDuplicate2(files_list []string, new_file string) error {

	m := map[uint64]bool{}
	readed := 0
	added := 0

	out, err := os.Create(new_file)
	if err != nil {
		return err
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	for _, src_file := range files_list {

		counter := 0
		percent := 0
		total, err := s.CalculateLines(src_file)
		if err != nil {
			return err
		}

		fmt.Printf("\n%s processing: ", src_file)

		in, err := os.Open(src_file)
		if err != nil {
			return err
		}
		defer in.Close()

		scanner := bufio.NewScanner(in)

		for scanner.Scan() {
			line := scanner.Text()
			line_hash := s.GetHashFvn64(line)

			readed++

			if _, seen := m[line_hash]; !seen {
				fmt.Fprintln(writer, line)
				m[line_hash] = true
				added++
			}

			counter++
			if counter == 100000 {
				percent += counter
				fmt.Printf("..%d%%", (percent * 100 / total))
				counter = 0
			}
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	if err := writer.Flush(); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println()
	fmt.Printf("|%-20s|%20d|\n", "Readed", readed)
	fmt.Printf("|%-20s|%20d|\n", "Removed", (readed - added))
	fmt.Printf("|%-20s|%20d|\n", "Result", added)

	return nil
}
