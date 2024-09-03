package main

import "fmt"

func main() {
	// Write your code here
	var (
		Bubblegum     float32 = 202
		Toffee        float32 = 118
		IceCream      float32 = 2250
		MilkChocolate float32 = 1680
		Doughnut      float32 = 1075
		Pancake       float32 = 80
	)
	income := Bubblegum + Toffee + IceCream + MilkChocolate + Doughnut + Pancake

	var expenses float32
	var otherExpenses float32

	fmt.Println("Earned amount:\n", "Bubblegum: $202\n", "Toffee: $118\n", "Ice cream: $2250\n", "Milk chocolate: $1680\n", "Doughnut: $1075\n", "Pancake: $80", "\nIncome: $"+fmt.Sprint(income))
	fmt.Println("Staff expenses:")
	fmt.Scanf("%f", &expenses)
	fmt.Println("Other Expenses:")
	fmt.Scanf("%f", &otherExpenses)

	totalExpenses := expenses + otherExpenses
	totalExpenses = -totalExpenses
	var netIncome float32 = income + totalExpenses

	fmt.Println("Net income: $" + fmt.Sprint(netIncome))
}
