package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome to my Pizza App"
	welcome2 := "Please rate our pizza between 1 and 10"
	fmt.Println(welcome + "\n" + welcome2)

	reader := bufio.NewReader(os.Stdin)

	//  \n to stop reading
	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating, ", input)

	// ! trim space or enter
	trimInput := strings.TrimSpace(input)
	numRating, err := strconv.ParseFloat(trimInput, 64)

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}

}
