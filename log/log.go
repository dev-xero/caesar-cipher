package log

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Title(word string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("231")).
		Background(lipgloss.Color("#42B0F5")).
		PaddingLeft(2).
		PaddingRight(2).
		Align(lipgloss.Center)

	fmt.Println()
	fmt.Println(style.Render(word))
	fmt.Println()
}

func Error(err error) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("231")).
		Background(lipgloss.Color("#FF584D")).
		PaddingLeft(2).
		PaddingRight(2).
		Align(lipgloss.Center)

	fmt.Println()
	fmt.Println(style.Render("ERROR"), err)
	fmt.Println()
}

func Result(res string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("231")).
		Background(lipgloss.Color("#42F590")).
		PaddingLeft(2).
		PaddingRight(2).
		Align(lipgloss.Center)

	fmt.Println()
	fmt.Println(style.Render("RESULT"), res)
	fmt.Println()
}
