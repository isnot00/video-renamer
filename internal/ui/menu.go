package ui

import "fmt"

func ShowMenu() int {

	fmt.Println(Purple + "══════════════════════════════" + Reset)
	fmt.Println(Bold + " Select Mode" + Reset)
	fmt.Println(Purple + "══════════════════════════════" + Reset)

	fmt.Println(" 1) Dry Run")
	fmt.Println(" 2) Execute Rename")
	fmt.Println(" 3) Exit")

	fmt.Print("\n> ")

	var choice int
	fmt.Scanln(&choice)

	return choice
}
