package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	inp, err := r.ReadString('\n')
	return strings.TrimSpace(inp), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add, s - save, t - tip): ", reader)
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Price amount: ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("price must be a number")
			promptOption(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added - ", name, price)
		promptOption(b)
	case "s":
		b.saveBill()
		fmt.Println("File saved", b.name)
	case "t":
		tip, _ := getInput("Tip amount ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("tip must be a number")
			promptOption(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added - ", tip)
		promptOption(b)
	default:
		fmt.Println("Re enter from provided options")
		promptOption(b)
	}
}

func main() {
	bill := createBill()
	promptOption(bill)
}
