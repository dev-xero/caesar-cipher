package log

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
)

func PrintTitle(word string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("231")).
		Background(lipgloss.Color("#42B0F5")).
		PaddingLeft(2).
		PaddingRight(2).
		Align(lipgloss.Center)

	fmt.Println(style.Render(word))
	fmt.Println()
}

func PrintError(err string) {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("231")).
		Background(lipgloss.Color("#FF584D")).
		PaddingLeft(2).
		PaddingRight(2).
		Align(lipgloss.Center)

	fmt.Println(style.Render(err))
	fmt.Println()
}
