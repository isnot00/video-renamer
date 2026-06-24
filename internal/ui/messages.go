package ui

import "fmt"

func Success(msg string) {
	fmt.Printf("%s✔ %s%s\n", Green, msg, Reset)
}

func Warning(msg string) {
	fmt.Printf("%s⚠ %s%s\n", Yellow, msg, Reset)
}

func Error(msg string) {
	fmt.Printf("%s✖ %s%s\n", Red, msg, Reset)
}

func Info(msg string) {
	fmt.Printf("%sℹ %s%s\n", Blue, msg, Reset)
}
