package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var onlyOnce sync.Once
var dice = []int{1, 2, 3, 4, 5, 6}

type Player struct {
	Name  string
	Score int
}

func main() {

	var input int
	var name string
	round := 4
	// for {
	// var winner Player
	var players []Player
	fmt.Println("How Many players")
	// for _, each := range menu {
	// 	fmt.Println(each)
	// }
	fmt.Print("===========\ninput: ")
	fmt.Scanf("%d", &input)
	for i := 1; i <= input; i++ {
		var player Player
		mess := fmt.Sprintf("===========\nplayer %d name: ", i)
		fmt.Print(mess)
		fmt.Scanf("%s", &name)
		player.Name = name
		players = append(players, player)
	}
	// log.Println(players)
	for j, each := range players {
		for i := 1; i <= round; i++ {
			num := rollDice()
			each.Score = CountScore(each.Score, num)
			mess := fmt.Sprintf("Player Name: %s, count rolling the dice : %d , result: %d, score : %d \n", each.Name, i, num, each.Score)
			players[j].Score = each.Score
			fmt.Print(mess)

		}
	}
	winners := GetWinner(players)
	for _, each := range winners {
		mess := fmt.Sprintf("The winner is player: %s ,with score :%d", each.Name, each.Score)
		fmt.Print(mess)
	}

}

func GetWinner(players []Player) []Player {
	var winners []Player
	var winner Player
	max := 0
	for _, each := range players {
		if each.Score >= max {
			max = each.Score
			winner.Name = each.Name
			winner.Score = each.Score
			winners = append(winners, winner)
		}
	}
	return winners
}

func CountScore(score, dice int) int {
	var point int
	switch dice {
	case 1, 3, 5:
		point = 10
	default:
		point = -5
	}

	return score + point
}

func rollDice() int {

	onlyOnce.Do(func() {
		rand.Seed(time.Now().UnixNano()) // only run once
	})

	return dice[rand.Intn(len(dice))]
}
