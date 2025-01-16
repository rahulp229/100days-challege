package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Enter your name")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	Teams := make(map[string][]string, 4)

	Teams["Red"] = []string{"Player-A", "Player-B", "Player-C", "Player-D"}
	Teams["Blue"] = []string{"Player-E", "Player-F", "Player-G", "Player-H"}
	Teams["Green"] = []string{"Player-I", "Player-J", "Player-K", "Player-L"}
	Teams["Yellow"] = []string{"Player-M", "Player-N", "Player-O", "Player-P"}

	fmt.Println("Welcome to Ludo Game")
	fmt.Println("Hello", name)
	checkPost := make(map[string]int)
	checkPost["1st"] = 8
	checkPost["2nd"] = 21
	checkPost["3rd"] = 34
	checkPost["4th"] = 47
	checkPost["final"] = 56
	//1st check post at 8th position
	//2nd check post at 21th position
	//3rd check post at 34th position
	//4th check post at 47th position
	//final path to reach destination at 56th position
	//oppent home postion yellow - 13th position
	//oppent home postion green - 26th position
	//oppent home postion red - 39th position
	rules := []string{"if all players reach final postion then that team is won",
		"if all players of any team not reach to final position then that team is lost"}

	fmt.Println(rules)
	game := make(map[string]bool)
	for key, v := range Teams {
		//playing the game
		fmt.Println(key, "Team", v)
		//gostacticcheck
		if _, ok := game[key]; !ok {
			fmt.Println(key, "Team", v)
			game[key] = true
		}
		//gostacticcheck
		if checkPost["final"] == 56 {
			fmt.Println("....")
			time.Sleep(2 * time.Second)
			fmt.Println("Game is over")

			break
		}
		if game[key] == true {
			fmt.Println(key, "Team", v)
			game[key] = false
		}
	}

}

func generateRandomNumber(min, max int) int {
	seed := 123456
	rand.New(rand.NewSource(int64(seed)))
	return rand.Intn(max-min) + min
}
