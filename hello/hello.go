package main

import (
	"fmt"
)

func main() {
	// TODO: FILE -> CREATE, WRITE, AND DELETE ARE LEFT

	// file, err := os.Open("testing.txt")

	// if err != nil {
	// 	fmt.Println(err)
	// }
	
	// defer file.Close()
	
	// contentSize, err := file.Stat()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fileContent := make([]byte, contentSize.Size())

	// Count, err := file.Read(fileContent)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("count: %d \n", Count)
	// fmt.Printf("file: %s \n`", fileContent)

	string := make([]byte, 5)
	string = []byte("HELLO")

	fmt.Println("string bytes: ", &string)
	string = []byte("A")
	fmt.Println("string bytes: ", &string)
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	fmt.Println(sample)















	// var i int
	// _ = i // a way to leave a variable empty
	// // fmt.Println(i)
	// // // var a
	// // // num1 := 5
	// // // // num2 := 2
	// // // var num3 int = 1
	// // fmt.Println(quote.Go())
	// // fmt.Println(sub(5, 5))

	// // var num int
	// // var num2 int
	// // fmt.Println("Enter number:")
	// // fmt.Scan(&num, &num2)
	// // fmt.Printf("Number1: %d Number2: %d \n", num, num2)

	// var x,j int

  // fmt.Print("Type two numbers: ")
  // fmt.Scanln(&x, &j)
  // fmt.Println("Your numbers are:", x, "and", j)
}
