package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// cards := newDeck()
	// // cards.print()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	cards2 := newDeckFromFile("my_cards1")
	fmt.Println(cards2)

}

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func main() {
// 	err := filepath.Walk("./UAT", func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			fmt.Println(err)
// 			return err
// 		}
// 		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
// 		return nil
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
