package ui

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func getTerminalSize() (*Size, error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	chunks := strings.Split(string(out[:len(out)-1]), " ")

	width, err := strconv.Atoi(chunks[1])
	if err != nil {
		return nil, err
	}
	height, err := strconv.Atoi(chunks[0])
	if err != nil {
		return nil, err
	}
	return &Size{
		Width:  width,
		Height: height,
	}, nil
}

const TOP_LEFT_CORNER = "╭"
const TOP_RIGHT_CORNER = "╮"
const BOTTOM_LEFT_CORNER = "╰"
const BOTTOM_RIGHT_CORNER = "╯"
const HORIZONTAL = "─"
const VERTICAL = "│"

func drawRectangle(width int, height int) string {
	if width < 2 || height < 2 {
		return ""
	}
	out := ""

	out += TOP_LEFT_CORNER + strings.Repeat(HORIZONTAL, width-2) + TOP_RIGHT_CORNER + "\n"

	for i := 1; i < height-2; i++ {
		out += VERTICAL + strings.Repeat(" ", width-2) + VERTICAL + "\n"
	}
	out += BOTTOM_LEFT_CORNER + strings.Repeat(HORIZONTAL, width-2) + BOTTOM_RIGHT_CORNER

	return out
}
