package main

import (
	file "GO_GETTING_STARTED/hello/Advent/pkg/fle"
	"fmt"
	"strconv"
	"strings"
)

var gamePlayableReqCubesMap map[string]int = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func main() {
	// reading input file
	inputString := []string(file.ReadFile("input.txt"))
	gameIdsSum := int(getPlayableGamesIdsSum(inputString))

	fmt.Printf("Answer: %d\n", gameIdsSum)
	
}

func getPlayableGamesIdsSum(input []string) int {

	gameIdsAcc := 0
	isGamePlayabale := true
	// iterate over each game
	for gameId, game := range input {
		fmt.Printf("****************Solving game no.: %d -> %s\n", gameId, game)

		// spliting game into words [" "]
		words := strings.Split(game, " ")

		// cube colors list
		cubeColors := [3]string{"red", "green", "blue"}

		// game config map inital
		gameConfigMap := map[string]int{
			"red": 0,
			"green": 0,
			"blue": 0,
		}

		// iterating over words
		for wordIdx, word := range words {
			// iterating over cube colors to find a match in words
			for _, color := range cubeColors {
				if strings.HasPrefix(word, color) {
					n, err := strconv.Atoi(words[wordIdx - 1])
					if err != nil {
						fmt.Println(err)
					}
					gameConfigMap[color] += n
					if strings.HasSuffix(word, ";") || wordIdx == len(words) - 1 {
						temp := map[string]int {
							"red": gameConfigMap["red"],
							"green": gameConfigMap["green"],
							"blue": gameConfigMap["blue"],
						}
						for key := range gameConfigMap {
							gameConfigMap[key] = 0
						}

						for key, value := range temp {
							if value > gamePlayableReqCubesMap[key] {
								isGamePlayabale = false
								fmt.Println("game is not playable:", gameId+1)
								
								break
							}
						}
					}
				}
				
			}

			fmt.Println(gameConfigMap)
		}



		// *************************************************************************

		// // cut out prefix from game ["Game id: "]
		// prefixGameStr := fmt.Sprintf("Game %d: ", gameId+1)
		// game, found := strings.CutPrefix(game, prefixGameStr)
		// fmt.Printf("Found prefix: %t, After cut: %s\n", found, game)

		// // split game sets [;]
		// gameSets := strings.Split(game, ";")

		// // iterate over each game set
		// gameConfigMap := map[string]int{
		// 	"red": 0,
		// 	"green": 0,
		// 	"blue": 0,
		// }
		// for setNum, set := range gameSets {
		// 	fmt.Printf("******************setNum: %d -> %s\n", setNum, set)

		// 	// spliting set into words [" "]
		// 	setWords := strings.Split(set, " ")
			
		// 	// iterate over each word
		// 	for wc, word := range setWords {
		// 		cubeColors := [3]string{"red", "green", "blue"}
		// 		fmt.Printf("%d -> %s\n", wc, word)

		// 		// iterating over cubecolors
		// 		for _, color := range cubeColors {
		// 			// checking if word matches with any of cubecolors
		// 			if strings.HasPrefix(word, color) {
		// 				fmt.Printf("word: %s has suffix: %s\n", word, color)
		// 				num, err := strconv.Atoi(setWords[wc - 1])
		// 				if err != nil {
		// 					fmt.Println(err)
		// 				}
						
		// 				// populating game config map
		// 				gameConfigMap[color] +=  num
		// 			}
		// 		}
				
		// 	}
			

			
			
		// }

		// fmt.Println("*******************gameconfigmap: ",gameConfigMap)

		// for key, value := range gamePlayableReqCubesMap {
		// 	fmt.Println(value)
		// 	gamePlayableReqCubesMap[key] -= gameConfigMap[key]
			
		// }
		// TODO: Just figure out how to accumulate gameids
		
		// iterating over game config to check if its valid game
		// for key, value := range gameConfigMap {
		// 	gameConfigMap[key] = int(gamePlayableReqCubesMap[key] - value)
		// 	if 0 > gameConfigMap[key] {
		// 		isGamePlayabale = false
		// 		fmt.Printf("Game:%d is not playable\n", gameId+1)
		// 		break
		// 	}
			
		// }

		if isGamePlayabale {
			fmt.Printf("Game:%d is playable\n", gameId+1)
			gameIdsAcc += int(gameId + 1)
		}else {
			isGamePlayabale = true
		}
	}

	
	return gameIdsAcc
}



// 12 red cubes, 13 green cubes, and 14 blue cubes
//  {
// 			"red": 12,
// 			"green": 13,
// 			"blue": 14
// 			Mapping of color of cube -> number of cubes
// 	}


// Things to do:
// You need to figure out, which games are playable
// You need to grab their ids
// Sum them up and return

// TODO1: [FIGURE OUT WHICH GAMES ARE PLAYABLE]
// You need to create a array(slice) of maps for each GAME [100 maps]
// Where each map would be storing the mapping of color of cube -> number of cubes