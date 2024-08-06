package main

import (
	file "GO_GETTING_STARTED/hello/Advent/pkg/fle"
	"fmt"
	"strconv"
)

func main() {
	input := file.ReadFile("input.txt")
	// fmt.Printf("Sum: %d\n", getPartsSum(input))
	fmt.Printf("Gear Sum: %d\n", getGearsSum(input))
}
// *****************************
func getGearsSum(s []string) int {
	sum := 0

	for lIdx, l := range s {
		fmt.Println("Line & LineNum: ", l, lIdx)
		
		for rIdx, r := range l {
			fmt.Println("rune & runeNum: ", r, rIdx)

			if r == '*' {
				fmt.Println("found asteric at: ", rIdx)

				var (
					pL string
					nL string
				)
				if lIdx != 0 {
					pL = s[lIdx - 1]
				}

				if lIdx != len(s) - 1 {
					nL = s[lIdx + 1]
				}

				isGear, parts := checkIsGear(pL, l, nL, rIdx)

				fmt.Println("isGear && parts********************", isGear, parts)

				if isGear {
					ps := parts[0] * parts[1]
					sum = sum + ps
				}
				
			}

		}
	}

	return sum
}

func checkIsGear(pL, cL, nL string, gLoc int) (bool, []int) {
	var (
		isGear bool
		parts []int
	)

	// checking for left wall
	if gLoc >= 0 {
		// checking for part on left side
		if cL[gLoc - 1] >= '0' && cL[gLoc - 1] <= '9' {
			n, _ := grabNum(cL, gLoc - 1)
			parts = append(parts, n)
		}
	}

	// checking for right wall
	if gLoc < len(cL) - 1 {
		// checking for part on right side
		if cL[gLoc + 1] >= '0' && cL[gLoc + 1] <= '9' {
			n, _ := grabNum(cL, gLoc + 1)
			parts = append(parts, n)
		}
	}

	// checking for top wall
	if pL != "" {
		// checking for part in upper wall left side
		if len(parts) < 2 && gLoc >= 0 {
			 if pL[gLoc - 1] >= '0' && pL[gLoc - 1] <= '9' {
					n, _ := grabNum(pL, gLoc - 1)
					parts = append(parts, n)
				}
		}

		// checking for part in upper wall middle
		if len(parts) < 2 {
			 if pL[gLoc] >= '0' && pL[gLoc] <= '9' {
					n, _ := grabNum(pL, gLoc)
					parts = append(parts, n)
				}
		}

		// checking for part in upper wall right side
		if len(parts) < 2 && gLoc < len(pL) - 1 {
			 if pL[gLoc + 1] >= '0' && pL[gLoc + 1] <= '9' {
					n, _ := grabNum(pL, gLoc + 1)
					parts = append(parts, n)
				}
		}
	}

	// checking for bottom wall
	if nL != "" {
		// checking for part in lower wall left side
		if len(parts) < 2 && gLoc >= 0 {
			 if nL[gLoc - 1] >= '0' && nL[gLoc - 1] <= '9' {
					n, _ := grabNum(nL, gLoc - 1)
					parts = append(parts, n)
				}
		}

		// checking for part in bottom wall middle
		if len(parts) < 2 {
			 if nL[gLoc] >= '0' && nL[gLoc] <= '9' {
					n, _ := grabNum(nL, gLoc)
					parts = append(parts, n)
				}
		}

		// checking for part in lower wall right side
		if len(parts) < 2 && gLoc < len(pL) - 1 {
			 if nL[gLoc + 1] >= '0' && nL[gLoc + 1] <= '9' {
					n, _ := grabNum(nL, gLoc + 1)
					parts = append(parts, n)
				}
		}
	}

	if len(parts) == 2 {
		isGear = true
	}

	return isGear, parts
}

// *****************************
func getPartsSum(s []string) int {
	sum := 0

	for lIdx, l := range s {
		fmt.Printf("Line%d:, %s\n", lIdx, l)
		psl := parseLine(l, lIdx, s)
		sum += psl
	}

	return sum
}

func parseLine(s string, lNum int, inpStr []string) int {
	sum := 0
	fmt.Println(len(s))
	line := []rune(s)

	for rIdx := 0; rIdx < len(line); rIdx++ {
		fmt.Println("if starts here")
		// checking if its a num
		if line[rIdx] >= '0' && line[rIdx] <= '9' {
			fmt.Println("if starts and condition checked")
			// checking if symbol present
			isP := checkSymbolPresence(inpStr, lNum, rIdx, len(line))
			if isP {
				// grab the num
				n, eIdx := grabNum(s, rIdx)
				// changing rune idx, to num end idx + 1
				if eIdx == len(line) - 1 {
					rIdx = eIdx
				} else {
					rIdx = eIdx + 1
				}
				fmt.Println("#####################",rIdx, eIdx, len(line))
				sum += n
			}
		}

		fmt.Println("forloop ends here", string(line[rIdx]), rIdx)
	}

	return sum
}


func checkIsSymbol(b byte) bool {
	r := rune(b)
	return !(r >= '0' && r <= '9') && !(r == '.')
}

func checkSymbolPresence(s []string, lNum int, nIdx int, lLen int) bool {
		fmt.Println("just cecking")
	if lNum != 0 {
		if checkIsSymbol(s[lNum - 1][nIdx]) {
			return true
		}

		if nIdx != 0 {
			if checkIsSymbol(s[lNum - 1][nIdx - 1]) {
			return true
			}
		}

		if lLen - 1 != nIdx {
			if checkIsSymbol(s[lNum - 1][nIdx + 1]) {
				return true
			}
		} 
	}
	if nIdx != 0 {
		if checkIsSymbol(s[lNum][nIdx - 1]) {
			return true
		}
	}

	if lLen - 1 != nIdx {
		if checkIsSymbol(s[lNum][nIdx + 1]) {
			return true
		}
	}

	if lNum != len(s) - 1 {
		if checkIsSymbol(s[lNum + 1][nIdx]) {
			return true
		}
		
		if nIdx != 0 {
			if checkIsSymbol(s[lNum + 1][nIdx - 1]) {
			return true
			}
		}

		if lLen - 1 != nIdx {
			if checkIsSymbol(s[lNum + 1][nIdx + 1]) {
				return true
			}
		}
	}

	return false
}



func grabNum(line string, nIdx int) (int, int) {

	lHalf := line[nIdx + 1 :]
	eHalf := line[: nIdx]
	
	var nEIdx int
	for i, r := range lHalf {
		if !(r >= '0' && r <= '9') {
			nEIdx = i + nIdx
			break
		}
	}
	

	if nEIdx == 0 {
		lHLR := rune(lHalf[len(lHalf) - 1])
		if lHLR >= '0' && lHLR <= '9' {
			nEIdx = nIdx + len(lHalf)
		}
	}
	

	
 
	

	var nSIdx int
	rEHalf := reverse(eHalf)
	for i, r := range rEHalf {
		if !(r >= '0' && r <= '9') {
			nSIdx = nIdx - i
			break
		}
	}

	n, err := strconv.Atoi(line[nSIdx:nEIdx + 1])
	fmt.Printf("atio returned: %d\n", n)
	if err != nil {
		fmt.Print(err)
	}
	return n, nEIdx
}

// function, which takes a string as 
// argument and return the reverse of string. 
func reverse(s string) string { 
    rns := []rune(s) // convert to rune 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
  
        // swap the letters of the string, 
        // like first with last and so on. 
        rns[i], rns[j] = rns[j], rns[i] 
    } 
  
    // return the reversed string. 
    return string(rns) 
} 