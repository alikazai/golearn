package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(s string, r *bufio.Reader) (string, error) {
	fmt.Print(s)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip)", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.additem(name, p)

		fmt.Println("Item Added - ", name, price)
		promptOptions(b)
	case "s":
		b.save()
		fmt.Println("You saved the file", b.name)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(t)

		ff := fmt.Sprintf("Tip amount added : %v\n", b.tip)
		fmt.Println(ff)
		promptOptions(b)
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b)
	}

}

func main() {
	os.Mkdir("bills", 0750)

	mybill := createBill()

	promptOptions(mybill)

	fmt.Println(mybill)
}
