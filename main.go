package main

// func main() {
// 	var card string = "Ace ofSpades"

// 	//compiler inffers string
// 	var othercard = "ax"

// 	//other way inffers string
// 	otherCard := "Abc"
// 	// := ONly ON ASSING NEW VARIABLE, not on next assigment

// 	//THis does not work
// 	//other = "abc"

// 	fmt.Println(card)
// 	fmt.Println(otherCard)
// 	fmt.Println(othercard)
// 	//fmt.Println(other)

// 	//Always need to define a variable
// }

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	hand.print()
	remainingDeck.print()
}
