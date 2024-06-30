package file

import (
	"bufio"
	"fmt"
	"os"
)


func ReadFile(fN string) []string  {
	file, err := os.Open(fN)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sliceOfPrintableStrs []string
	for scanner.Scan() {
		sliceOfPrintableStrs = append(sliceOfPrintableStrs, scanner.Text())
		
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return sliceOfPrintableStrs
}