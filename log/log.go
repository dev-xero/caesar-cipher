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
		Width(22).
		Align(lipgloss.Center)

	fmt.Println(style.Render(word))
}
