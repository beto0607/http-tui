package ui

import "fmt"

var ANSI_MAP = map[string]string{
	"HOME":              "\033[H",
	"CLEAR":             "\033[2J",
	"HIDE_CURSOR":       "\033[25",
	"MOVE_CURSOR":       "\033[%d;%dH",
	"MOD_BOLD":          "\033[1m",
	"MOD_ITALIC":        "\033[3m",
	"MOD_UNDERLINE":     "\033[4m",
	"MOD_STRIKETHROUGH": "\033[9m",
	"MOD_END":           "\033[0m",
}

func MoveCursorTo(x int, y int) string {
	return fmt.Sprintf(ANSI_MAP["MOVE_CURSOR"], x, y)
}

func MakeBold(content string) string {
	return ANSI_MAP["MOD_BOLD"] + content + ANSI_MAP["MOD_END"]
}

func MakeItalic(content string) string {
	return ANSI_MAP["MOD_ITALIC"] + content + ANSI_MAP["MOD_END"]
}

func MakeUndeline(content string) string {
	return ANSI_MAP["MOD_UNDERLINE"] + content + ANSI_MAP["MOD_END"]
}

func MakeStrikethrough(content string) string {
	return ANSI_MAP["MOD_STRIKETHROUGH"] + content + ANSI_MAP["MOD_END"]
}
