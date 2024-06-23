package main

import (
	"fmt"
	"strconv"

	// "log"
	"os"
	"strings"
)

// 1abc2			 --> 12
// pqr3stu8vwx --> 38
// a1b2c3d4e5f --> 15
// treb7uchet  --> 77

// steps:
// 1. traverse input document line by line -> 1abc2 (solve 1abc2)
// 2. traverse picked line left-to-right check for first numeric-digit and last numeric-digit -> 1 2
// 3. convert those numeric-digits into a decimal -> 12
// 4. If there is only one numeric-digit then, consider a digit 10x of it -> line:4 has only 7 so 7 7x10
// 5. Move to next line then next line until you reach the end of the document
// 6. Accumulate those decimals from each line with the next line's.
// 7. Return the accumulated number

// ideas:
// 2. traverse the line from ltr to get the first number then break out then traverse from rtl to get the last number.
// 3. decimal = first-number x 10 + last-number (1x10 + 2)
// 6. decimal += first-number x 10 + last-number -> solves accumulation problem initialize decimal with 0

func findCalibrationValuesSum(docSlice []string) int {
	accumulator := 0
	firstNumber := 0
	lastNumber := 0
	for _, word := range docSlice {

		//traversing word from ltr
		for _, char := range word {
    	c := string(char)

			n, err := strconv.Atoi(c)
			if err == nil {
				firstNumber = n
				break
			}
		}
	
		firstNumber *= 10
		word = revStr(word)
		//traversing word from rtl
		for _, char := range word {
    	c := string(char)

			n, err := strconv.Atoi(c)
			if err == nil {
				lastNumber = n
				break
			}
		}
		accumulator += firstNumber + lastNumber
	}
	return accumulator
}

func getInpSlice() []string {
	file, err := os.Open("input.txt")
	
	if err != nil {
		fmt.Println(err)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	fileSize := int(fileInfo.Size())

	inpSliceBytes := make([]byte, fileSize)
	count, err := file.Read(inpSliceBytes)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("count: ", count)
	defer file.Close()

	inpString := string(inpSliceBytes[:])
	
	inpSliceStrings := strings.Split(inpString, "\n")

	return inpSliceStrings
}

func main() {
	fmt.Println("trebuchet!")
	// file, err := os.Open("input.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fileInfo, errs := file.Stat()
	// if errs != nil {
	// 	fmt.Println(err)
	// }
	// fileSize := fileInfo.Size()
	// fmt.Println(fileSize)
	// b := make([]byte, fileSize)
	// count, err := file.Read(b)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(count)
	// str1 := string(b[23])
	// if str1 == "\n" {
	// 	fmt.Println("new line")
	// }
	// strBytes := string(b[:])
	// fmt.Println(strBytes)
	// fmt.Println(str1)

	fmt.Println("total: ",findCalibrationValuesSum(getInpSlice()))
}

func revStr(str string) string {
	s := []rune(str)
	i := int(0)
	j := int(len(str)) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
	return string(s)
}