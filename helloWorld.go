package main

import "fmt"

const (
	message4 = "#4 The answer to life is %d \n\n"
	answer4  = 105

	message6 = "#6 iota exampled:  %d  %d  %d\n"
	answer6a = iota // Auto Increment variable for enums
	answer6b
	answer6c
)

var (
	message5 = "#5 The answer to life is %d \n\n"
	answer5  = 105

	message7 = "#7 The float point exmple %.2f\n\n" // only 2 pooints of precision
	answer7  = 3.1468234

	message8 = "#8 The float point exmple %.2f\n\n" // only 2 pooints of precision
	answer8  = float64(3.1468234)

	message9 = "#9 The unsigned int exmple %d\n\n" // only 2 pooints of precision
	answer9  = uint(32)                            //  uint8() uint32()  uint64()

	message10 = "isTrue is %t\n\n"
	answer10  = true
)

func main() {

	var message1 string

	message1 = "#1 Hello World!\n\nThese Nuts\n\n"

	message2 := "#2 Hello World!\n\nSecond Style\n\n"

	fmt.Printf(message1)
	fmt.Printf(message2)

	message3 := "#3 The answer to life is %d \n\n"
	answer := 42

	fmt.Printf(message3, answer)
	fmt.Printf(message4, answer4)
	answer5 := answer5 + 1
	answer5 += 5
	fmt.Printf(message5, answer5)
	fmt.Printf(message6, answer6a, answer6b, answer6c)
	fmt.Printf(message7, answer7)
	fmt.Printf(message8, answer8)
	fmt.Printf(message9, answer9)
	fmt.Printf(message10, answer10)

}
