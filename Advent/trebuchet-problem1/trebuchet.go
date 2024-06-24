package main

import (
	"fmt"
	"strconv"

	// "log"
	"os"
	"strings"
)

	var strBuffer strings.Builder


// 1abc2			 --> 12
// pqr3stu8vwx --> 38
// a1b2c3d4e5f --> 15
// treb7uchet  --> 77

// pre-processing:
// #Convert letter-word -> numeric-word [Build an api called (letterWord string) convLetterToNumeric() (numericWord string)]

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

var lToN map[string]string = map[string]string{
	"one": "1",
	"two": "2",
	"three": "3",
	"four": "4",
	"five": "5",
	"six": "6",
	"seven": "7",
	"eight": "8",
	"nine": "9",
}
var letterNumToIdx map[string]int = map[string]int{
	"one": -1,
	"two": -1,
	"three": -1,
	"four": -1,
	"five": -1,
	"six": -1,
	"seven": -1,
	"eight": -1,
	"nine": -1,
}
var numericNumToIdx map[int]int = map[int]int{
	1: -1,
	2: -1,
	3: -1,
	4: -1,
	5: -1,
	6: -1,
	7: -1,
	8: -1,
	9: -1,
}

// sample input
// two1nine 					-> 219
// eightwothree				-> 823
// abcone2threexyz		-> abc123xyz
// xtwone3four				-> x2134
// 4nineeightseven2		-> 49872
// zoneight234				-> z18234
// 7pqrstsixteen			-> 7pqrst6teen  5max 3min

func decToBin(n int) int {
	p := int(1)
	bin := int(0)
	for i := 0; n != 0; i++ {
		r := int(n % 2)
		n = int(n / 2)

		bin += r * p
		p *= 10
	}

	return bin
}

func getSs(s string, str, end int) string {
	strBuffer.WriteString("")
	strSlice := []rune(s)
	for i := str; i < end; i++ {
		strBuffer.WriteString(string(strSlice[i]))
	}

	ss := strBuffer.String()
	strBuffer.Reset()

	return ss
}


func findLetterNum(word string, numStr string) int {
	nLen := len(numStr) // 3
	i := 0
	
	for i <= len(word) - nLen {
		j := i + nLen
		ss := string(getSs(word, i, j))
		
		if ss == numStr {
			fmt.Println("matched!")
			return i
		}
		i++
	}

	return -1
}

func convLetterToNumeric(word string) string {
	strBuffer.Reset()
	subWord := ""
	for i, char := range word{
		strBuffer.WriteString(string(char))
		if i >= 3 {
			for key, value := range lToN{
				subWord = strings.Replace(strBuffer.String(), key, value, -1)
				strBuffer.Reset()
				strBuffer.WriteString(subWord)
			}
			// fmt.Println("subword:",subWord)
		}
	}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	// wordSlice := []rune(letWord)


	// nOI := 1
	// for i := 0; i < len(wordSlice); i++ {
	// 	nOI *= 2
	// }
	// nOI--

	// for i := 0; i < nOI; i++ {
	// 	fmt.Println(decToBin(i))
	// }

	return subWord
}

func findCalibrationValuesSum(docSlice []string) int {
	accumulator := 0
	firstNumericNum := 0
	firstNumericNumIdx := -1
	lastNumericNum := 0
	lastNumericNumIdx := -1
	// firstLetterNum := 0
	// firstLetterNumIdx := -1
	// lastLetterNum := 0
	// lastLetterNumIdx := -1
	fmt.Println(len(docSlice))
	fmt.Println(docSlice)

	for i, word := range docSlice {
		fmt.Println(i, ": ", word)
		// for key, value := range lToN {
		// 	word = strings.Replace(word, key, value, -1)

		// 	// fmt.Println(word)
		// }
		// fmt.Println(word)

		// idxOLN := -1
		// for key := range lToN {


		// 	// idxOLN = findLetterNum(word, key)
		// 	// if idxOLN != -1 {
		// 	// 	// fmt.Println(idxOLN)
		// 	// 	// fmt.Println(value)
		// 	// 	if letterNumToIdx[key] == 2
		// 	// }

			
		// }
		// fmt.Println(letterNumToIdx)

		word = convLetterToNumeric(word)
		fmt.Println("word:",word)



		//traversing word from ltr
		for i, char := range word {
    	c := string(char)

			n, err := strconv.Atoi(c)
			if err == nil {
				firstNumericNum = n
				firstNumericNumIdx = i
				break
			}
		}
		fmt.Println("firstNum:",firstNumericNum)
		fmt.Println("firstNum idx:",firstNumericNumIdx)
	
		firstNumericNum *= 10
		word = revStr(word)
		//traversing word from rtl
		for i, char := range word {
    	c := string(char)

			n, err := strconv.Atoi(c)
			if err == nil {
				lastNumericNum = n
				lastNumericNumIdx = i
				break
			}
		}
		fmt.Println("lastNum:",lastNumericNum)
		fmt.Println("lastNumidx:",lastNumericNumIdx)
		accumulator += firstNumericNum + lastNumericNum
		fmt.Println("acc:", accumulator)
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
	// convLetterToNumeric("abc")
	// fmt.Println(getSs("two1nine", 0, 3))


// two1nine 					-> 219
// eightwothree				-> 823
// abcone2threexyz		-> abc123xyz
// xtwone3four				-> x2134
// 4nineeightseven2		-> 49872
// zoneight234				-> z18234
// 7pqrstsixteen			-> 7pqrst6teen  5max 3min

	fmt.Println(findLetterNum("abseight", "eight"))
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